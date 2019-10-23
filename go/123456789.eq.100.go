package main

import "fmt"
import "strings"

// import "time"

var nums = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var max = [10]int{123456789, 23456789, 3456789, 456789, 56789, 6789, 789, 89, 9}

// 123 - (456789)

func find(prefix string, index, expect int) {
	fmt.Println("find")
	if expect > 0 {
		if max[index] < expect {
			return
		}
	}
	if expect < 0 {
		if -max[index] > expect {
			return
		}
	}
	if index == 9 && expect == 0 {
		if prefix[0] == '+' {
			fmt.Println(prefix[1:])
		}
		return
	}
	var value int
	// +
	value = 0
	for i := index; i < 9; i++ {
		value = value*10 + i + 1
		find(prefix+"+"+strings.Join(nums[index:i+1], ""), i+1, expect-value)
	}
	// -
	value = 0
	for i := index; i < 9; i++ {
		value = value*10 - i - 1
		find(prefix+"-"+strings.Join(nums[index:i+1], ""), i+1, expect-value)
	}
}

// 1+/-(...) = 100
func main() {
	find("", 0, 100)
}
