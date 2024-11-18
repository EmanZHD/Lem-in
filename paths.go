package main

import (
	"fmt"
)

func findShortPath(info *Info, collect bool) {
	target := info.Start
	tempPath := []string{}
	fmt.Println("path ", info.PathSource)
	for range info.PathSource {
		prev := info.PathSource[target]
		for prev != info.End {
			// tempPath = append([]string{prev}, tempPath...)
			tempPath = append(tempPath, prev)
			prev = info.PathSource[prev]
		}
		// tempPath = append([]string{info.Start}, tempPath...)
		if prev == info.End {
			tempPath = append(tempPath, info.End)
			break
		}
	}
	fmt.Println("test ", tempPath)
	if collect {
		fmt.Println("\033[31m"+"COLLECT ", info.Collect, ""+"\033[0m")
		info.Collect = append(info.Collect, tempPath)
	}
	info.PathSource = map[string]string{}
	info.ShortestPaths = append(info.ShortestPaths, tempPath)
}

func FindGroups(info *Info) {
	fmt.Println("-==================Start finding Groups======================\n", info.ShortestPaths)
	// Visited := make(map[string]string)
	for _, path := range info.ShortestPaths {
		fmt.Println("----> ", path)
		info.Collect = append(info.Collect, path)
		BfsHelper(info, info.Collect)
		fmt.Println("--------------------------- ", info.Collect)
		info.Groups = append(info.Groups, info.Collect)

		info.Collect = [][]string{}
	}
	// fmt.Println("******************************************************\n", info.Groups)
	for _, group := range info.Groups {
		fmt.Println("******************** ", group)
	}
}
