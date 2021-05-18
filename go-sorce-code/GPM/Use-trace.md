# 可视化的GMP编程

## trace
- 创建trace `f, err := os.Create("trace.out")`
- 启动trace `err = trace.Start(f)`
- 停止trace `trace.Stop()`
- `go build` 获得 `trace.out` 文件

运行 `go tool trace trace.out`

## DEBUG
- 将代码编译 `go build XXX.go`
- `GODEBUG=schedtrace=1000 ./[可执行文件]` 1000代表1000ms打印一次
[详细介绍] (https://www.kancloud.cn/aceld/golang/1958305#5GMP_305)
