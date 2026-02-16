package tasks

import (
	"Golang/cryptomath"
	"fmt"
)

func Hellman() {
	fmt.Println("\n    Task 4: Диффи-Хеллман    ")
	
	p := readBigInt("Введите общее простое число p: ")
	g := readBigInt("Введите основание g: ")

	aPriv := readBigInt("[Alice] Секретный ключ: ")
	bPriv := readBigInt("[Bob] Секретный ключ: ")

	aPub := cryptomath.PowerMod(g, aPriv, p)
	bPub := cryptomath.PowerMod(g, bPriv, p)

	fmt.Printf("Публичный ключ Alice: %s\n", aPub)
	fmt.Printf("Публичный ключ Bob:   %s\n", bPub)

	secA := cryptomath.PowerMod(bPub, aPriv, p)
	secB := cryptomath.PowerMod(aPub, bPriv, p)

	fmt.Printf("Секрет Alice: %s\n", secA)
	fmt.Printf("Секрет Bob:   %s\n", secB)

	if secA.Cmp(secB) == 0 {
		fmt.Println("Ключи совпали. Канал защищен.")
	} else {
		fmt.Println("Ошибка. Ключи разные.")
	}
}