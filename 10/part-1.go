package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	// rawDat, _ := ioutil.ReadFile("example2.txt")
	file := make([]int, len(strings.Split(string(rawDat), "\n"))-1)
	for i, j := range strings.Split(string(rawDat[:len(rawDat)-1]), "\n"){
		file[i], _ = strconv.Atoi(j)
	}
	sort.Ints(file)
	file = append(file[:1], file...)
	file[0] = 0
	file = append(file, file[len(file)-1] + 3)
	fmt.Println(file)
	fmt.Println(part1(file))
}

func part1(data []int) int {
	jot1, jot3 := 0,0
	
	for i, x := range data{
		if i == len(data)-1 {
			fmt.Println(jot1, jot3)
			return jot1*jot3
		} else {
			if data[i+1] - x == 1 {
				jot1++
			}else if data[i+1] - x == 3 {
				jot3++
			}
		}
	}
	return 0
}
