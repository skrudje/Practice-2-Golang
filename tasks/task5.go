package tasks

import (
	"Golang/cryptomath"
	"fmt"
	"math/big"
)

func RunTask5() {
	fmt.Println("\n--- Task 5: Диофантово уравнение ---")
	
	var choice int
	fmt.Println("1. Вариант 1 (143x + 169y = 5)")
	fmt.Println("2. Ввести свои числа")
	fmt.Print("Выбор: ")
	fmt.Scan(&choice)

	var A, B, C *big.Int

	if choice == 1 {
		// 143a + 169b = 5
		A = big.NewInt(143)
		B = big.NewInt(169)
		C = big.NewInt(5)
	} else {
		A = readBigInt("Введите A: ")
		B = readBigInt("Введите B: ")
		C = readBigInt("Введите C: ")
	}

	res := cryptomath.ExtendedGCD(A, B)
	fmt.Printf("НОД(%s, %s) = %s\n", A, B, res.D)

	// проверка делимости C % GCD == 0
	rem := new(big.Int).Mod(C, res.D)
	
	if rem.Cmp(big.NewInt(0)) != 0 {
		fmt.Println("Решений в целых числах НЕТ (C не делится на НОД).")
	} else {
		// mult = C / GCD
		mult := new(big.Int).Div(C, res.D)
		
		// x = x0 * mult
		// y = y0 * mult
		finalX := new(big.Int).Mul(res.X, mult)
		finalY := new(big.Int).Mul(res.Y, mult)

		fmt.Printf("Частное решение:\n x = %s\n y = %s\n", finalX, finalY)
	}
}