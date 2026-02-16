package main

import (
	"Golang/tasks"
	"fmt"
)

func main() {
	for {
		fmt.Println("1. Ферма (mod pow)")
		fmt.Println("2. Евклид (ax+by)")
		fmt.Println("3. Обратный элемент")
		fmt.Println("4. Диффи-Хеллман")
		fmt.Println("5. Диофантово уравнение")
		fmt.Println("6. Эмуляция атаки на базе программы 4 таска")
		fmt.Println("0. Выход")
		fmt.Print("> ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1: tasks.Ferma()
		case 2: tasks.Evklid()
		case 3: tasks.ReverseElement()
		case 4: tasks.Hellman()
		case 5: tasks.Diofantovo()
		case 6: tasks.Attack()
		case 0: return
		default: fmt.Println("Неверный ввод.")
		}
	}
}