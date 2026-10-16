package main

import "testing"

func TestGetResult(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		operator string
		expected float64
	}{
		{"soma", "2", "3", "+", 5},
		{"subtração", "10", "4", "-", 6},
		{"multiplicação", "3", "4", "*", 12},
		{"divisão", "10", "2", "/", 5},
		{"números negativos", "-5", "3", "*", -15},
		{"números decimais", "2.2", "4", "*", 8.8},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := getResult(
				[]string{test.num1, test.num2},
				test.operator,
			)

			if err != nil {
				t.Fatalf("não esperávamos um erro: %v", err)
			}

			if result != test.expected {
				t.Errorf(
					"esperávamos %g, mas recebemos %g",
					test.expected,
					result,
				)
			}
		})
	}
}

func TestGetResultErrors(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		operator string
	}{
		{"divisão por zero", "10", "0", "/"},
		{"número inválido", "abc", "2", "+"},
		{"segundo número vazio", "2", "", "*"},
		{"primeiro número vazio", "", "2", "+"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := getResult(
				[]string{test.num1, test.num2},
				test.operator,
			)

			if err == nil {
				t.Errorf("esperávamos um erro, mas não recebemos nenhum")
			}
		})
	}
}
