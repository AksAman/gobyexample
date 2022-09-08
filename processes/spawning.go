package processes

import (
	"fmt"
	"io"
	"os/exec"
)

func runSpawningExamples() {
	Title("runSpawningExamples")

	Title("Output()")
	dateCmd := exec.Command("date")
	dateOut, _ := dateCmd.Output()
	fmt.Println("> date")
	fmt.Printf("dateOut: %v\n", string(dateOut))

	Title("errors()")
	invalidOutput, err := exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", e)
		case *exec.ExitError:
			fmt.Println("command exit with code=", e.ExitCode())
		default:
			panic(err)
		}
	} else {
		fmt.Printf("invalidOutput: %v\n", string(invalidOutput))
	}

	Title("pipes")
	grepCmd := exec.Command("grep", "hell")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello world\nworld hello\n"))
	grepIn.Write([]byte("hellboi"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> | grep hell")
	fmt.Printf("string(grepBytes): %v\n", string(grepBytes))

}
