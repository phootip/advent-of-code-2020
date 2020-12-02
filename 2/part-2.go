package main

import (
	"fmt"
	"strings"
	// "io"
	"io/ioutil"
	"strconv"
)

func main()  {
	dat, _ := ioutil.ReadFile("./input.txt")
	// dat, _ := ioutil.ReadFile("./example.txt")
	data := strings.Split(string(dat),"\n")
	// fmt.Println(data)
	ans := 0
	for _, line := range data[:len(data)-1]  {
		// fmt.Println(line)
		// Tokenize
		lineArr := strings.Split(line," ")
		r := strings.Split(lineArr[0],"-")
		fmt.Println(r)
		low, _ := strconv.Atoi(string(r[0]))
		high, _ := strconv.Atoi(string(r[1]))
		char := string(lineArr[1][0])
		pass := string(lineArr[2])
		
		// Calculation
		if (string(pass[low-1]) == char && string(pass[high-1]) != char) || (string(pass[low-1]) != char && string(pass[high-1]) == char) {
			ans++
		}
	}
	fmt.Println(ans)
}
