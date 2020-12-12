package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type inst struct{
	ops	byte
	value int
}

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]),"\n")
	fmt.Println(file)
	data := tokenize(file)
	fmt.Println(data)
	ans1 := part1(data)
	fmt.Println("part 1: ", ans1)
}

func part1(data []inst) int{
	pos := [2]int{0,0}
	dir := [2]int{10,-1}
	for i, inst := range data {
		fmt.Println(i,inst)
		switch inst.ops {
		case byte('N'):
			dir[1] -= inst.value
		case byte('S'):
			dir[1] += inst.value
		case byte('E'):
			dir[0] += inst.value
		case byte('W'):
			dir[0] -= inst.value
		case byte('F'):
			pos[0] += dir[0]*inst.value
			pos[1] += dir[1]*inst.value
		case byte('R'),
		byte('L'):
			switch inst.value {
			case 90:
				dir = turn(dir, inst.ops)
			case 180:
				dir[0] = -dir[0]
				dir[1] = -dir[1]
			case 270:
				if inst.ops == byte('R') { dir = turn(dir, byte('L')) }
				if inst.ops == byte('L') { dir = turn(dir, byte('R')) }
			}
		default:
			panic("Unknow Ops")
		}
		fmt.Println("Current dir:", dir)
		fmt.Println("Current Pos:", pos)
	}

	fmt.Println("Final Pos: ", pos)
	return int(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}

// 			-4,-10
// -10,4		10,-4
// 			4,10
func turn(dir [2]int, ops byte) [2]int{
	if ops == byte('R') {
		dir[0], dir[1] = -dir[1], dir[0]
	} else if ops == byte('L') {
		dir[0], dir[1] = dir[1], -dir[0]
	}
	return dir
}

func tokenize(file []string) []inst{
	result := []inst{}
	for _, line := range file {
		temp, _ := strconv.Atoi(line[1:])
		result = append(result, inst{line[0], temp})
	}
	return result
}

// 583 too low
