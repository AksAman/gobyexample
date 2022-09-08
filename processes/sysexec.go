package processes

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func runSysCallExample() {
	Title("runSysCallExample")
	lsBinary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	fmt.Printf("ls binary found at : %v\n", lsBinary)

	args := []string{"ls", "-a", "-l", "-h"}

	execErr := syscall.Exec(lsBinary, args, os.Environ())
	if execErr != nil {
		panic(execErr)
	}
}
