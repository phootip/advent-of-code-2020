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
	x,y,ans := 0,0,0
	for _, line := range data[:len(data)-1] {
		maps = append(maps, line)
	}
	maxX := len(maps[0])
	for _, line := range maps {
		fmt.Println(line, x, y, string(line[x]))
		if string(line[x]) == "#" {
			ans++
		}
		x += 3
		if x >= maxX {
			x -= maxX
		}
		y++
	}
	fmt.Println(ans)
}
