package tasks

import (
	"Golang/cryptomath"
	"fmt"
)

func Attack() {
	fmt.Println("\n    Task 6: Атака Man-In-The-Middle (MITM)    ")
	
	p := readBigInt("Введите общее простое число p: ")
	g := readBigInt("Введите основание g: ")

	aPriv := readBigInt("\n[Alice] Секретный ключ: ")
	aPub := cryptomath.PowerMod(g, aPriv, p)
	fmt.Printf("[Alice] Отправляет A = %s\n", aPub)

	fmt.Println("\n--- [Eva] ПЕРЕХВАТ A ---")
	ePriv := readBigInt("[Eva] Секретный ключ Евы: ")
	ePub := cryptomath.PowerMod(g, ePriv, p)
	fmt.Printf("[Eva] Отправляет Бобу поддельный E = %s\n", ePub)

	bPriv := readBigInt("\n[Bob] Секретный ключ: ")
	bPub := cryptomath.PowerMod(g, bPriv, p)
	
	secretBob := cryptomath.PowerMod(ePub, bPriv, p)
	fmt.Printf("[Bob] Отправляет B = %s\n", bPub)
	fmt.Printf("[Bob] Вычислил секрет: %s\n", secretBob)

	fmt.Println("\n--- [Eva] ПЕРЕХВАТ B ---")
	fmt.Printf("[Eva] Отправляет Алисе поддельный E = %s\n", ePub)

	secretEveAlice := cryptomath.PowerMod(aPub, ePriv, p)
	secretEveBob := cryptomath.PowerMod(bPub, ePriv, p)

	secretAlice := cryptomath.PowerMod(ePub, aPriv, p)
	fmt.Printf("\n[Alice] Вычислила секрет: %s\n", secretAlice)

	fmt.Println("\n--- РЕЗУЛЬТАТ АТАКИ ---")
	fmt.Printf("Канал Алиса-Ева: %s <==> %s\n", secretAlice, secretEveAlice)
	fmt.Printf("Канал Боб-Ева:   %s <==> %s\n", secretBob, secretEveBob)

	if secretAlice.Cmp(secretEveAlice) == 0 && secretBob.Cmp(secretEveBob) == 0 {
		fmt.Println("\n[УСПЕХ]\n\n")
	}
}