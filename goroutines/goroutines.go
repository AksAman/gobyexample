package goroutines

import (
	"fmt"
	"time"
)

func task(arg string) {
	fmt.Println("Task started", arg)
	for i := 0; i < 3; i++ {
		fmt.Printf("%v:%d\n", arg, i)
	}
	fmt.Println("Task finished", arg)
}

func Run() {
	task("direct")

	go task("indirect")

	go func(arg string) {
		task(arg)
	}("anonymous")

	time.Sleep(time.Second / 50)
	fmt.Println("EOP")
}
