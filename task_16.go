package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func Task_16_solution() []int {

	array := []int{22, 56, 23, 76, -23, 45, 1, 12, 53, 86, -64, -1, 76, 23}
	sortArray := quickSort(array, 0, len(array)-1)
	fmt.Println(sortArray)
	return sortArray

}
func quickSort(arr []int, leftBorder, rightBorder int) []int {
	if leftBorder < rightBorder {
		var p int
		// Берём опорную точку pivot, элементы меньше опорной точки оказываются слева,
		// элементы больше оказываются справая
		arr, p = partition(arr, leftBorder, rightBorder)
		// берём последовательность до pivot и после и повторяем процедуру
		arr = quickSort(arr, leftBorder, p-1)
		arr = quickSort(arr, p+1, rightBorder)
		// рекусрия продолжается до момента пока сортируем кусок не будет иметь длину 1
	}
	return arr

}

func partition(arr []int, leftBorder, rightBorder int) ([]int, int) {
	pivot := arr[rightBorder]
	i := leftBorder
	for j := leftBorder; j < rightBorder; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[rightBorder] = arr[rightBorder], arr[i]
	return arr, i
}
