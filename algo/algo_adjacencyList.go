package algo

import "strings"

// BuildAdjList Build the adjacency list from the graph
func BuildAdjList(info *Info) map[string][]string {
	links := info.Links
	rooms := info.Rooms
	path := map[string][]string{}
	for _, link := range links {
		a := strings.Split(link, "-")
		r1, r2 := a[0], a[1]
		for _, room := range rooms {
			switch true {
			case r1 == r2:
				Error("Invalid link")
			case room == r1:
				path[room] = append(path[room], r2)
			case room == r2:
				path[room] = append(path[room], r1)
			}
		}
	}
	return path
}
