# Golang 的堆和栈
先呈现一个Golang中的现象
```
package main

func foo(arg_val int)(*int) {

    var foo_val int = 11;
    return &foo_val;
}

func main() {

    main_val := foo(666)

    println(*main_val)
}
```
在其他如C/C++语言中，这段代码无法编译，原因是 `foo` 函数返回了局部变量 `foo_val` 的地址，这在其他语言中是禁止的，原因是 `foo_val`变量的生命周期只在 `foo` 函数内，当它在 `main` 函数中被调用时，已经被销毁了。但是，这段代码却可以在 Golang中正常运行。

## Golang 编译器的逃逸分析
Golang 编译器会自动决定把一个变量放在栈还是堆上，编译器会做出逃逸分析(escape analysis)， 当变量的作用域仍在函数范围内，就放在栈上，否则放在堆上。

