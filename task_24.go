package main

import (
	"fmt"
	"math"
)

/*Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.*/

// Структура для точки с инкапсулированными параметрами x,y и конструктором
type Point struct {
	x, y float64
}

// метод конструктор
func (p *Point) Init(x, y float64) {
	p.x = x
	p.y = y
}

// метод для расчёта дистанции
func distance(p1, p2 *Point) float64 {

	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))

}

func Task_24_solution() {
	var P1 Point
	var P2 Point

	// Вызов конструкторов для точек
	P1.Init(5, 6)
	P2.Init(20, 10)

	fmt.Printf("Дистанция между точками: %g", distance(&P1, &P2))
	/*
		output
		Дистанция между точками: 15.524174696260024
	*/

}
