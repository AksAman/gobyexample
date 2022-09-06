package goroutines

import (
	"fmt"
)

const MAX = 2

func worker1(val *int, sumChan chan int, workerName string) int {
	for i := 0; i <= MAX; i++ {
		fmt.Printf("%v: adding %d to initial: %d\n", workerName, i, *val)
		*val += i
		fmt.Printf("%v: Added %d, equals: %d\n", workerName, i, *val)
	}
	sumChan <- *val
	return *val
}

func RunRaceCondition() {
	val := 0

	totalChan := make(chan int)

	go worker1(&val, totalChan, "worker-1")
	go worker1(&val, totalChan, "worker-2")
	<-totalChan
	fmt.Printf("after first val: %v\n", val)

	sum := <-totalChan

	directSum := MAX * (MAX + 1) / 2

	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("directSum: %v\n", directSum)

}
