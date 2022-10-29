package main

import "fmt"

//Поменять местами два числа без создания временной переменной

func Task_13_solution() {

	a := 3
	b := 5
	swap1(&a, &b)
	swap2(&a, &b)
	swap3(&a, &b)
	swap4(&a, &b)

}

func swap1(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Printf("Swap1 result: a: %d, b: %d\n", *a, *b)
}

func swap2(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
	fmt.Printf("Swap2 result: a: %d, b: %d\n", *a, *b)

}

func swap3(a *int, b *int) {
	*a = *a * *b
	*b = *a / *b
	*a = *a / *b
	fmt.Printf("Swap3 result: a: %d, b: %d\n", *a, *b)
}

func swap4(a *int, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
	fmt.Printf("Swap4 result: a: %d, b: %d\n", *a, *b)

}
