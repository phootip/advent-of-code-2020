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
		count := 0
		for _, i := range pass{
			if string(i) == char {
				count++
			}
		}
		if low <= count && count <= high {
			ans++
		}
	}
	fmt.Println(ans)
}

// 269 wrong
