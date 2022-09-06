package channels

import "fmt"

func RunNonBlockingChannels() {
	messagesChan := make(chan string)

	select {
	case msg := <-messagesChan:
		fmt.Printf("msg rcvd: %v\n", msg)
	default:
		fmt.Println("no message rcvd")
	}

	msg := "hi"
	select {
	case messagesChan <- msg:
		fmt.Printf("msg sent: %v\n", msg)
	default:
		fmt.Println("no message sent")
	}

}
