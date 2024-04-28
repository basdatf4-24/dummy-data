package main

import (
	"os"

	"github.com/mhmdhaekal/dummydata/cmd"
)

func main() {
	var args = os.Args
	if len(args) < 2 {
		panic("Argument not valid")
	}

	command := args[1]
	switch command {
	case "contributor":
		cmd.GenerateContributor()
	case "user":
		cmd.GenerateUser()
	default:
		panic("Unknown command")
	}

}
