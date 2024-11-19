package cmd

import (
	"fmt"
	"sort"
)

func PrintPaths(info *Info) {
	bestPath := make(map[int][]int)
	// tempArr := []int{}
	for _, group := range info.Groups {
		// SortPaths(group)
		fmt.Println("** ", SortPaths(group))
		// i := 0
		// for i <= info.Nbr {
		// 	for j := 0; j < len(group); j++ {
		// 		fmt.Println("GG ", group)

		// 		// for len(group[j]) < len(group[j+1]) && (j <= len(group)-1) {
		// 		// 	tempArr = append(tempArr, i)
		// 		// 	i--
		// 		// }
		// 		// bestPath[j] = tempArr
		// 	}
		// }
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
