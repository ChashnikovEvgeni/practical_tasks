package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func WordReverse(words []string) string {
	n := len(words)
	// парными заменами переворачиваем слова в строке
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}

	return strings.Join(words, " ")
}

// функция чтобы парсить строку по пробелам " "
// возвращает слайс с отдельными словами
func ParserFiels(s string) []string {
	words := strings.Fields(s)
	return words

}

func Task_20_solution() {
	st := "dog sun snow kek mem resulter"
	fmt.Println(WordReverse(ParserFiels(st)))
}
