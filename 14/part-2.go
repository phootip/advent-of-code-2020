package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example2.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n")
	fmt.Println(file)
	ans2 := part2(file)
	fmt.Println("part2: ",ans2)
}

func part2(data []string) int {
	ans := 0
	mask := ""
	mem := make(map[int]int)
	for _, line := range data {
		lineToken := strings.Split(line, " = ")
		fmt.Println(lineToken)
		if lineToken[0] == "mask" {
			mask = lineToken[1]
		} else {
			pos := []int{}
			for _, posS := range getAllPos(lineToken[0][4:len(lineToken[0])-1], mask) {
				fmt.Println(posS)
				temp, _ := strconv.ParseInt(posS, 2, 64)
				pos = append(pos, int(temp))
			}
			value, _ := strconv.Atoi(lineToken[1])

			for _, p := range pos {
				mem[p] = value
			}
			fmt.Println("pos,value:",pos,value)
			fmt.Println("---------------------------------------------------------------------------------------------------------------")
		}
	}
	for _, v := range mem {
		ans += v
	}
	return ans
}

func getAllPos(pos string, mask string) []string {
	fmt.Println("Generating All Positions....", pos, mask)
	temp, _ := strconv.Atoi(pos)
	cPos := strconv.FormatInt(int64(temp),2)
	cPos = strings.Repeat("0",len(mask)-len(cPos)) + cPos
	fmt.Println(cPos)

	cPos = applieMask(cPos,mask)
	allPos := []string{cPos}
	newAllPos := genPos(allPos)
	return newAllPos
}

func genPos(allPos []string) []string {
	newAllPos := make([]string, len(allPos))
	copy(newAllPos,allPos)
	if !strings.Contains(allPos[0],"X"){
		return allPos
	}
	pos := newAllPos[0]
	newAllPos = newAllPos[1:]
	for i := range pos {
		if pos[i] == byte('X') {
			newAllPos = append(newAllPos, pos[0:i] + "1" + pos[i+1:])
			newAllPos = append(newAllPos, pos[0:i] + "0" + pos[i+1:])
			return genPos(newAllPos)
		}
	}
	return newAllPos
}

func applieMask(pos string, mask string) string {
	buffer := make([]byte,len(pos))
	for i := range pos {
		if mask[len(mask)-i-1] == byte('X') || mask[len(mask)-i-1] == byte('1') {
			buffer[len(buffer)-i-1] = mask[len(mask)-i-1]
		}else {
			buffer[len(buffer)-i-1] = pos[len(pos)-i-1]
		}
	}
	return string(buffer)
}
