package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка

func Task_17_solution() {

	sortArray := Task_16_solution()
	b, i := binarySearch(-1, sortArray)
	fmt.Printf("Найдено?: %t \n количество итераций: %d", b, i)

}

func binarySearch(needleElem int, sortData []int) (bool, int) {
	low := 0
	high := len(sortData) - 1
	for low <= high {
		center := (low + high) / 2
		if sortData[center] < needleElem {
			low = center + 1
		} else {
			high = center - 1
		}
	}

	if low == len(sortData) || sortData[low] != needleElem {
		return false, 0
	}
	return true, low + 1
}
