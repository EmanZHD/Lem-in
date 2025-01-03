package algo

import (
	"fmt"
	"strings"
)

// BestGroup finds the best group that contain minimum steps and turns
func BestGroup(mapArr []map[int][]string, info *Info) {
	stepsArr := []int{}
	tempS, tempT := 0, 0
	stepTurns := make(map[int][]int)
	for i, mapA := range mapArr {
		path := len(info.Groups[i][0])
		ants := len(mapArr[i][0])
		tempT = path - 1 + ants
		for j, grp := range mapA {
			tempS += len(grp) * len(info.Groups[i][j])
		}
		stepsArr = append(stepsArr, tempS, tempT)
		stepTurns[i] = stepsArr
		tempS = 0
		stepsArr = nil
	}

	if len(mapArr) == 0 {
		Error("No path was found")
	}
	minSteps := MinSteps(stepTurns)
	info.BestGroup = mapArr[minSteps]
	info.Turns = stepTurns[minSteps][1]
	PrintAnts(info, minSteps)
}

// PrintAnts Prints the ants movements
func PrintAnts(info *Info, index int) {
	AntMoves := info.BestGroup
	paths := info.Groups[index]
	var ants []Ant

	for i, moves := range AntMoves {
		for moveIndex, move := range moves {
			elem := Ant{
				Id:        move,
				Path:      paths[i],
				StartLine: moveIndex,
			}
			ants = append(ants, elem)
		}
	}
	lines := make([][]string, info.Turns)
	for _, ant := range ants {
		node := 0
		for i := ant.StartLine; i < info.Turns; i++ {
			if node < len(ant.Path) {
				lines[i] = append(lines[i], ant.Id+"-"+ant.Path[node])
				node++
			}
		}
	}
	output := ""
	for i, line := range lines {
		if i == len(lines)-1 {
			output += strings.Join(line, " ")
		} else {
			output += strings.Join(line, " ") + "\n"
		}
	}

	fmt.Println(strings.Join(info.Data, "\n") + "\n\n" + output)
}
