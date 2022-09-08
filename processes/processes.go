package processes

import (
	"fmt"
)

func Title(s string) {
	fmt.Printf("\n----- %v -----\n", s)
}

func Run() {
	runSpawningExamples()
	runSysCallExample()
	runSignalsExample()
}
