package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var action string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка ввода первого числа")
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка ввода второго числа")
		return
	}

	fmt.Print("Введите действие (+ - * / % ^): ")
	_, err = fmt.Scan(&action)
	if err != nil {
		fmt.Println("Ошибка ввода действия")
		return
	}

	var result float64
	var valid bool = true

	switch action {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на 0")
			valid = false
		} else {
			result = num1 / num2
		}
	case "%":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на 0 (остаток)")
			valid = false
		} else {
			result = float64(int64(num1) % int64(num2))
		}
	case "^":
		result = math.Pow(num1, num2)
	default:
		fmt.Printf("Ошибка: неизвестная операция '%s'\n", action)
		valid = false
	}

	if valid {
		fmt.Printf("%.4f %s %.4f = %.4f\n", num1, action, num2, result)
	}
}
