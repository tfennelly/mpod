package main

import (
	"errors"
	"fmt"
	"os"
)

// Defines the structure of a valid command
// this will have to be checked to see if the args are valid for the command
// NOTE that the flags - such as name are arbitrary, and validity is only
// dependent on position thus they are omitted from this struct
// 		e.g. the name of container to start must be start.index + 1
type command struct {
	use   string
	args  []string // any arbitrary args - this struct doesn't verify validity of args
	short string
	long  string
	run   func()
}

func main() {
	if len(os.Args) == 1 {
		errors.New("Please provide some arguments\n")
		return
	}
	// detect what the command was
	// if the command is a valid command, pass it to the relevant handler function for that command
	// if the command is invalid, raise an error based on the type of misuse
	createCommand(os.Args)
}

func createCommand(programArgs []string) (command, error) {
	switch programArgs[1] {
	case "start":
		fmt.Println(programArgs[2])
		// check for validity of start command logic
		thisStartCommand := command{
			use:   programArgs[0],
			args:  []string{},
			short: "scrapShort",
			long:  "scrapLong",
			run: func() {
				fmt.Println("scrapPrint")
			},
		}
		return thisStartCommand, nil
	case "stop":
		// check for validity of stop command logic
		thisStopCommand := command{
			use:   programArgs[0],
			args:  []string{},
			short: "scrapShort",
			long:  "scrapLong",
			run: func() {
				fmt.Println("scrapPrint")
			},
		}
		return thisStopCommand, nil
	default:
		err := errors.New("this is the default err")

		return command{}, err
	}
}
