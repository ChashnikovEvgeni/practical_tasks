package main

import (
	"fmt"
	"unicode/utf8"
)

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
} */

func Task_15_solution() {

	someFunc()
	RightGet()

}

var justString string

func createHugeString(n int) string {
	return "♣Борис四五六DSSDgfdhfghdfgjfhkghjlghjlhjk;h;hjkl;jkl;jkl;jkl;jktryj;klrtfjydfgkhjdfgkl;hdfgl;jfghjfghjfghkfhjkhjkuyiryuiyuityukjghkhj"
}

func someFunc() {
	/*Первое негативное последствие подобной реализаци - невозможность полученния данных из этих переменных
	за пределами функции. Если подразумевается, что v и justString будет использоваться только в данной функции,
	то всё в порядке, иначе их нужно возращаться return.
	*/
	v := createHugeString(1 << 10)
	fmt.Println(utf8.RuneCountInString(v), "runes")
	fmt.Println(len(v))
	/* Если человек, написавший этот код, расчитывает, что используя v[:100] он получит первые
	100 символов v, то это не так. Тип string является срезом из байт. Значит, применяя подобный синтаксис,
	человек ВОЗМОЖНО получит не первые 100 символов, а первые 100 байт. Да, если использовать стандартный набор
	символов в строке он и правда получит 100 символов, ведь они кодируются 1 байтом.
	Но, если в строке будут спец. символы, допустим ♣, то он будет кодироваться большим количеством байт(2 байта)
	Аналогичная проблема с символами других языков(привёден пример русского(Борис) и японского(四五六) в строке выше)
	*/
	justString = v[:100]
	fmt.Println("Количество символов в полученной строке:")

	fmt.Println(utf8.RuneCountInString(justString), "runes")
	/*
			output:
		Количество символов в полученной строке:
		87 runes

	*/
	for i, w := 0, 0; i < len(justString); i += w {
		runeValue, width := utf8.DecodeRuneInString(justString[i:])
		fmt.Printf("%#U Начинается с позиции байта %d\n", runeValue, i)
		w = width
	}

	/*  Часть output:
	U+2663 '♣' Начинается с позиции байта 0
	U+0411 'Б' Начинается с позиции байта 3
	U+043E 'о' Начинается с позиции байта 5
	U+0440 'р' Начинается с позиции байта 7
	U+0438 'и' Начинается с позиции байта 9
	U+0441 'с' Начинается с позиции байта 11
	U+56DB '四' Начинается с позиции байта 13
	U+4E94 '五' Начинается с позиции байта 16
	U+516D '六' Начинается с позиции байта 19
	U+0044 'D' Начинается с позиции байта 22
	U+0053 'S' Начинается с позиции байта 23
	U+0053 'S' Начинается с позиции байта 24
	*/

}

// Правильная реализация
func RightGet() string {
	v := createHugeString(1 << 10)
	fmt.Println(utf8.RuneCountInString(v), "runes")

	// Если нужно получить первые несколько символов, то лучше поступить так
	RightRune := make([]rune, 100)
	i := 0
	for _, r := range v {
		if i >= cap(RightRune) {
			break
		}
		RightRune[i] = r
		i++
	}

	RightString := string(RightRune)
	fmt.Println(utf8.RuneCountInString(RightString), "Рун в  RightString")
	return RightString
}
