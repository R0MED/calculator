package main

import (
	"fmt"
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
