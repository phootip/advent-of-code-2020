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
	ans2 := part2(file)
	fmt.Println("part2: ",ans2)
}

func part2(data []string) int64 {
	ans := int64(0)
	mask := ""
	mem := make(map[int]int64)
	for _, line := range data {
		temp := strings.Split(line, " = ")
		fmt.Println(temp)
		if temp[0] == "mask" {
			mask = temp[1]
		} else {
			pos, _ := strconv.Atoi(temp[0][4:len(temp[0])-1])
			temp2, _ := strconv.Atoi(temp[1])
			value := strconv.FormatInt(int64(temp2),2)
			value = strings.Repeat("0",len(mask)-len(value)) + value
			fmt.Println(pos,value,mask)
			buffer := make([]byte,len(value))
			for i := range value {
				if mask[len(mask)-i-1] != byte('X') {
					buffer[len(buffer)-i-1] = mask[len(mask)-i-1]
				}else {
					buffer[len(buffer)-i-1] = value[len(value)-i-1]
				}
			}
			fmt.Println(string(buffer))
			mem[pos], _ = strconv.ParseInt(string(buffer), 2, 64)
			fmt.Println(mem)
		}
	}
	for _, v := range mem {
		ans += v
	}
	return ans
}
