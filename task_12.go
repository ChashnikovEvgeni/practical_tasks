package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func Task_12_solution() {
	set1 := make(map[string]struct{})
	subsequence := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, st := range subsequence {
		set1[st] = struct{}{}
	}

	for key, _ := range set1 {
		fmt.Println(key)
	}
}
