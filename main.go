package main

import (
	"bufio"
	"fmt"
	"os"
	"string-calculator/calc"
)

func main() {
	fmt.Println("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		os.Exit(1)
	}

	// Убираем лишние пробелы и символы новой строки
	expression = expression[:len(expression)-1]

	// Разбираем выражение и выполняем операцию
	result, err := calc.EvaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	// Обрезаем результат, если он длиннее 40 символов
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println("Результат:", result)
}
