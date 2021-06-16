package concurrent

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
)

func doPrintInOrder() {
	f := &Foo{}
	go f.third()
	go f.first()
	go f.second()

	time.Sleep(1 * time.Second)
}

type Foo struct {
}

func (f *Foo) first() {
	fmt.Println("first")
	ch1 <- 1
}

func (f *Foo) second() {
	<-ch1
	fmt.Println("second")
	ch2 <- 1
}

func (f *Foo) third() {
	<-ch2
	fmt.Println("third")
}
