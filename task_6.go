package main

//Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"sync"
)

func Task_6_solution() {
	var wg sync.WaitGroup
	rotineClose1(&wg)
	wg.Wait()
	rotineClose2(&wg)
	wg.Wait()
	rotineClose3(&wg)
	wg.Wait()
	rotineClose4(&wg)
	wg.Wait()
}

// Когда все операции будут выполнены, горутина остановится
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
		//тут представлен цикл, который закончиться после закрытия канала и получения из него всех данных
		// переданных до закрытия
		for number := range ch {
			fmt.Printf("Получено число: %d", number)
			wg.Done()
		}
	}()
	// передача данных в канал
	ch <- 1
	// Закрытие канала
	close(ch)
}

func rotineClose3(wg *sync.WaitGroup) {
	fmt.Println("Остановка по получению сообщения из канала")
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for {
			select {
			// если сработает данные case, то горутина завершит свою работы
			// применяется например с таймерами
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
	// Создание контекста с отменой из родительского контекста
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			// как только сработает cancel() и контекст заверщается, прихожит сообщение и горутина завершается
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
