package main

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int
	mutex sync.Mutex
}

type CounterAtomic struct {
	count int32
}

func Task_18_solution() {
	var amount int
	var iter int
	var counter Counter
	var wg sync.WaitGroup
	var counterAtomic CounterAtomic
	fmt.Println("Введите количество инкрементеров")
	fmt.Scanf("%d\n", &amount)
	fmt.Println("Введите количество операция для каждого инкреметера")
	fmt.Scanf("%d\n", &iter)

	atiomicAmount := amount

	wg.Add(2 * amount)
	for amount > 0 {
		go increment(&counter, &wg, iter)
		amount--
	}

	for atiomicAmount > 0 {
		go incrementAtomic(&counterAtomic, &wg, iter)
		atiomicAmount--
	}

	wg.Wait()
	fmt.Printf("Итоговое значение increment: %d \n", counter.count)
	fmt.Printf("Итоговое значение incrementAtomic: %d", counterAtomic.count)

}

func increment(counter *Counter, wg *sync.WaitGroup, iter int) {
	for iter > 0 {
		counter.mutex.Lock()
		counter.count++
		counter.mutex.Unlock()
		iter--
	}
	wg.Done()

}

func incrementAtomic(counter *CounterAtomic, wg *sync.WaitGroup, iter int) {
	for iter > 0 {
		atomic.AddInt32(&counter.count, 1)
		iter--
	}
	wg.Done()

}
