package main

import "fmt"

//Удалить i-ый элемент из слайса

func Task_23_solution() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Исходный слайс: %#v\n", slice)
	Appenddelete(slice, 3)

}

// Удаление с помощью переписывание исходного слайса
// берётся кусок слайса до элемента и после и первой части дописывается вторая
// элементы второй части сдвинаются на 1 индекс влево
func Appenddelete(slice []int, i int) {
	appendSlice := append(slice[:i], slice[i+1:]...)
	fmt.Printf("Appenddelete, итоговый слайс: %#v\n", appendSlice)
}
