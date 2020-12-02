package main

import (
	"fmt"
	"strings"
	// "io"
	"io/ioutil"
	"strconv"
)

func main() {
	// fmt.Println("Hello Go")
	// dat, err := ioutil.ReadFile("./example.txt")
	dat, err := ioutil.ReadFile("./input.txt")
	fmt.Println(err)
	// fmt.Println(string(dat))
	data := strings.Fields(string(dat))
	fmt.Println(data, len(data))
	for i := 0; i < len(data); i++ {
		first, _ := strconv.Atoi(data[i])
		for j := i; j < len(data); j++ {
			second, _ := strconv.Atoi(data[j])
			if first + second < 2020 { 
				for k := j; k < len(data); k++{
					third, _ := strconv.Atoi(data[k])
					if first+second+third == 2020 {
						fmt.Println("ans=",first,second,third,first*second*third)
					}
				}
			}
		}
	}
}
