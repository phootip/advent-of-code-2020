package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sort"
)

func main() {
	rawFile, _ := ioutil.ReadFile("input.txt")
	// rawFile, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawFile[:len(rawFile)-1]),"\n")
	fmt.Println(file)
	var mem []int
	for _, line := range file {
		fmt.Println(rawToBi(line))
		result, _ := strconv.ParseInt(rawToBi(line),2,64)
		fmt.Println(result)
		mem = append(mem, int(result))
	}
	sort.Ints(mem)
	fmt.Println(mem)

	for i, _ := range mem {
		if i != len(mem)-1 && mem[i+1] - mem[i] != 1 {
			fmt.Println(mem[i]+1)
		}
	}
}

func rawToBi(raw string) string{
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(raw, "B", "1", -1), "F", "0", -1), "R", "1", -1), "L", "0", -1)
}

// 1023 too high
