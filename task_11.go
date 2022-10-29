package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func Task_11_solution() {
	multiplicity1 := make(map[int]struct{})
	multiplicity2 := make(map[int]struct{})
	resultMultiplicity := make(map[int]struct{})

	set(1, 10, multiplicity1)
	set(7, 20, multiplicity2)
	fmt.Println(multiplicity1)
	fmt.Println(multiplicity2)
	intersection(multiplicity1, multiplicity2, resultMultiplicity)
	fmt.Println("Пересечение:")
	fmt.Println(resultMultiplicity)

}

func set(start int, end int, multiplicity map[int]struct{}) {
	for ; start < end; start++ {
		multiplicity[start] = struct{}{}
	}

}

func intersection(multiplicity1 map[int]struct{}, multiplicity2 map[int]struct{}, resultMultiplicity map[int]struct{}) {
	if len(multiplicity1) < len(multiplicity2) {
		multiplicity1, multiplicity2 = multiplicity2, multiplicity1
	}

	for value := range multiplicity1 {
		_, ok := multiplicity2[value]
		if ok {
			resultMultiplicity[value] = struct{}{}
		}
	}

}
