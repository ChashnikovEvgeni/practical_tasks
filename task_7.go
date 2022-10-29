package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

func Task_7_solution() {

	var wg sync.WaitGroup
	var mutex sync.Mutex
	people := map[string]int{
		"Vlad":  3711,
		"Bob":   2138,
		"Slava": 1908,
		"Pavel": 912,
	}
	names := make([]string, len(people))

	i := 0
	for k := range people {
		names[i] = k
		i++
		fmt.Println(k)
	}

	wg.Add(2)
	go payDay(&wg, &mutex, people, names, 2000)
	go payDay(&wg, &mutex, people, names, 4000)
	wg.Wait()

	for name, value := range people {
		fmt.Printf("Имя: %s, Баланас: %d \n", name, value)
	}

}

func payDay(wg *sync.WaitGroup, mutex *sync.Mutex, people map[string]int, names []string, money int) {
	for _, name := range names {
		mutex.Lock()
		people[name] = people[name] + money
		mutex.Unlock()
	}
	wg.Done()

}
