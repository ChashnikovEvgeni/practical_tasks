package main

import (
	"fmt"
	"reflect"
)

/*Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{} */

func Task_14_solution() {
	type km int64
	var way km
	way = 56
	Inter := []interface{}{52, "hello world!", true, make(chan int), way}

	//определение типа с помощью самописной функции
	for _, data := range Inter {
		fmt.Println("С помощью dataTypeQualifier получен тип данных: \n", dataTypeQualifier(data))
	}

	// определние типа с помщью  fmt.Printf и verbs %T
	for _, data := range Inter {
		fmt.Printf("С помощью Printf получен тип данных: %T \n", data)
	}

	// определние типа с помщью  reflect.TypeOf()
	for _, data := range Inter {
		fmt.Println("С помощью reflect.TypeOf получен тип данных: \n", reflect.TypeOf(data))
	}

	// определение с помощью reflect.ValueOf(data).Kind()
	for _, data := range Inter {
		fmt.Println("С помощью reflect.ValueOf.Kind получен тип данных: \n", reflect.ValueOf(data).Kind())
	}
}

// своя функция для определения типа
func dataTypeQualifier(inter interface{}) string {
	switch inter.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float32:
		return "int"
	case float64:
		return "int"
	default:
		return "unknown type"
	}
}
