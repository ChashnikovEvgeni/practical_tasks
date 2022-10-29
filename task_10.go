package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func Task_10_solution() {
	groups := make(map[int][]float32)
	groupBy(groups)
	for key, value := range groups {

		fmt.Printf("Ключ: %d \n", key)
		for _, n := range value {
			fmt.Printf(" %g \n", n)
		}
	}

}

func groupBy(groups map[int][]float32) {
	subsequence := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, number := range subsequence {
		group := int(number/10) * 10
		groups[group] = append(groups[group], number)

	}
}
