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

	for _, data := range Inter {
		fmt.Println("С помощью dataTypeQualifier получен тип данных: \n", dataTypeQualifier(data))
	}

	for _, data := range Inter {
		fmt.Printf("С помощью Printf получен тип данных: %T \n", data)
	}

	for _, data := range Inter {
		fmt.Println("С помощью reflect.TypeOf получен тип данных: \n", reflect.TypeOf(data))
	}

	for _, data := range Inter {
		fmt.Println("С помощью reflect.ValueOf.Kind получен тип данных: \n", reflect.ValueOf(data).Kind())
	}
}

func dataTypeQualifier(inter interface{}) string {
	switch inter.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown type"
	}
}
