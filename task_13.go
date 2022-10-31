package main

import (
	"fmt"
)

//Поменять местами два числа без создания временной переменной

func Task_13_solution() {
	//инициализация переменных
	a := 3
	b := 5
	swap1(&a, &b)
	swap2(&a, &b)
	swap3(&a, &b)
	swap4(&a, &b)

}

// обмен местами с помощью =
func swap1(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Printf("Swap1 result: a: %d, b: %d\n", *a, *b)
}

// обмен местами с помошью 1-ого сложения и 2-ух вычитаний
func swap2(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
	fmt.Printf("Swap2 result: a: %d, b: %d\n", *a, *b)

}

// обмен местами с помощью 1-ого умножения и 2-ух делений
func swap3(a *int, b *int) {
	*a = *a * *b
	*b = *a / *b
	*a = *a / *b
	fmt.Printf("Swap3 result: a: %d, b: %d\n", *a, *b)
}

func swap4(a *int, b *int) {
	// если значение битов одинаков то 0, если разно то 1
	*a ^= *b
	*b ^= *a
	*a ^= *b
	fmt.Printf("Swap4 result: a: %d, b: %d\n", *a, *b)

}
