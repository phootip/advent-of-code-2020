package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

var (
	mem [][][]string
)

func main() {
	fmt.Println("Start Day 11...")
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := [][]string{}
	for _, i := range strings.Split(string(rawDat[:len(rawDat)-1]),"\n"){
		file = append(file, strings.Split(i,""))
	}
	pprint("start:",file)
	ans := part2(file)
	fmt.Println(ans)
}

func part2(data [][]string) int{
	fmt.Println("Starting Part 1...")
	nextData  := make([][]string, len(data))
	deepcopy(nextData,data)
	for y, row := range data {
		for x, col := range row {
			_ = col
			switch data[y][x] {
				case ".":
					continue
				case "L":
					if getAdjOcc(data,x,y) == 0 {
						nextData[y][x] = "#"
					}
				case "#":
					if getAdjOcc(data,x,y) >= 4 {
						nextData[y][x] = "L"
					}
				default:
					panic("Wrong seat type")
			}
		}
	}
	// pprint("data:", data)
	pprint("nextData:", nextData)
	temp := make([][]string, len(data))
	deepcopy(temp, nextData)
	if contain(nextData) {
		return getAns(nextData)
	}
	mem = append(mem, temp)
	return part1(nextData)
}

func getAdjOcc(data [][]string, x int, y int) int {
	// checkList := [][]int{}
	result := 0
	if data[y][x] == "#" {result--}
	for i := x-1; i <= x+1; i++ {
		for j := y-1; j <= y+1; j++ {
			func () {
				defer func() {
				if err := recover(); err != nil {
				}
				}()
				if data[j][i] == "#" {
					result++
				}
			}()
		}
	}
	return result
}

func remove(s [][]int, i int) [][]int {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func pprint(init string, data [][]string) {
	fmt.Println(init)
	for _, i := range data {
		fmt.Println(i)
	}
}

func deepcopy(nextData [][]string, data [][]string) {
	for i := range data {
		nextData[i] = make([]string, len(data[i]))
		copy(nextData[i],data[i])
	}
}

func contain(data [][]string) bool{
	for _, i := range mem {
		if reflect.DeepEqual(i,data) {
			return true
		}
	}
	return false
}

func getAns(data [][]string) int {
	ans := 0
	for _, i := range data {
		for _, j := range i {
			if j == "#" {
				ans++
			}
		}
	}
	return ans
}
