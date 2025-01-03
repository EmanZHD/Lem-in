package algo

import (
	"fmt"
	"os"
	"strings"
)

// Struct of ant data
type Info struct {
	Data          []string
	AntNbr        int
	Rooms         []string
	Start         string
	End           string
	Links         []string
	AdjList       map[string][]string
	SourcePath    map[string]string
	ShortestPaths [][]string
	Collect       [][]string
	Groups        [][][]string
	BestGroup     map[int][]string
	Turns         int
	Error         error
}

// Ant struct for printing
type Ant struct {
	Id        string
	Path      []string
	StartLine int
}

// Print Error
func Error(mssg any) {
	fmt.Println("\033[31m"+"Error:", fmt.Sprint(mssg)+"\033[0m")
	os.Exit(0)
}

// Mark rooms in a found path as occupied
func Contain(paths [][]string) map[string]bool {
	exist := make(map[string]bool)
	for _, path := range paths {
		for _, elem := range path {
			exist[elem] = true
		}
	}
	return exist
}

// Check if a line is valid
func IsValidLine(line string, key string) bool {
	switch true {
	case key == "room":
		room := strings.Split(line, " ")
		if len(room) != 3 || room[0][0] == 'L' {
			Error("invalid room format")
		}
	case key == "link":
		link := strings.Split(line, "-")
		if len(link) != 2 || link[0][0] == 'L' {
			Error("invalid link format")
		}
	}
	return true
}

// Get the group of path with the minimum number of steps and turns
func MinSteps(stepTurns map[int][]int) int {
	first := true
	minKey, minValue := 0, 0
	var index int
	for i, value := range stepTurns {
		if first {
			minKey = value[0]
			minValue = value[1]
			first = false
			index = i
		} else if value[1] < minValue || (value[1] == minValue && value[0] < minKey) {
			minKey = value[0]
			minValue = value[1]
			index = i
		}
	}
	return index
}
