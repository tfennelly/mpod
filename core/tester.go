package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		// this case checks if the type of e (=err.(type)) is equal to an
		//	execution error type (exec.Error)
		case *exec.Error:
			fmt.Println("failed executing: ", err)
		//	checks if the error is of type exec.ExitError
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
}
