package main

import (
	"fmt"
)

func main() {
	// data := []int{0,3,6} // 175594
	// data := []int{1,3,2} // 1
	// data := []int{2,1,3} // 10
	// data := []int{1,2,3} // 27
	// data := []int{2,3,1} // 78
	// data := []int{3,2,1} // 438
	// data := []int{3,1,2} // 1836
	data := []int{0,14,1,3,7,9} // 763 
	fmt.Println("Part1 :", part1(data))
}

func part1(data []int) int {
	mem := make(map[int][]int)
	for i, num := range data {
		mem[num] = append(mem[num], i)
	}
	lastNum := data[len(data)-1]
	nextNum := 0
	count := len(data)
	for count < 30000000 {
		nextNum = getNextNum(mem, lastNum)
		mem[nextNum] = append(mem[nextNum], count)
		lastNum = nextNum
		count++
		fmt.Println(count)
	}
	
	return nextNum
}

func getNextNum(mem map[int][]int, lastNum int) int{
	temp := mem[lastNum]
	if len(temp) == 1 {
		return 0
	}
	return temp[len(temp)-1]-temp[len(temp)-2]
}
