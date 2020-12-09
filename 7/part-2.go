package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day7")
	// rawDat, _ := ioutil.ReadFile("input.txt")
	rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]),"\n")
	rules := make(map[string][][]string)
	for _, rawRule := range file {
		rawRule = rawRule[:len(rawRule)-1]
		rule := strings.Split(rawRule, " contain ")
		left := strings.Split(rule[0], " ")[0] + strings.Split(rule[0], " ")[1]
		if rule[1] == "no other bags" {
			continue
		} else {
			temp := strings.Split(rule[1], ", ")
			right := [][]string{}
			for _, i := range temp {
				temp2 := strings.Split(i, " ")
				right = append(right,[]string{temp2[0],temp2[1]+temp2[2]})
			}
			// fmt.Println(right[0][1])
			rules[left] = right
		}
	}
	fmt.Println(rules)
	fmt.Println(part2(rules,"shinygold"))
}

func part2(rules map[string][][]string,bag string) int {
	ans := 0
	needed := rules[bag]
	fmt.Println(needed)
	return ans
}

func contain(strArr []string, str string) bool{
	for _, s := range strArr{
		if str == s{
			return true
		}
	}
	return false
}
