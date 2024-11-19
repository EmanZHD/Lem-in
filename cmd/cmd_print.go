package cmd

import (
	"fmt"
	"sort"
	"strconv"
)

func PrintPaths(info *Info) {
	bestPath := make(map[int][]int)
	// tempArr := []int{}
	for i, group := range info.Groups {
		group = SortPaths(group)
		for j := 0; j < len(group)-1; j++ {
			fmt.Println()
			fmt.Println("test ", i)
			index := strconv.Itoa(j)
			ant := "L" + index
			for len(group[j]) <= len(group[j+1])+info.Nbr {
				group[j] = append(group[j], ant)
				info.Nbr--
				if info.Nbr <= 0 {
					break
				}
			}
			fmt.Println("-- ", group[j])
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
