package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func Task_5_solution() {
	dataChan := make(chan int)
	var seconds time.Duration
	fmt.Println("Введите время таймера")
	fmt.Scanf("%d", &seconds)
	timer := time.After(time.Second * seconds)

	makeReader(dataChan)
loop:
	for i := 0; ; i++ {
		select {
		case <-timer:
			fmt.Println("Время вышло")
			break loop
		default:
			dataChan <- i

		}
	}
}

func makeReader(dCh chan int) {
	go func() {
		for number := range dCh {
			fmt.Printf("Получено число %d: ", number)
		}
	}()

}
