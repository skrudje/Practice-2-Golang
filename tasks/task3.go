package tasks

import (
	"Golang/cryptomath"
	"fmt"
)

func ReverseElement() {
	fmt.Println("\n    Task 3: Обратный элемент (C^-1 mod m)    ")
	c := readBigInt("Введите C: ")
	m := readBigInt("Введите m: ")

	inv := cryptomath.ModInverse(c, m)

	if inv == nil {
		fmt.Println("Обратного элемента не существует (НОД != 1).")
	} else {
		fmt.Printf("Обратный элемент: %s\n", inv)
	}
}