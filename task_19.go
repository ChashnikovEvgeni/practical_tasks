package main

import "fmt"

/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

func ReverseString(s string) string {
	// переводим string  в слайс рун
	r := []rune(s)
	n := len(r)
	// Ввиду парных замен, количество итераций меньше в 2 раза длины
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return string(r)
}

func Task_19_solution() {
	// инициализируем переменные
	var s string
	var resultS string
	// получаем строку для переворота
	fmt.Println("Введите строку")
	fmt.Scanf("%s", &s)
	// вызываем функция для переворота и выводим результат
	resultS = ReverseString(s)
	fmt.Println(resultS)
}
