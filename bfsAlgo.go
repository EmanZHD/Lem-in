package main

func BfsHelper(info *Info, visitedPath [][]string) {
	visitedPath = info.Collect
	for _, adj := range info.Graph[info.End] {
		BfsAlgo(info, adj, visitedPath)
		// delete(info.Visit, adj)
		// delete(info.Visit, info.Start)
	}
}

func BfsAlgo(info *Info, adj string, visitedPath [][]string) {
	q := []string{adj}
	node := adj
	prevNOde := info.End
	Visit := make(map[string]string)
	// var exist bool
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		Visit[node] = prevNOde
		if current != info.Start {
			for _, child := range info.Graph[current] {
				visitedPath := info.Collect
				exist := Contain(visitedPath)
				_, visited := Visit[child]
				if child != info.End {
					if !visited && !exist[child] && !exist[current] {
						// fmt.Printf("EXIST ============= %v - \n", child)
						Visit[child] = current
						q = append(q, child)
					}
				}
			}
		} else {
			info.PathSource = Visit
			if len(visitedPath) == 0 {
				findShortPath(info, false)
			} else {
				findShortPath(info, true)
			}
			break
		}
	}

	// fmt.Println("visited ", info.PathSource)
	// return visit
}
