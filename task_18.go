package main

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// структура счётчик
type Counter struct {
	// поле для инкремента
	count int
	// mutex для избегания некорректной работы счётчика
	mutex sync.Mutex
}

// иной вариант структуры счётчика
type CounterAtomic struct {
	count int32
}

func Task_18_solution() {
	// инициализируем переменные
	var amount int
	var iter int
	var counter Counter
	var wg sync.WaitGroup
	var counterAtomic CounterAtomic

	//Вводим данные
	fmt.Println("Введите количество инкрементеров")
	fmt.Scanf("%d\n", &amount)
	fmt.Println("Введите количество операция для каждого инкреметера")
	fmt.Scanf("%d\n", &iter)

	atiomicAmount := amount

	// двойной размер группы т.к. будут работать два счётчика
	wg.Add(2 * amount)
	for amount > 0 {
		// создаём для каждой функции инкримента отдельную горутину
		go increment(&counter, &wg, iter)
		amount--
	}

	for atiomicAmount > 0 {
		// создаём для каждой функции инкримента отдельную горутину
		go incrementAtomic(&counterAtomic, &wg, iter)
		atiomicAmount--
	}

	// ожидаем завершения всех горутин
	wg.Wait()
	fmt.Printf("Итоговое значение increment: %d \n", counter.count)
	fmt.Printf("Итоговое значение incrementAtomic: %d", counterAtomic.count)

}

// функция для инкримента
func increment(counter *Counter, wg *sync.WaitGroup, iter int) {
	for iter > 0 {
		// блокируем мьютекс на время записи, так одни данные не будут переписываться поверх других
		counter.mutex.Lock()
		counter.count++
		// разблокируем мьютекс
		counter.mutex.Unlock()
		iter--
	}
	// сигнализируем группе о выполнеии
	wg.Done()

}

// атомарное инкременьтрование
// мьютекс не требуется т.к. реализуются на аппаратном уровне
func incrementAtomic(counter *CounterAtomic, wg *sync.WaitGroup, iter int) {
	for iter > 0 {
		atomic.AddInt32(&counter.count, 1)
		iter--
	}
	wg.Done()

}
