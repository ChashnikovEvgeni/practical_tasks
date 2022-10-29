package main

//Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"sync"
)

func Task_6_solution() {
	var wg sync.WaitGroup
	//rotineClose1(&wg)
	//wg.Wait()
	//rotineClose2(&wg)
	//wg.Wait()
	rotineClose3(&wg)
	wg.Wait()
	//rotineClose4(&wg)
	//wg.Wait()
}

func rotineClose1(wg *sync.WaitGroup) {
	fmt.Println("Остановка по завершению операций")
	ch := make(chan int)
	wg.Add(1)
	go func() {
		number := <-ch
		fmt.Printf("Результат: %d", number*number)
		wg.Done()
	}()
	ch <- 1
}

func rotineClose2(wg *sync.WaitGroup) {
	fmt.Println("Остановка по закрытию канала")
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for number := range ch {
			fmt.Printf("Получено число: %d", number)
			wg.Done()
		}
	}()
	ch <- 1
	close(ch)
}

func rotineClose3(wg *sync.WaitGroup) {
	fmt.Println("Остановка по получению сообщения из канала")
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Сообщение получено, горутина остановлена")
				wg.Done()
				return
			default:
				fmt.Println("Работаем")

			}
		}
	}()
	ch <- 1
}

func rotineClose4(wg *sync.WaitGroup) {
	fmt.Println("Остановка по вызову функции отмены котнекста")
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина заврешает работу")
				wg.Done()
				return
			default:
				fmt.Println("Горутина 4 продоолжает работу")

			}

		}
	}(ctx)
	cancel()
}
