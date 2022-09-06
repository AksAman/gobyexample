// <- chan means data going outside chan, or sending channels
// chan <- means data going inside  chan, or recieving channels

package channels

import "fmt"

type (
	senderChan   <-chan string
	receiverChan chan<- string
)

func ping(pingsChan receiverChan, msg string) {
	pingsChan <- msg
}

func pong(pingsChan senderChan, pongsChan receiverChan) {
	msg := <-pingsChan
	pongsChan <- msg
}

func RunChannelDirections() {
	pingsChannel := make(chan string, 1)
	pongsChannel := make(chan string, 1)

	ping(pingsChannel, "ping message 1")
	pong(pingsChannel, pongsChannel)

	pongMessage := <-pongsChannel
	fmt.Printf("pongMessage: %v\n", pongMessage)
}
