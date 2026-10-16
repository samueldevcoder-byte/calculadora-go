package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Scan(&input)

	var operator string
	operatorIndex := -1

	for i, char := range input {
		if i > 0 && (char == '+' || char == '-' || char == '*' || char == '/') {
			operator = string(char)
			operatorIndex = i
			break
		}
	}	
	if operator == "" {
		fmt.Println("Erro: operador inválido")
		return
	}

	parts := []string{
		input[:operatorIndex],
		input[operatorIndex+1:],
	}

	result, err := getResult(parts, operator)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("%s %s %s = %g\n", parts[0], operator, parts[1], result)
}

func getResult(parts []string, operator string) (float64, error) {
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return 0, fmt.Errorf("enter two numbers and an operator, e.g. 2*2")
	}

	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}
	num2, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}

	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Operator not valid")
	}
}
