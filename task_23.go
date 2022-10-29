package main

import "fmt"

//Удалить i-ый элемент из слайса

func Task_23_solution() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Исходный слайс: %#v\n", slice)
	Appenddelete(slice, 3)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Copydelete(slice, 4)

}

// ломаем слайс
func Appenddelete(slice []int, i int) {
	appendSlice := append(slice[:i], slice[i+1:]...)
	fmt.Printf("Appenddelete, итоговый слайс: %#v\n", appendSlice)
}

// работает криво
func Copydelete(slice []int, i int) {
	copySlice := slice[:i+copy(slice[i:], slice[i+1:])]
	fmt.Printf("Copydelete, итоговый слайс: %#v\n", copySlice)
}
