package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false */

func Task_26_solution() {
	st := "Wwe"
	fmt.Printf("Строка содержит только уникальные значения?: %t ", Checkunique(st))

}

// Проверка уникальности
func Checkunique(s string) bool {
	// переводим строку в нижний регистр, так проверка станет регистронезависимой
	s = strings.ToLower(s)
	//инициализируем множества с помощью которого будем проводить проверку уникальности
	checklist := make(map[rune]struct{})
	// переводим строку в набор рун
	runes := []rune(s)
	// если данная рунна не находиться в множестве - добавляем
	// если она будет найдена - проверка уникальности не пройдена
	for _, rune := range runes {
		_, ok := checklist[rune]
		if ok {
			return false
		}
		// добавляем руну в множество
		checklist[rune] = struct{}{}

	}
	return true
}
