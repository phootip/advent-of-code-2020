package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var mem = make(map[int]int)

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
	fmt.Println(part2(file,0,0))
}

func part2(data []int, adapter int, index int) int {
	// fmt.Println("mem default:",mem[adapter])
	fmt.Println(adapter)
	if mem[adapter] != 0 {
		return mem[adapter]
	} else if adapter == data[len(data)-1]{
		return 1
	}
	result := 0
	for i := index+1; i <= index+3 && i < len(data); i++ {
		if data[i] - adapter <=3 {
			result += part2(data, data[i], i)
		} else {
			break
		}
	}
	fmt.Println("result: ", result)
	mem[adapter] = result
	return result
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
