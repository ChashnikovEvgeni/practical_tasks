package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func Task_12_solution() {
	// инициализация множества и последовательности строк
	set1 := make(map[string]struct{})
	subsequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// проходимся циклом по последовательности строк и добавляем каждую в множество
	// повторные значение сохраняться не будет
	for _, st := range subsequence {
		set1[st] = struct{}{}
	}

	// Выводим полученное множество
	for key, _ := range set1 {
		fmt.Println(key)
	}
	/*
			output:
		dog
		tree
		cat
	*/

}
