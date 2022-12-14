package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func Task_10_solution() {
	// Данные для группировки
	subsequence := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// map в котором будут храниться сгруппированные значения температуры
	groups := make(map[int][]float32)
	//вызов функции группировки
	groupBy(groups, &subsequence)

	// вывод групп
	for key, value := range groups {

		fmt.Printf("Группа: %d \n", key)
		for _, n := range value {
			fmt.Printf(" %g \n", n)
		}
	}

	/*
		output:
		Группа: 30
		 32.5
		Группа: -20
		 -25.4
		 -27
		 -21
		Группа: 10
		 13
		 19
		 15.5
		Группа: 20
		 24.5
	*/

}

// Функция группировку
func groupBy(groups map[int][]float32, subsequence *[]float32) {
	for _, number := range *subsequence {
		// находим значение целой части от деления, умножаем на 10 и получаем имя группы который принадлжит значение температуры
		group := int(number/10) * 10
		// после установки принадлежности, записываем значение в слайс данной группы
		groups[group] = append(groups[group], number)

	}
}
