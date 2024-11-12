package calc

import (
	"errors"
	"strings"
)

// Add складывает две строки.
func Add(str1, str2 string) (string, error) {
	return str1 + str2, nil
}

// Subtract вычитает одну строку из другой.
func Subtract(str1, str2 string) (string, error) {
	return strings.ReplaceAll(str1, str2, ""), nil
}

// Multiply умножает строку на число.
func Multiply(str string, num int) (string, error) {
	if num < 1 || num > 10 {
		return "", errors.New("число должно быть от 1 до 10 включительно")
	}
	return strings.Repeat(str, num), nil
}

// Divide делит строку на число.
func Divide(str string, num int) (string, error) {
	if num < 1 || num > 10 {
		return "", errors.New("число должно быть от 1 до 10 включительно")
	}
	length := len(str) / num
	if length == 0 {
		return "", nil
	}
	return str[:length], nil
}
