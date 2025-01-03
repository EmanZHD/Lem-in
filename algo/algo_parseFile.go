package algo

import (
	"os"
	"strconv"
	"strings"
)

// Reading test file and splitting it's content to lines
func ParseFIle() *Info {
	file := os.Args[1]
	data, _ := os.ReadFile(file)
	if len(data) == 0 {
		Error("Empty file")
	}
	lines := []string{}

	// deleting empty lines
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) != 0 {
			lines = append(lines, line)
		} else {
			continue
		}
	}
	info := &Info{
		Data: lines,
	}

	ProcessLines(info)
	BfsHelper(info, [][]string{})
	return info
}

// Process file's lines and fill in the info in struct
func ProcessLines(info *Info) {
	antTotal, err := strconv.Atoi(info.Data[0])
	switch true {
	case err != nil:
		Error(err)
	case antTotal <= 0:
		Error("Invalid ant number")
	}

	doubleS := 0
	doubleE := 0

	for i := 1; i < len(info.Data); i++ {

		if info.Data[i] == "##start" {
			doubleS++
			if IsValidLine(info.Data[i+1], "room") {
				info.Start = strings.Split((info.Data[i+1]), " ")[0]
			}
		}
		if info.Data[i] == "##end" {
			doubleE++
			if IsValidLine(info.Data[i+1], "room") {
				info.End = strings.Split((info.Data[i+1]), " ")[0]
			}
		}

		if info.Data[i][0] != '#' {
			if strings.Contains(info.Data[i], "-") && IsValidLine(info.Data[i], "link") {
				info.Links = append(info.Links, info.Data[i])
			} else if IsValidLine(info.Data[i], "room") {
				info.Rooms = append(info.Rooms, strings.Split(info.Data[i], " ")[0])
			}
		}
		if doubleE > 1 || doubleS > 1 {
			Error("Start or End room already exists")
		}
	}

	if len(info.Start) == 0 || len(info.End) == 0 || len(info.Rooms) == 0 || len(info.Links) == 0 {
		Error("Incomplete data")
	}

	info.AntNbr = antTotal
	info.AdjList = BuildAdjList(info)
}
