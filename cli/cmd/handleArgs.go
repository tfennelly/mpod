package main

import (
	"errors"
	"fmt"
	"os"
)

// Defines the *structure* of a valid command in application logic
// NOTE that the flags - such as name are arbitrary, and validity is only dependent on position thus they are omitted from this struct
type command struct {
	use          string
	argsAndFlags []string // any arbitrary args - this struct doesn't verify validity of args
	short        string
	long         string
	run          func()
}

func main() {
	if len(os.Args) == 1 {
		errors.New("Please provide some arguments\n")
		return
	}
	// detect what the command was
	// if the command is a valid command, pass it to the relevant handler function for that command
	// if the command is invalid, raise an error based on the type of misuse
	command, _ := createCommand(os.Args)
	command.run()
}

func isValidIndex(slice []string, index int) bool {
	return index >= 0 && index < len(slice)
}

func createCommand(programArgs []string) (command, error) {
	switch programArgs[1] {
	case "start":
		// start has non-mandatory flag of '--name'
		if isValidIndex(programArgs, 2) {
			if programArgs[2] != "--name" {

			}
		}
		// check for validity of start command logic
		thisStartCommand := command{
			use:          programArgs[0],
			argsAndFlags: []string{},
			short:        "scrapShort",
			long:         "scrapLong",
			run: func() {
				fmt.Println("scrapPrint")
			},
		}
		return thisStartCommand, nil
	case "stop":
		// check for validity of stop command logic
		thisStopCommand := command{
			use:          programArgs[0],
			argsAndFlags: []string{},
			short:        "scrapShort",
			long:         "scrapLong",
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
