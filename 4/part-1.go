package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	// dat, _ := ioutil.ReadFile("example.txt")
	dat = dat[:len(dat)-2]
	data := strings.Split(string(dat), "\n\n")
	// data = data[:len(data)-1]
	// fmt.Println(dat)
	// fmt.Println(strings.Join(data,","))
	ans := 0
	for _, passport := range data {
		// fmt.Println(passport)
		tokenized := strings.Split(strings.Replace(passport,"\n"," ",-1)," ")
		mem := make(map[string]string)
		for _, rawData := range tokenized {
			data := strings.Split(rawData,":")
			mem[data[0]] = data[1]
		}
		fmt.Println(mem)

		category := make(map[string]bool)
		for k := range mem {
			if IsValidCategory(k) {
				category[k] = true
			}
		}
		fmt.Println(category)
		if len(category) == 7 {
			ans++
		}
		fmt.Println()
	}
	fmt.Println(ans)
}

func IsValidCategory(category string) bool {
	switch category {
	case
			"byr",
			"iyr",
			"eyr",
			"hgt",
			"hcl",
			"ecl",
			"pid":
			// "cid":
			return true
	}
	return false
}

// 207 too low
