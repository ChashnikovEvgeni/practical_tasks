package main

import (
	"fmt"
)

// Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй —
// результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func Task_9_solution() {
	// инициализация слайса с числами
	numbers := []int{2, 3, 4, 5, 6}
	// полечения канала для получения из него данных
	channel := set_data(numbers)
	// получения канала из которого будут получаться числа
	ch := pow_data(channel)
	// вывод данных из канала по их готовности
	for number := range ch {
		// fmt.Println использует для вывода stdout
		fmt.Println(number)
	}

}

// Функция получается слайс чисел и в отдельной горутине передаёт их в канал, возращает однонаправленный канал из которого возможно только получить данные
// По заверщению передачи, закрывает канал
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

// Получает однонаправленный канал, получаеет из него данные, возводит в квадрат и предаёт в другой канал
// в главное горутине из него получат данные
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
