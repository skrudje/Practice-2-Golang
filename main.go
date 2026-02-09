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
			// очистка ввода если ошибка
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			tasks.RunTask1()
		case 2:
			tasks.RunTask2()
		case 3:
			tasks.RunTask3()
		case 4:
			tasks.RunTask4()
		case 5:
			tasks.RunTask5()
		case 6:
			tasks.RunTask6()
		case 0:
			return
		default:
			fmt.Println("Неверный ввод.")
		}
	}
}