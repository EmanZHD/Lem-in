package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Info struct {
	Data         []string
	Nbr          int
	Rooms        []string
	Start        string
	End          string
	Links        []string
	Paths        map[string][]string
	BuildPath    map[string]string
	ShortestPath []string
}

func main() {
	test := ParseFIle()
	fmt.Println(test.Paths)
	fmt.Println("shortestPath ", test.ShortestPath)
	PrintPaths(test)
}

func Error(mssg any) {
	fmt.Println("\033[31m"+"Error:", fmt.Sprint(mssg)+"\033[0m")
	os.Exit(0)
}

func ProcessLines(info *Info) {
	total, err := strconv.Atoi(info.Data[0])
	if err != nil {
		Error(err)
	}
	info.Nbr = total
	for i := 1; i < len(info.Data); i++ {
		if info.Data[i] == "##start" {
			if IsValid(info.Data[i+1], "room") {
				info.Start = strings.Split((info.Data[i+1]), " ")[0]
			}
		}
		if info.Data[i] == "##end" {
			if IsValid(info.Data[i+1], "room") {
				info.End = strings.Split((info.Data[i+1]), " ")[0]
			}
		}
		// exp, _ := regexp.MatchString("^(L|#)", info.Data[i])
		if info.Data[i] != "##end" && info.Data[i] != "#comment" && info.Data[i] != "##start" {
			if strings.Contains(info.Data[i], "-") && IsValid(info.Data[i], "link") {
				info.Links = append(info.Links, info.Data[i])
			} else if IsValid(info.Data[i], "room") {
				info.Rooms = append(info.Rooms, strings.Split(info.Data[i], " ")[0])
			}
		}
	}
	if len(info.Start) == 0 || len(info.End) == 0 || len(info.Rooms) == 0 || len(info.Links) == 0 {
		Error("INcomplete data")
	}
	info.Paths = LinkPaths(info)
	// return &Info{
	// 	Start: info.Start,
	// 	End:   info.End,
	// 	Rooms: info.Rooms,
	// 	Links: info.Links,
	// 	Paths: info.Paths,
	// }
}

func LinkPaths(info *Info) map[string][]string {
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
	// fmt.Println("paths", paths)
	return paths
}

func ParseFIle() *Info {
	file := os.Args[1]
	data, _ := os.ReadFile(file)
	if len(data) == 0 {
		Error("Empty file")
	}
	lines := strings.Split(string(data), "\n")
	info := &Info{
		Data: lines,
	}
	ProcessLines(info)
	BfsAlgo(info)
	return info
}

func IsValid(line string, key string) bool {
	if key == "room" {
		room := strings.Split(line, " ")
		if len(room) != 3 {
			Error("invalid room format")
		}
	}
	if key == "link" {
		link := strings.Split(line, "-")
		if len(link) != 2 {
			Error("invalid link format")
		}
	}
	return true
}

func BfsAlgo(info *Info) {
	q := []string{info.Start}
	visit := make(map[string]string)
	node := info.Start
	prevNOde := "None"
	// q = append(q, info.Start)
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		visit[node] = prevNOde
		if current != info.End {
			for _, child := range info.Paths[current] {
				if _, visited := visit[child]; !visited {
					visit[child] = current
					q = append(q, child)
				}
			}
		} else {
			break
		}
	}
	info.BuildPath = visit
	findShortPath(info)
	fmt.Println("visited ", visit)
	// return visit
}

func findShortPath(info *Info) {
	target := info.End
	tempPath := []string{target}
	for range info.BuildPath {
		// fmt.Println("TEST ", info.BuildPath[startNode])
		prev := info.BuildPath[target]
		for prev != info.Start {
			tempPath = append([]string{prev}, tempPath...)
			prev = info.BuildPath[prev]
		}
		// tempPath = append([]string{info.Start}, tempPath...)
		if prev == info.Start {
			break
		}
	}
	info.ShortestPath = tempPath
}

func PrintPaths(info *Info) {
	path := ``
	for i := 0; i < info.Nbr-1; i++ {
		for _, node := range info.ShortestPath {
			path += fmt.Sprintf("L%v-%v ", i+1, node)
		}
	}
	fmt.Println("=> ", path)
}

func init() {
	switch true {
	case len(os.Args) != 2:
		Error("Usage: go run <programName> <fileName>")
	case !(strings.HasSuffix(os.Args[1], ".txt")):
		Error("INvalid file extension")
	}
}

// L1-2 L2-3
// L1-1 L2-1 L3-2
// L3-1
