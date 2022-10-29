package main

import (
	"fmt"
)

// Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй —
// результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func Task_9_solution() {
	numbers := []int{2, 3, 4, 5, 6}
	channel := set_data(numbers)
	ch := pow_data(channel)
	for number := range ch {
		fmt.Println(number)
	}

}

func set_data(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}

		close(out)
	}()
	return out
}

func pow_data(dataCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range dataCh {
			out <- n * n
		}
		close(out)
	}()
	return out
}
