package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n")
	iCorrupted := 0
	acc, pc := 0,0
	for pc != len(file) {
		fmt.Println(file)
		mem := make([]string,len(file))
		copy(mem,file)
		mem, iCorrupted = changeMem(mem, iCorrupted)
		acc, pc = 0,0
		visited := make(map[int]bool)
		for pc != len(file) {
			if visited[pc] {
				fmt.Println("last inst: ", pc)
				fmt.Println("answer: ", acc)
				break
			}
			visited[pc] = true
			compute(mem[pc], &acc, &pc)
		}
	}
	fmt.Println("part2:",acc)
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

func changeMem(mem []string, i int) ([]string, int) {
	fmt.Println("Changing mem....")
	for ; mem[i][:3] == "acc"; i++ {}
	switch mem[i][:3] {
		case "nop":
			mem[i] = strings.Replace(mem[i],"nop","jmp",1)
		case "jmp":
			mem[i] = strings.Replace(mem[i],"jmp","nop",1)
		default:
			errors.New("instruction is wrong")
	}
	// fmt.Println(mem[i][:3])
	// fmt.Println("mem changed:", mem)
	return mem, i+1
}
