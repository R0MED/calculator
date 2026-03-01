package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	var action string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	fmt.Print("Введите действие (плюс +, минус -, делить /, умножить *, остато от деления %, степень ^): ")
	fmt.Scan(&action)

}
