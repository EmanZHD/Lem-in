package cmd

import (
	"fmt"
	"sort"
	"strconv"
)

func PrintPaths(info *Info) {
	bestPath := make(map[int][]int)
	// tempArr := []int{}
	count := 1
	for i, group := range info.Groups {
		group = SortPaths(group)
		fmt.Println("test ", i)
		for j := 0; j < len(group); j++ {
			fmt.Println("group ", j, " path ", group[j])
			for info.Nbr > count {
				if j+1 <= len(group[j]) {
					for len(group[j]) <= len(group[j+1]) {
						fmt.Println("track ants ", info.Nbr)
						index := strconv.Itoa(count)
						ant := "L" + index
						group[j] = append(group[j], ant)
						count++
						info.Nbr--
					}
				}
			}
			fmt.Println("-- ", group[j])
			// }

		}
	}

	fmt.Println("LAST ", bestPath)
	// fmt.Println("=> ", path)
}

func SortPaths(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	return paths
}

// func BestGroup(info *Info) {
// 	for _, group := range info.Groups {

// 	}
// }

// func print
func PrintAnts(info *Info) {
	for _, group := range info.Groups {
		countAnts := 1
		// group := info.Groups[1]
		for countAnts <= info.Nbr {
			for i := 0; i < len(group); i++ {
				path := group[i]
				if len(group) == 1 {
					for {
						group[i] = append(group[i], fmt.Sprintf("L%v-%v", countAnts, path[i]))
						countAnts++
						if countAnts == info.Nbr+1 {
							break
						}
					}
				} else {
					if i+1 < len(group) {
						for len(group[i]) <= len(group[i+1]) {
							fmt.Println("countAnts", countAnts, "info.Nbr", info.Nbr)
							if countAnts == info.Nbr+1 {
								break
							}
							group[i] = append(group[i], fmt.Sprintf("L%v-%v", countAnts, path[i]))
							countAnts++
						}
						if len(group[i]) > len(group[i+1]) {
							group[i+1] = append(group[i+1], fmt.Sprintf("L%v--%v", countAnts, group[i+1][i]))
							countAnts++
							if countAnts <= info.Nbr+1 {
								break
							}

						}
					}
				}

			}
		}
		fmt.Println("path ", group)
	}
}
