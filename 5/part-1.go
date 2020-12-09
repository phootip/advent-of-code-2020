package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"math"
)

func main() {
	rawFile, _ := ioutil.ReadFile("input.txt")
	// rawFile, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawFile[:len(rawFile)-1]),"\n")
	fmt.Println(file)
	ans := 0
	for _, line := range file {
		fmt.Println(rawToBi(line))
		result, _ := strconv.ParseInt(rawToBi(line),2,64)
		fmt.Println(result)
		ans = int(math.Max(float64(ans),float64(result)))
	}
	fmt.Println(ans)
}

func rawToBi(raw string) string{
	return strings.Replace(strings.Replace(strings.Replace(strings.Replace(raw, "B", "1", -1), "F", "0", -1), "R", "1", -1), "L", "0", -1)
}

// 1023 too high
