package main

import (
	"os"
	"strings"
)

type Info struct {
	Data          []string
	Nbr           int
	Rooms         []string
	Start         string
	End           string
	Links         []string
	Graph         map[string][]string
	PathSource    map[string]string
	ShortestPaths [][]string
	Collect       [][]string
	Groups        [][][]string
	// Visit        map[string]string
}

func main() {
	test := ParseFIle()
	FindGroups(test)
	PrintPaths(test)
}

func init() {
	switch true {
	case len(os.Args) != 2:
		Error("Usage: go run <programName> <fileName>")
	case !(strings.HasSuffix(os.Args[1], ".txt")):
		Error("INvalid file extension")
	}
}

// L1-2 L2-3
// L1-1 L2-1 L3-2
// L3-1
