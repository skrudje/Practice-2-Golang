package tasks

import (
	"Golang/cryptomath"
	"fmt"
)

func RunTask2() {
	fmt.Println("\n--- Task 2: Алгоритм Евклида (ax + by = GCD) ---")
	a := readBigInt("Введите a: ")
	b := readBigInt("Введите b: ")

	res := cryptomath.ExtendedGCD(a, b)

	fmt.Printf("НОД(%s, %s) = %s\n", a, b, res.D)
	fmt.Printf("Коэффициенты: x = %s, y = %s\n", res.X, res.Y)
}