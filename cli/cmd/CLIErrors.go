package main

import (
	"strings"
	"fmt"
)

type CLIErr struct {
	name    string
	message string
}

// An example of an argument is 'mpod start', where start is the argument
func InvalidArg(command string, invalidArg string) CLIErr {
	return newCLIErr("InvalidArgumentError", command, invalidArg)
}

// A flag can be thought of as 'an argument to an argument'
// e.g 'mpod start --name "my-container"' where '--name' is the flag
func InvalidFlag(argument string, invalidFlag string) CLIErr {
	return newCLIErr("InvalidFlagError", command, invalidArg)
}

func newCLIErr(errorType string, argument string, invalidFlag string) CLIErr {
	invalidFlagError := CLIErr{
		name:    "InvalidFlagError",
		message: fmt.Sprintf("%s: '%s' is not a valid flag for argument: %s", errorType, argument, invalidFlag),
	}

	return invalidFlagError
}
