package main

import (
	"os"
	"runtime/trace"
)

// trace
// 1. 创建
// 2. 启动
// 3. 停止
func main() {
	// 1
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 2
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常的业务

	// 3
	trace.Stop()
}
