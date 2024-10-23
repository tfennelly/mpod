package main

import (
	"errors"
	"fmt"
	"os"
)


// ************************************************************************
// I don't think any of this is needed of you use the flags library
// ************************************************************************



func main() {
	if len(os.Args) == 1 {
		errors.New("Please provide some arguments\n")
		return
	}
	// detect what the command was
	// if the command is a valid command, pass it to the relevant handler function for that command
	// if the command is invalid, raise an error based on the type of misuse
	command, err := createCommand(os.Args)
	if err != nil {
		panic(err)
	}
	command()
}

func isValidIndex(slice []string, index int) bool {
	return index >= 0 && index < len(slice)
}

func createCommand(programArgs []string) (func(), error) {
	switch programArgs[1] {
	case "start":
		let name = "default"
		// start has non-mandatory flag of '--name'
		if isValidIndex(programArgs, 2) {
			if programArgs[2] != "--name" {

			}
		}
		return func() {
			fmt.Println("Starting the %q thing ..", name)
			// add code to start ...
		}, nil
	case "stop":
		return func() {
			fmt.Println("Stopping the thing ..")
			// add code to stop ...
		}, nil
	default:
		err := errors.New("this is the default err")

		return nil, err
	}
}
