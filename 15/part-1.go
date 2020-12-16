package main

import (
	"fmt"
)

func main() {
	data := []int{0,3,6}
	fmt.Println(data)
	fmt.Println("Part1 :", part1(data))
}

func part1(data []int) int {
	mem := make(map[int][]int)
	for i, num := range data {
		mem[num] = append(mem[num], i)
	}

	
	fmt.Println(mem)
	return 0
}

func getNextNum(mem map[int][]int, lastNum int) int{
	temp := mem[lastNum]
		result := 
	fmt.Println(temp)
}
