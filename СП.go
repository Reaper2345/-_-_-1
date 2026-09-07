package main

import (
	"fmt"
	"math"
)

// Задание 1
func main() {
	fmt.Printf("Задание 1:\n")
	rent := 95000.0
	NewRent := rent * 1.10
	fmt.Printf("новая стоимость аренды: %.2f руб.\n\n", NewRent)

	// Задание 2
	fmt.Printf("Задание 2:\n")
	laptops := 55480 * 6
	monitors := 21830 * 3
	mice := 890 * 11
	keyboards := 1560 * 5
	sum := laptops + monitors + mice + keyboards
	fmt.Printf("Общая сумма покупки: %d руб.\n\n", sum)

	// Задание 3
	fmt.Printf("Задание 3:\n")
	a := 5000
	b := 256
	c := a / b
	d := a % b
	fmt.Printf("Можно разместить файлов: %d\n", c)
	fmt.Printf("останется свободного места: %d Гб\n\n", d)

	// Задание 4
	fmt.Printf("Задание 4:\n")
	var temp float64
	fmt.Print("Введите температуру в фаренгейтах: ")
	fmt.Scan(&temp)
	cel := 5.0 / 9.0 * (temp - 32)
	fmt.Printf("Температура в цельсиях: %.2f C\n\n", cel)

	// Задание 5
	fmt.Printf("Задание 5:\n")
	var r float64
	fmt.Print("Введите радиус клумбы: ")
	fmt.Scan(&r)
	L := 2 * math.Pi * r
	S := math.Pi * r * r
	fmt.Printf("Длина окружности: %.2f\n", L)
	fmt.Printf("Площадь круга: %.2f\n\n", S)

	// Задание 6
	fmt.Printf("Задание 6:\n")
	var i, a1, years float64
	fmt.Print("Введите начальную сумму вклада: ")
	fmt.Scan(&i)
	fmt.Print("Введите годовую процентную ставку: ")
	fmt.Scan(&a1)
	fmt.Print("Введите количество лет: ")
	fmt.Scan(&years)
	result := i * math.Pow(1+a1/100, years)
	fmt.Printf("Итоговая сумма вклада: %.2f\n\n", result)

	// Задание 9
	fmt.Printf("Задание 9:\n")
	var PurchaseSum float64
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&PurchaseSum)
	discount := PurchaseSum * 0.8
	fmt.Printf("Сумма со скидкой 20%%: %.2f\n\n", discount)

	// Задание 10
	fmt.Printf("Задание 10:\n")
	var A, B float64
	fmt.Print("Введите число a: ")
	fmt.Scan(&A)
	fmt.Print("Введите число b: ")
	fmt.Scan(&B)
	delenie := A / B
	roundedUp := math.Ceil((delenie))
	roundedDown := math.Floor(float64(delenie))
	fmt.Printf("Результат деления: %.4f\n", delenie)
	fmt.Printf("Округление вверх: %.0f\n", roundedUp)
	fmt.Printf("Округление вверх: %.0f\n", roundedDown)
}
