package main

import (
	"os"
	"strings"

	"Lem-in-izahid/algo"
)

func main() {
	data := algo.ParseFIle()
	algo.FindGroups(data)
	algo.MoveAnts(data)
}

func init() {
	switch true {
	case len(os.Args) != 2:
		algo.Error("Usage: go run <programName> <fileName>")
	case !(strings.HasSuffix(os.Args[1], ".txt")):
		algo.Error("Invalid file extension")
	}
}
