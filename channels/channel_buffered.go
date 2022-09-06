package channels

import "fmt"

func RunBuffered() {
	fmt.Println("Buffered Example: ")
	messagesChan := make(chan string, 2)

	messagesChan <- "m1"
	messagesChan <- "m2"

	fmt.Println(<-messagesChan)
	fmt.Println(<-messagesChan)

}
