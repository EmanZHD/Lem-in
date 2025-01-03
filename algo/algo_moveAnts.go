package algo

import (
	"fmt"
)

// MoveAnts divide paths between ants and define the best group
func MoveAnts(info *Info) {
	mapArr := []map[int][]string{}
	copy := make([][][]string, len(info.Groups))
	for i, group := range info.Groups {
		copy[i] = make([][]string, len(group))
		for j, subGroup := range group {
			copy[i][j] = append([]string{}, subGroup...)
		}
	}
	for _, group := range copy {
		AntsPath := make(map[int][]string)
		countAnts := 1
		for countAnts <= info.AntNbr {
			for i := 0; i < len(group); i++ {
				if len(group) != 1 {
					if i+1 < len(group) {
						if len(group[i]) <= len(group[i+1]) {
							for len(group[i]) <= len(group[i+1]) {
								group[i] = append(group[i], fmt.Sprintf("L%v", countAnts))
								AntsPath[i] = append(AntsPath[i], fmt.Sprintf("L%v", countAnts))
								countAnts++
							}
						} else if len(group[i]) > len(group[i+1]) {
							group[i+1] = append(group[i+1], fmt.Sprintf("L%v", countAnts))
							AntsPath[i+1] = append(AntsPath[i+1], fmt.Sprintf("L%v", countAnts))
							countAnts++

						}
						if countAnts == info.AntNbr+1 {
							break
						}

					} else if i == len(group)-1 {
						group[len(group)-1] = append(group[len(group)-1], fmt.Sprintf("L%v", countAnts))
						AntsPath[len(group)-1] = append(AntsPath[len(group)-1], fmt.Sprintf("L%v", countAnts))
						countAnts++
					}
				} else {
					for {
						group[i] = append(group[i], fmt.Sprintf("L%v", countAnts))
						AntsPath[i] = append(AntsPath[i], fmt.Sprintf("L%v", countAnts))
						countAnts++
						if countAnts == info.AntNbr+1 {
							break
						}
					}
				}
			}
		}
		mapArr = append(mapArr, AntsPath)
	}

	BestGroup(mapArr, info)
}
