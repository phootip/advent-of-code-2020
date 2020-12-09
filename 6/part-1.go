package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n\n")
	a := 0
	for _, ans := range file {
		data := strings.Replace(strings.Replace(ans,"\n", "", -1), " ", "", -1)
		a += strToSet(data)
		
	}
	fmt.Println(a)
}

func strToSet(data string) int{
	mem := make(map[rune]bool)
	for _, char := range data {
		mem[char] = true
	}
	return len(mem)
}
