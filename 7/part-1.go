package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Day7")
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
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
	fmt.Println(part1(rules,[]string{"shinygold"}, make(map[string]bool)))
}

func part1(rules map[string][][]string, checklist []string, counted map[string]bool) int {
	// fmt.Println(checklist)
	result := 0
	newChecklist := []string{}
	for left, right := range rules {
		for _, bag := range right {
			if contain(checklist,bag[1]) && !counted[bag[1]] {
				result++
				newChecklist = append(newChecklist, left)
				break
			}
		}
	}
	for _, i := range checklist {
		counted[i] = true
	}
	// fmt.Println("result: ", result)
	// fmt.Println("counted: ", counted)
	if len(newChecklist) != 0 {
		fmt.Println("real ans?:", len(counted))
		return result + part1(rules, newChecklist, counted)
	}
	fmt.Println("real ans?:", len(counted)-1) // Yeah this one is real
	return 0
}

func contain(strArr []string, str string) bool{
	for _, s := range strArr{
		if str == s{
			return true
		}
	}
	return false
}

// 190 too high 
// 161 too high
// 134 too low
// 144 correct
