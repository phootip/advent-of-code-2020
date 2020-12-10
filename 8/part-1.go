package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n")
	fmt.Println(file)

	acc,pc := 0,0
	visited := make(map[int]bool)
	for {
		if visited[pc] {
			fmt.Println("last inst: ", pc)
			fmt.Println("answer: ", acc)
			break
		}
		visited[pc] = true
		compute(file[pc], &acc, &pc)
	}
}

func compute(inst string, acc *int, pc *int) {
	opcode := strings.Split(inst, " ")[0]
	value, _  := strconv.Atoi(strings.Split(inst," ")[1])
	switch opcode {
		case "nop":
			*pc++
		case "acc":
			*acc += value
			*pc++
		case "jmp":
			*pc += value
	}
	return
}
