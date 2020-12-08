package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	// "strings"
	"regexp"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	// dat, _ := ioutil.ReadFile("example2.txt")
	// dat, _ := ioutil.ReadFile("example3.txt")
	dat = dat[:len(dat)-1]
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
		for k, v := range mem {
			if IsValidField(k, v) {
				fmt.Println(k,v)
				category[k] = true
			}
		}
		fmt.Println(category)
		if len(category) == 7 {
			fmt.Println("valid!!")
			ans++
		}
		fmt.Println()
	}
	fmt.Println(ans)
}

func IsValidField(field string, value string) bool {
	switch field {
	case "byr":
		// result, _ := regexp.MatchString(`^[1920-2002]$`, value)
		temp, _ := strconv.Atoi(value)
		result := 1920 <= temp && temp <= 2002 && len(value) == 4
		return result
	case "iyr":
		temp, _ := strconv.Atoi(value)
		result := 2010 <= temp && temp <= 2020 && len(value) == 4
		return result
	case "eyr":
		temp, _ := strconv.Atoi(value)
		result := 2020 <= temp && temp <= 2030 && len(value) == 4
		return result
	case "hgt":
		isCM, _ := regexp.MatchString(`^\d*cm$`, value)
		if isCM {
			height, _ := strconv.Atoi(value[:len(value)-2])
			if 150 <= height && height <= 193 {
				return true
			}
		}
		isIN, _ := regexp.MatchString(`^\d*in$`, value)
		if isIN {
			height, _ := strconv.Atoi(value[:len(value)-2])
			if 59 <= height && height <= 76 {
				return true
			}
		}
		return false
	case "hcl":
		result, _ := regexp.MatchString(`^#([0-9]|[a-f]){6}$`, value)
		return result
	case "ecl":
		result, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, value)
		return result
	case "pid":
		result, _ := regexp.MatchString(`^[0-9]{9}$`, value)
		return result
	}
	return false
}

// 170 too high
// 168 too high
