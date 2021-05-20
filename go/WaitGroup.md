# 0. 主要内容
- WaitGroup的定义
- WaitGroup的实现 
    - (`CAS`的使用)
- `ADD`与`DONE`应该放哪里？ 
    - 前者在协程外(即在主协程中)，后者在协程内
- 适合的场景 
    - 适合逻辑结构相同的协程。
    - 复杂场景不适用，因为其代码不简洁
# 1. WaitGroup
`WaitGroup` 是一个定义在`sync`包中用于控制协程的结构体类型。

主协程通过调用`Add`方法来设置需要等待的协程数量，每个协程在它运行完毕时调用`Done`方法，再在主协程中调用`Wait`方法可以阻塞到所有协程运行完毕。

# 2. 实例 与 注意事项
- 这是简单的`WaitGroup`使用场景。
    - 每开启一个`goroutine`前需要通过`Add`方法通知`WaitGroup`需要等待协程数增加。
    - 每个协程结束前通过`Done`方法通知 `WaitGroup`
    - 在主协程中调用 `Wait()` 阻塞主协程。
```
func wg(){
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            fmt.Println(i)
            defer wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println("finished")
}
```
- `WaitGroup` 是一个值而非引用，不要拷贝
    - 在该案例中，我们将 `WaitGroup` 作为参数传入协程中，但协程中的 `WaitGroup` 和主协程中的不是同一个对象。
```
func wg(){
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int, wg sync.WaitGroup) {
            fmt.Println(i)
            defer wg.Done()
        }(i, wg)
    }
    wg.Wait()
    fmt.Println("finished")
}

```

- `Add` 中的数不要随便自定义。
    - 在 `Add` 函数中,每次执行加减，会执行一次信号量的交互，随意增加这个值会频繁交互，浪费性能。
    ```
    for ; w != 0; w-- {
		runtime_Semrelease(semap, false, 0)
	}
    ```

# 3. 源码学习
- 原子操作（重点`CompareAndSwap  -CAS`）
    - `Add` 中:
        - `state := atomic.AddUint64(statep, uint64(delta)<<32)`
    - `Wait`:
        - `state := atomic.LoadUint64(statep)`
        - `atomic.CompareAndSwapUint64(statep, state, state+1)`
    - `CAS` 操作性能比锁要高