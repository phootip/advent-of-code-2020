package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	preamble := 25
	// rawDat, _ := ioutil.ReadFile("example.txt")
	// preamble := 5
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
