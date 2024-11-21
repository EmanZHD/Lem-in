package cmd

import (
	"fmt"
)

// func BestGroup(info *Info) {
// 	for _, group := range info.Groups {

// 	}
// }

// func print
func PrintAnts(info *Info) {
	// for _, group := range info.Groups {
	AntsPath := make(map[int][]string)
	countAnts := 1
	group := info.Groups[0]
	for countAnts <= info.Nbr {
		for i := 0; i < len(group); i++ {
			path := group[i]
			if len(group) != 1 {
				if i+1 < len(group) {
					if len(group[i]) <= len(group[i+1]) {
						for len(group[i]) <= len(group[i+1]) {
							group[i] = append(group[i], fmt.Sprintf("L%v", countAnts))
							AntsPath[i] = append(AntsPath[i], fmt.Sprintf("L%v-%v", countAnts, path[i]))
							countAnts++
						}
					} else if len(group[i]) > len(group[i+1]) {
						group[i+1] = append(group[i+1], fmt.Sprintf("L%v", countAnts))
						AntsPath[i+1] = append(AntsPath[i+1], fmt.Sprintf("L%v-%v", countAnts, group[i+1][i]))
						countAnts++

					}
					if countAnts == info.Nbr+1 {
						break
					}

				} else if i == len(group)-1 {
					group[len(group)-1] = append(group[len(group)-1], fmt.Sprintf("L%v", countAnts))
					AntsPath[len(group)-1] = append(AntsPath[len(group)-1], fmt.Sprintf("L%v-%v", countAnts, path[i]))
					countAnts++
				}
			} else {
				for {
					group[i] = append(group[i], fmt.Sprintf("L%v", countAnts))
					AntsPath[i] = append(AntsPath[i], fmt.Sprintf("L%v-%v", countAnts, path[i]))
					countAnts++
					if countAnts == info.Nbr+1 {
						break
					}
				}
			}

		}
	}
	fmt.Println("TEST ", AntsPath)
	fmt.Println("path ", group)
	// }
}

func Printtttt(){
	
}