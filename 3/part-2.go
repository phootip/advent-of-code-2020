package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	// "strconv"
)

func main() {
	dat, _ := ioutil.ReadFile("./input.txt")
	// dat, _ := ioutil.ReadFile("./example.txt")
	data := strings.Split(string(dat),"\n")
	var maps []string
	slopes := [5][2]int {
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}
	for _, line := range data[:len(data)-1] {
		maps = append(maps, line)
	}
	maxX := len(maps[0])
	var buffer []int
	for _, slope := range slopes {
		slopeX, slopeY := slope[0], slope[1]
		x,y,ans := 0,0,0
		fmt.Println(slopeX, slopeY)
		for _, line := range maps {
			if slopeY != 1 && y == slopeY-1 {
				y = 0
				continue
			}
			fmt.Println(line, x, y, string(line[x]))
			if string(line[x]) == "#" {
				ans++
			}
			x += slopeX
			if x >= maxX {
				x -= maxX
			}
			y++
		}
		fmt.Println(ans)
		buffer = append(buffer,ans)
	}
	fmt.Println(buffer)
	ans := 1
	for _, temp := range buffer {
		ans *= temp
	}
	fmt.Println(ans)
}

// 4187192320 wrong
