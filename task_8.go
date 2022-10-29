package main

import (
	"fmt"
	"strconv"
)

func On(n int64, i int) int64 {
	return n | 1<<(i-1)
}

func Off(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func Task_8_solution() {
	var number int64
	number = 9223372036854775806

	on := On(number, 1)
	off := Off(number, 2)

	fmt.Printf("%s\n", strconv.FormatInt(number, 2))
	fmt.Printf("%s\n", strconv.FormatInt(on, 2))
	fmt.Printf("%s\n", strconv.FormatInt(off, 2))
}
