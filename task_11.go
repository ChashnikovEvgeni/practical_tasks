package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func Task_11_solution() {
	// инициализация множеств
	multiplicity1 := make(map[int]struct{})
	multiplicity2 := make(map[int]struct{})
	// инициализация множества для хранения пересечения
	resultMultiplicity := make(map[int]struct{})

	//задаём множества так, чтобы у них были некоторые одинаковые значения
	set(1, 10, multiplicity1)
	set(7, 20, multiplicity2)
	fmt.Println(multiplicity1)
	fmt.Println(multiplicity2)
	intersection(multiplicity1, multiplicity2, resultMultiplicity)
	fmt.Println("Пересечение:")
	fmt.Println(resultMultiplicity)

}

// Функция для задачи множества
func set(start int, end int, multiplicity map[int]struct{}) {
	for ; start < end; start++ {
		multiplicity[start] = struct{}{}
	}

}

func intersection(multiplicity1 map[int]struct{}, multiplicity2 map[int]struct{}, resultMultiplicity map[int]struct{}) {
	// выбираем самое длинное можество
	if len(multiplicity1) < len(multiplicity2) {
		multiplicity1, multiplicity2 = multiplicity2, multiplicity1
	}

	// проходимся по самому длинному множеству
	for value := range multiplicity1 {
		// если число, хранящиеся в 1-ом множестве было найдено (found == true) во 2-ом
		// оно добавляется в итоговое множество
		_, found := multiplicity2[value]
		if found {
			resultMultiplicity[value] = struct{}{}
		}
	}

}
