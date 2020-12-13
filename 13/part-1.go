package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n")
	fmt.Println(file)
	ans1 := part1(file)
	fmt.Println("part1: ",ans1)
}

func part1(data []string) int {
	time, _ := strconv.Atoi(data[0])
	buses := []int{}
	fmt.Println("Time: ", time)
	for _, i := range strings.Split(data[1],",") {
		if i == "x" { continue }
		temp, _ := strconv.Atoi(i)
		buses = append(buses, temp)
	}
	fmt.Println(buses)
	ans, ansBus := math.Inf(1), -1
	fmt.Println(ans)
	for _, bus := range buses {
		next := bus - time % bus
		fmt.Println(next + time)
		if float64(next) < ans { 
			ans = float64(next)
			ansBus = bus
		}
	}
	fmt.Println(ans)
	return int(ans)*ansBus
}
