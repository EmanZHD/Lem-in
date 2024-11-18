package main

import "strings"

// [t E a m end] [0 o n e end] [h A c k end]
func BuildGraph(info *Info) map[string][]string {
	links := info.Links
	rooms := info.Rooms
	paths := map[string][]string{}
	for _, link := range links {
		a := strings.Split(link, "-")
		r1, r2 := a[0], a[1]
		for _, room := range rooms {
			switch true {
			case r1 == r2:
				Error("invalid link")
			case room == r1:
				paths[room] = append(paths[room], r2)
			case room == r2:
				paths[room] = append(paths[room], r1)
			}
		}
	}
	return paths
}
