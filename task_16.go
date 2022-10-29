package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func Task_16_solution() []int {

	array := []int{22, 56, 23, 76, -23, 45, 1, 12, 53, 86, -64, -1, 76, 23}
	sortArray := quickSort(array, 0, len(array)-1)
	fmt.Println(sortArray)
	return sortArray

}
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr

}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
