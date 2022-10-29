package main

import (
	"fmt"
	"sync"
	"time"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func Task_3_solution() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	channel := make(chan int, 5)
	start := time.Now()

	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			pow := n * n
			channel <- pow
		}(number)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			sum = sum + <-channel
		}
	}()

	wg.Wait()
	close(channel)

	fmt.Printf("Сумма квадратов равна: %d \n", sum)
	finalTime := time.Now().Sub(start).Seconds()
	fmt.Printf("Время выполнения: %g \n", finalTime)
}
