package main

import (
	"errors"
	"fmt"
	"os"
)

type command struct {
	use   string
	short string
	long  string
	run   func()
}

var startCMD = command{
	use:   "start",
	short: "short: start",
	long:  "long: start",
}

var stopCMD = command{
	use:   "stop",
	short: "short: stop",
	long:  "long: stop",
}

func main() {
	if len(os.Args) == 1 {
		errors.New("Please provide some arguments\n")
		return
	}
	// detect what the command was
	// if the command is a valid command, pass it to the relevant handler function for that command
	// if the command is invalid, raise an error based on the type of misuse
}

func detectCommand(passedCMD command) (command, error) {
	validCommandsMap := map[string]bool{"start": true, "stop": true}
	commandsIndex := map[string]command{"start": startCMD, "stop": stopCMD}

	if validCommandsMap[passedCMD.use] == true {
		thisCommand := commandsIndex[passedCMD.use]
		return thisCommand, nil
	}

	return command{}, fmt.Errorf("%s is not a recognised command\n", passedCMD.use)
}

func startContainer() {
	fmt.Println("start a container")
}
