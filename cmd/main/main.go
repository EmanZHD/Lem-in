package main

import (
	"os"
	"strings"

	"Lem-in-izahid/cmd"
)

func main() {
	test := cmd.ParseFIle()
	cmd.FindGroups(test)
	cmd.PrintPaths(test)
}

func init() {
	switch true {
	case len(os.Args) != 2:
		cmd.Error("Usage: go run <programName> <fileName>")
	case !(strings.HasSuffix(os.Args[1], ".txt")):
		cmd.Error("INvalid file extension")
	}
}

// L1-2 L2-3
// L1-1 L2-1 L3-2
// L3-1
