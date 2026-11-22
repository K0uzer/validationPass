package main

// // Напишите программу, которая проверяет корректность пароля.

// // Следуйте указаниям по тексту программы.
// // Чтобы протестировать своё решение, запустите go test -v в папке с упражнением.
// // Не меняйте сигнатуры функций и определение типа validator, чтобы тесты прошли корректно.

import (
	"fmt"
	"slices"
	// "strings"
	// "unicode"
	// "unicode/utf8"
)

// // validator проверяет строку на соответствие некоторому условию
// // и возвращает результат проверки
// type validator func(s string) bool

// // digits возвращает true, если s содержит хотя бы одну цифру
// // согласно unicode.IsDigit(), иначе false
// func digits(s string) bool {
// 	for _, v := range s {
// 		if ok := unicode.IsDigit(v); ok {
// 			return true
// 		}
// 	}
// 	return false
// }

// // letters возвращает true, если s содержит хотя бы одну букву
// // согласно unicode.IsLetter(), иначе false
// func letters(s string) bool {
// 	for _, v := range s {
// 		if ok := unicode.IsLetter(v); ok {
// 			return true
// 		}
// 	}
//     return false
// }

// // minlen возвращает валидатор, который проверяет, что длина
// // строки согласно utf8.RuneCountInString() не меньше указанной
// func minlen(length int) validator {
// 	return func(s string) bool {
// 		if utf8.RuneCountInString(s) > length {
// 			return true
// 		}
// 		return false
// 	}
// }

// // and возвращает валидатор, который проверяет, что все
// // переданные ему валидаторы вернули true
// func and(funcs ...validator) validator {
// 	return func(s string) bool {
// 		for _, el := range funcs {
// 			if !el(s) {
// 				return false
// 			}
// 		}
// 		return  true
// 	}
// }

// // or возвращает валидатор, который проверяет, что хотя бы один
// // переданный ему валидатор вернул true
// func or(funcs ...validator) validator {
// 	return func(s string) bool {
// 		for _, el := range funcs {
// 			if el(s) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// }



type User struct {
	Name string
}

type Add func(x, y int) int

func (a Add) Summ(x, y int) int {
	fmt.Println("Привет")
	return a(x, y)
}

func add(x, y int) int {
	return x + y
}

func main() {

	var a Add
	a = add

	fmt.Println(a(1, 5))
	fmt.Println(a.Summ(1, 5))

    s := []int{10, 20, 30, 40, 50}

    s = slices.Delete(s, 1, 3)

    fmt.Println(s) // [10 40 50]

	

	// a := []int{}
	// b := a[:2]
	// c := a[2:]
	// b = append(b, 100)
	// c = append(c, 200, 201, 202, 203)
	// fmt.Println(a) // 1 2 100 4 5
	// fmt.Println(b) // 1 2 100
	// fmt.Println(c) // 100 4 5 200 201 202 203
}