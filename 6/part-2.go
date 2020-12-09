package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	rawDat, _ := ioutil.ReadFile("input.txt")
	// rawDat, _ := ioutil.ReadFile("example.txt")
	file := strings.Split(string(rawDat[:len(rawDat)-1]), "\n\n")
	a := 0
	for _, ans := range file {
		data := strings.Split(strings.Replace(ans, " ", "", -1), "\n")
		fmt.Println(data)
		a += inclusion(data)
		fmt.Println("end loop:",a)
	}
	fmt.Println(a)
}

func inclusion(data []string) int{
	mem := strings.Split(data[0],"")
	for _, char := range data[1:] {
		mem = intersect(mem,strings.Split(char,""))
	}
	return len(mem)
}

func intersect(str []string, str2 []string) []string{
	fmt.Println(str, str2)
	temp := make(map[string]bool)
	result := []string{}
	for _, i := range str{
		temp[i] = true
	}
	for _, i := range str2{
		if temp[i] {
			result = append(result, i)
		}
	}
	return result
}
