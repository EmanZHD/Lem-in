package algo

// Finding the shortest paths between start and end
func BfsHelper(info *Info, visitedPath [][]string) {
	for _, adj := range info.AdjList[info.End] {
		BfsSearch(info, adj, visitedPath)
	}
}

// Performing BFS search on the graph's nodes and finding the shortest path
func BfsSearch(info *Info, adj string, visitedPath [][]string) {
	// Declaring a queue to use in the BFS Search
	queue := []string{adj}
	path := make(map[string]string)
	path[adj] = info.End
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current == info.Start {
			info.SourcePath = path
			if len(visitedPath) == 0 {
				BuildPath(info, false)
			} else {
				BuildPath(info, true)
			}
			break
		}

		for _, child := range info.AdjList[current] {
			visitedPath := info.Collect
			exist := Contain(visitedPath)
			_, visited := path[child]
			if child != info.End {
				if !visited && !exist[child] && !exist[current] {
					// fmt.Printf("EXIST ============= %v - \n", child)
					path[child] = current
					queue = append(queue, child)
				}
			}
		}

	}
}
