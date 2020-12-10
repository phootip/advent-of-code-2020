package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	preamble := 25 // target := 26134589
	// rawDat, _ := ioutil.ReadFile("example.txt")
	// preamble := 5 // target := 127
	file := make([]int, len(strings.Split(string(rawDat), "\n"))-1)
	for i, j := range strings.Split(string(rawDat[:len(rawDat)-1]), "\n"){
		file[i], _ = strconv.Atoi(j)
	}
	fmt.Println(file)
	i := preamble
	for {
		if isValid(file[i], file[i-preamble:i]){
			i++
			continue
		}
		break
	}
	fmt.Println(file[i])
	fmt.Println("Starting Part2...")
	ans2 := part2(file[i], file)
	fmt.Println(ans2)
}

func part2(target int, data []int) int{
	fmt.Println("Target: ", target)
	for i := 0; i< len(data); i++ {
		for j := i+2; j < len(data); j++ {
			fmt.Println(sum(data[i:j]))
			if sum(data[i:j]) == target {
				fmt.Println("Found Answer")
				return min(data[i:j]) + max(data[i:j])
			}else if sum(data[i:j]) > target {
				break
			}
		}
	}
	return 0
}

func sum(arr []int) int{
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}

func min(arr []int) int{
	cMin := arr[0]
	for _, i := range arr {
		if i < cMin {
			cMin = i
		}
	}
	return cMin
}

func max(arr []int) int{
	cMax := arr[0]
	for _, i := range arr {
		if i > cMax {
			cMax = i
		}
	}
	return cMax
}

func isValid(num int, list []int) bool{
	fmt.Println("Start isValid...")
	fmt.Println(num, list)
	for i := 0; i < len(list); i++ {
		first := list[i]
		for j:=i; j < len(list); j++ {
			second := list[j]
			if first + second == num { 
				return true
			}
		}
	}
	return false
}
