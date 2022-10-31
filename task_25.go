package main

import (
	"fmt"
	"log"
	"time"
)

//Реализовать собственную функцию sleep.

func Ownsleep2(sec int) {
	log.Println("OwnSleep2 отсчёт начал")
	t := time.Now()
	// функция sleep опирается на значение времени прошедшие с момента засечения
	// как только оно будет равно длительности, sleep завершиться
	for time.Now().Sub(t).Seconds() < float64(sec) {
	}
	log.Println("OwnSleep2 таймер вышел")
}

func Ownsleep1(sec int) {
	log.Println("OwnSleep1 отсчёт начал")
	// как только таймер подаст сигнал, sleep завершится
	<-time.After(time.Duration(sec) * time.Second)
	log.Println("OwnSleep1 таймер вышел")
}

func Task_25_solution() {

	var timer int
	fmt.Println("Введите время ожидания")
	fmt.Scanf("%d\n", &timer)
	Ownsleep1(timer)
	Ownsleep2(timer)
}
