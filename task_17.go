package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка

func Task_17_solution() {
	//бинарный поиск осуществляется по отсортированному массиву
	sortArray := Task_16_solution()
	//[-64 -23 -1 1 12 22 23 23 45 53 56 76 76 86]

	var element int
	fmt.Println("Введите значение элемента для бинарного поиска")
	// получение длительность таймера по вводу пользователя
	fmt.Scanf("%d", &element)
	found, i := binarySearch(element, sortArray)
	if found {
		fmt.Printf("Найдено, индекс: %d", i)
	} else {
		fmt.Println("Объект не найден")
	}

}

func binarySearch(element int, sortArray []int) (bool, int) {
	// Старотовые значения левой и правой границы
	leftBorder := 0
	rightBorder := len(sortArray) - 1
	// искать будем до тех пор, пока значение границ не приобретут одно и тоже значение
	for leftBorder <= rightBorder {
		// находим условный центр
		center := (leftBorder + rightBorder) / 2
		// если значение элемента который мы ищем больше значени элемента по индексу центра, то мы увеличиваем индекс центр
		// иначе уменьшаем
		if sortArray[center] < element {
			leftBorder = center + 1
		} else {
			rightBorder = center - 1
		}
	}

	// если пройдя весь срез и на границы осуществив попытку поиска мы не найдём элемент
	// то он будет признан отсутствующим
	if leftBorder == len(sortArray) || sortArray[leftBorder] != element {
		return false, 0
	}
	return true, leftBorder
}
