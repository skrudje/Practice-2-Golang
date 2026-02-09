package tasks

import (
	"Golang/cryptomath"
	"fmt"
	"math/big"
)

// вспомогательная функция для ввода bigInt
func readBigInt(label string) *big.Int {
	var s string
	fmt.Print(label)
	fmt.Scan(&s)
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}

func RunTask1() {
	fmt.Println("\n--- Task 1: Теорема Ферма ---")
	a := readBigInt("Введите a: ")
	x := readBigInt("Введите степень x: ")
	p := readBigInt("Введите модуль p: ")

	if !cryptomath.IsPrime(p) {
		fmt.Println("Ошибка: Модуль p не является простым")
		return
	}

	// проверка теоремы ферма: a^(p-1) == 1
	pMinus1 := new(big.Int).Sub(p, big.NewInt(1))
	check := cryptomath.PowerMod(a, pMinus1, p)
	
	if check.Cmp(big.NewInt(1)) == 0 {
		fmt.Printf("Теорема выполняется: %s^(%s) mod %s = 1\n", a, pMinus1, p)
	} else {
		fmt.Println("Теорема не выполняется (или a кратно p).")
	}

	result := cryptomath.PowerMod(a, x, p)
	fmt.Printf("Результат: %s\n", result)
}