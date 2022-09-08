package processes

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*

Sometimes we’d like our Go programs to intelligently handle Unix signals.
For example,
we might want a server to gracefully shutdown
	when it receives a SIGTERM,
or a command-line tool to stop processing input
	if it receives a SIGINT.
*/

func runSignalsExample() {
	Title("runSignalsExample")

	// os signals should be buffered
	sigs := make(chan os.Signal, 1)

	// register types of signals
	// SIGINT is the signal sent when we press Ctrl+C
	// The SIGTERM and SIGQUIT signals are meant to terminate the process. (default signal when we use the kill command.)
	// SIGTSTP: Ctrl+Z
	// SIGQUIT: Ctrl+\
	// https://www.baeldung.com/linux/sigint-and-other-termination-signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)

	sigReceived := make(chan bool, 1)

	go func() {
		// wait for signals
		sig := <-sigs
		fmt.Printf("sig: %v, %#v\n", sig, sig)
		sigReceived <- true
	}()

	fmt.Println("Awaiting signal")
	<-sigReceived
	fmt.Println("exiting")
}
