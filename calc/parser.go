package calc

import (
	"errors"
	"regexp"
	"strconv"
)

// EvaluateExpression разбирает выражение и выполняет соответствующую операцию.
func EvaluateExpression(expression string) (string, error) {
	re := regexp.MustCompile(`^"([^"]{1,10})"\s*([\+\-\*/])\s*(?:"([^"]{1,10})"|(\d{1,2}))$`)
	matches := re.FindStringSubmatch(expression)

	if matches == nil {
		return "", errors.New("некорректный формат ввода")
	}

	str1 := matches[1]
	operator := matches[2]
	str2 := matches[3]
	numStr := matches[4]

	switch operator {
	case "+":
		if str2 == "" {
			return "", errors.New("второй операнд должен быть строкой для операции сложения")
		}
		return Add(str1, str2)

	case "-":
		if str2 == "" {
			return "", errors.New("второй операнд должен быть строкой для операции вычитания")
		}
		return Subtract(str1, str2)

	case "*":
		if numStr == "" {
			return "", errors.New("второй операнд должен быть числом для операции умножения")
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", errors.New("некорректное число")
		}
		return Multiply(str1, num)

	case "/":
		if numStr == "" {
			return "", errors.New("второй операнд должен быть числом для операции деления")
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", errors.New("некорректное число")
		}
		return Divide(str1, num)

	default:
		return "", errors.New("неподдерживаемая операция")
	}
}
