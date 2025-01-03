package algo

// BuildPath extracts the path form the map found by the BfsSearch
func BuildPath(info *Info, collect bool) {
	target := info.Start
	tempPath := []string{}

	for range info.SourcePath {
		prev := info.SourcePath[target]
		for prev != info.End {
			tempPath = append(tempPath, prev)
			prev = info.SourcePath[prev]
		}
		if prev == info.End {
			tempPath = append(tempPath, info.End)
			break
		}
	}
	if collect && len(tempPath) != 1 {
		// fmt.Println("\033[31m"+"COLLECT ", info.Group, ""+"\033[0m")
		info.Collect = append(info.Collect, tempPath)
	}
	info.SourcePath = map[string]string{}
	info.ShortestPaths = append(info.ShortestPaths, tempPath)
}

// FindGroups find the group of path that don't intersect with the path source
func FindGroups(info *Info) {
	// fmt.Println("-==================Start finding Groups======================\n", info.ShortestPaths)
	for _, path := range info.ShortestPaths {
		// fmt.Println("----> ", path)
		info.Collect = append(info.Collect, path)
		BfsHelper(info, info.Collect)
		// fmt.Println("--------------------------- ", info.Group)
		info.Groups = append(info.Groups, info.Collect)
		info.Collect = [][]string{}
	}
	// fmt.Println("******************** ", info.Groups)
}
