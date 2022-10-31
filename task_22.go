package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

// суммирование
func sum(a, b *big.Int) *big.Int {
	var c big.Int
	c.Add(a, b)
	return &c
}

// вычитание
func sub(a, b *big.Int) *big.Int {
	var c big.Int
	c.Sub(a, b)
	return &c
}

// Деление получение целой части
func div(a, b *big.Int) *big.Int {
	var c big.Int
	c.Div(a, b)
	return &c
}

// Деление получение остатка
func mod(a, b *big.Int) *big.Int {
	var c big.Int
	c.Mod(a, b)
	return &c
}

// Деление получение целой части и остатка
func divMod(a, b *big.Int) (*big.Int, *big.Int) {
	var c big.Int
	m := new(big.Int)
	c.DivMod(a, b, m)
	return &c, m
}

// Умножение
func mul(a, b *big.Int) *big.Int {
	var c big.Int
	c.Mul(a, b)
	return &c
}

func Task_22_solution() {

	a := big.NewInt(11)
	b := big.NewInt(5)
	fmt.Printf("Сумма: %v\n", sum(a, b))
	fmt.Printf("Вычитание: %v\n", sub(a, b))
	fmt.Printf("Деление целое: %v\n", div(a, b))
	fmt.Printf("Деление остаток: %v\n", mod(a, b))
	d, m := divMod(a, b)
	fmt.Printf("Деление: %v\n и остаток:%v\n ", d, m)
	fmt.Printf("Умножение: %v\n", mul(a, b))

}
