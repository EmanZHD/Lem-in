package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Info struct {
	Data          []string
	Nbr           int
	Rooms         []string
	Start         string
	End           string
	Links         []string
	AdjList       map[string][]string
	PathSource    map[string]string
	ShortestPaths [][]string
	Collect       [][]string
	Groups        [][][]string
	// Visit        map[string]string
}

func Error(mssg any) {
	fmt.Println("\033[31m"+"Error:", fmt.Sprint(mssg)+"\033[0m")
	os.Exit(0)
}

func Contain(paths [][]string) map[string]bool {
	exist := make(map[string]bool)
	for _, path := range paths {
		for _, elem := range path {
			exist[elem] = true
		}
	}
	// fmt.Println("rrrrrrrr ", exist)
	return exist
}

func SortPaths(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

func IsValid(line string, key string) bool {
	exp := regexp.MustCompile("^(#|L)")
	switch true {
	case exp.MatchString(line):
		Error("Invalid room => cannot start with '#' or 'L'")
	case key == "room":
		room := strings.Split(line, " ")
		if len(room) != 3 {
			fmt.Println("ff ", room)
			Error("invalid room format")
		}
	case key == "link":
		link := strings.Split(line, "-")
		if len(link) != 2 {
			Error("invalid link format")
		}
	}
	return true
}
