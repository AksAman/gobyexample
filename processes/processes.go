package processes

import (
	"fmt"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal("error: ", e)
	}
}

func Title(s string) {
	fmt.Printf("\n----- %v -----\n", s)
}

func Run() {
	// runSpawningExamples()
	// runSysCallExample()
	runSignalsExample()
}
