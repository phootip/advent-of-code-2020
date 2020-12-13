package main

import (
	"fmt"
	"io/ioutil"
	// "math"
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

func part2(data []string) int{
	buses := []int{}
	for _, i := range strings.Split(data[1],",") {
		if i == "x" { 
			buses = append(buses, 0)
		} else {
			temp, _ := strconv.Atoi(i)
			buses = append(buses, temp)
		}
	}
	fmt.Println(buses)
	c,k := 0, buses[0]
	for i, bus := range buses {
		if bus != 0 {
			c,k = getCK(k,bus, c, i)
			fmt.Printf("%v + %vk\n",c,k)
			fmt.Println("-------------------------------")
		}
	}
	return c
	// fmt.Println((77+91)%7)
	// fmt.Println((77+91+1)%13)
	// fmt.Println((350+5369)%7)
	// fmt.Println((350+5369+1)%13)
	// fmt.Println((350+5369+4)%59)
	return 0
}

func getCK(a int, b int, c1 int, c2 int) (int,int){
	fmt.Printf("%v+%vN | %v+%vM: \n", c1,a,c2,b)
	fmt.Println("LCM:", lcm(a,b))
	fmt.Println("GCD:", gcd(a,b))
	n := 0
	for {
		right := (c1+a*n+c2)
		if right % b == 0 {
			m := right/b
			fmt.Println("The number n,m: ", n, m)
			fmt.Println("Time: ", n*a)
			break
		}
		n++
	}
	return a*n+c1, lcm(a,b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
