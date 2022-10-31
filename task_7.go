package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

func Task_7_solution() {

	var wg sync.WaitGroup
	// инициализация Mutex
	var mutex sync.Mutex
	// Инициализация структуры с банковским балансом указанных людей
	people := map[string]int{
		"Vlad":  3711,
		"Bob":   2138,
		"Slava": 1908,
		"Pavel": 912,
	}

	wg.Add(2)
	// создание горутин начисляющих средства на счета людей
	// предполагается, что счёт каждого человека увеличиться на 6000 единиц
	go payDay(&wg, &mutex, people, 2000)
	go payDay(&wg, &mutex, people, 4000)
	// ожидание выполнение всех задач
	wg.Wait()

	// вывод итоговых данных
	for name, value := range people {
		fmt.Printf("Имя: %s, Баланас: %d \n", name, value)
	}

	/*
			output:
		Имя: Vlad, Баланас: 9711
		Имя: Bob, Баланас: 8138
		Имя: Slava, Баланас: 7908
		Имя: Pavel, Баланас: 6912


	*/

}

// функция начисления средства
func payDay(wg *sync.WaitGroup, mutex *sync.Mutex, people map[string]int, money int) {
	// проходимся циклом по всем людям
	for name, _ := range people {
		// перед началом записи необходимо заблокировать мютекс
		// так две горутины не перепишут данные поверх друг друга и оба числа прибавятся к текущему балансу
		mutex.Lock()
		people[name] = people[name] + money
		// разблокировка после окончания записи
		mutex.Unlock()
	}
	// уменьшение счётчика активных элементов waitgroup
	wg.Done()

}
