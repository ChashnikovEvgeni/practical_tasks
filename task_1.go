package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//	Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

// Структура Human с полями Id, Name, Age
// и методам Birhday() и Print_name()
type Human struct {
	Id   int
	Name string
	Age  int
}

// Увеличение возраста человека
func (h *Human) Birthday() {
	h.Age = h.Age + 1
}

// Печать имени человека
func (h Human) Print_name() {
	fmt.Println(h.Name)

}

// Структура машина с полями Brand, Mileage
// и методом Traveling()
type Car struct {
	Brand   string
	Mileage float32
}

// Метод увеличивает пробег машина на указанное значение
func (c *Car) Traveling(km float32) {
	c.Mileage = c.Mileage + km
}

// Структура Action, в которую встраиваются структуры Human и Car
type Action struct {
	Id int
	Human
	Car
}

// Функция для демонстрации работы встраивания
func Task_1_solution() {
	fmt.Println("---------------Решение задачи 1---------------")
	action := Action{
		Id:    1,
		Human: Human{Id: 2, Name: "Влад", Age: 25},
		Car:   Car{Brand: "BMW", Mileage: 500.65},
	}
	//Пример shorthands, данный синтаксис можно использовать вместо action.Human.Name и action.Human.Age и т.п. , если бы поле Name
	// присутствовало в структуре Car, которая тоже встраивается в структуру Action, то использование такого синтаксиска было бы невозможным. Это был бы пример colliding
	fmt.Printf("Id: %d , Имя: %s, Возраст: %v, Машина марки: %s, Пробег машины: %g ", action.Id, action.Name, action.Age, action.Brand, action.Mileage)
	// output : Id: 1 , Имя: Влад, Возраст: 25, Машина марки: BMW, Пробег машины: 500.65
	// значение поля Id взято именно из струкруты Action, т.к. произошёл Shadowing, и взято было первое встреченное поле с указанным названием поля.
	//Для получения поля Id встроенной структуры Human нужно использоваться следующийсинтаксис: action.Human.Id
	fmt.Printf("Human Id : %d, Action Id: %d", action.Human.Id, action.Id)
	//Вызов методов Print_name() и Birthday, полученного входе встраивания структуры Human в структуру Action
	action.Print_name()
	action.Birthday()
	//Вызов методa Traveling(), полученного входе встраивания структуры Car в структуру Action
	action.Traveling(1000.23)
	//Вывод значение полей структуры Action после изменений произведённых методами Birthday(), Traveling()
	fmt.Printf("Id: %d , Имя: %s, Возраст: %v, Машина марки: %s, Пробег машины: %g ", action.Id, action.Name, action.Age, action.Brand, action.Mileage)
	fmt.Println("---------------Конец решения задачи 1---------------")

}
