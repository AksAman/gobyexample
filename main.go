package main

import (
	"github.com/AksAman/gobyexample/sync_examples"
)

func main() {
	// structembeds.Run()
	// generics.Run()
	// errorhandling.Run()

	// goroutines.Run()
	// goroutines.RunInBackground()
	// goroutines.RunRaceCondition()

	// channels.Run()
	// channels.RunUnbuffered()
	// channels.RunBuffered()
	// channels.RunChannelSynced()
	// channels.RunChannelDirections()
	// channels.RunChannelSelect()
	// channels.RunCloseChannels()
	// channels.RunRangeChannels()

	// timemachine.RunTimeOutExample()
	// timemachine.RunTimers()
	// timemachine.RunTickers()
	// timemachine.RunRateLimiting()
	// timemachine.RunRateLimiting2()

	// workerpools.RunWorkers()

	// sync_examples.RunWithWaitGroups()
	sync_examples.RunAtomicCounter()

}
