package main

import (
	"os"
	"strconv"
	"strings"
)

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
	BfsHelper(info, [][]string{})
	// BfsAlgo(info)
	return info
}

func ProcessLines(info *Info) {
	total, err := strconv.Atoi(info.Data[0])
	switch true {
	case err != nil:
		Error(err)
	case total == 0:
		Error("INvalid ant number zas provided")
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
		// comment := info.Data[i][0] != '#' && IsValid(info.Data[i], "link")
		if info.Data[i] != "##end" && info.Data[i] != "##start" {
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
	info.Graph = BuildGraph(info)
}
