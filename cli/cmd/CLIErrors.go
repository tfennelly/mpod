package main

import (
	"strings"
)

type CLIErr struct {
	name    string
	message string
}

// An example of an argument is 'mpod start', where start is the argument
func InvalidArg(command string, invalidArg string) CLIErr {
	var messageSB strings.Builder
	messageSB.WriteString("InvalidArgumentError: ")
	messageSB.WriteString("'" + invalidArg + "'")
	messageSB.WriteString(" is not a valid argument for command: ")
	messageSB.WriteString(command)

	InvalidArgumentError := CLIErr{
		name:    "InvalidArgumentError",
		message: messageSB.String(),
	}

	return InvalidArgumentError
}

// A flag can be thought of as 'an argument to an argument'
// e.g 'mpod start --name "my-container"' where '--name' is the flag
func InvalidFlag(argument string, invalidFlag string) CLIErr {
	var messageSB strings.Builder
	messageSB.WriteString("InvalidFlagError: ")
	messageSB.WriteString("'" + invalidFlag + "'")
	messageSB.WriteString(" is not a valid flag for argument: ")
	messageSB.WriteString(argument)

	InvalidFlagError := CLIErr{
		name:    "InvalidFlagError",
		message: messageSB.String(),
	}

	return InvalidFlagError
}
