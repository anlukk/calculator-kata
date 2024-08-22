package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Print("введите выражение:")
	fmt.Scanln(&input)

	result, err := calculator(input)
	if err != nil {
		panic(err)
	}

	if result == "" {
		return
	}

	fmt.Println("ваш результат равен:", result)

}

func calculator(input string) (string, error) {
	input = strings.ReplaceAll(input, " ", "")

	var (
		a, b    int
		operator string
		isRoman bool
		isMixed bool
	)

	if strings.Contains(input, "+") {
		operator = "+"
		parts := strings.Split(input, "+")

		if len(parts) != 2 {
			return "", errors.New("некорректный формат выражения")
		}

		a, b, isRoman, isMixed = parseNumber(
			parts[0]), 
			parseNumber(parts[1]), 
			isRomanNumber(parts[0], parts[1]), 
			isMixedNumber(parts[0], parts[1])
	} else if strings.Contains(input, "-") {
		operator = "-"
		parts := strings.Split(input, "-")

		if len(parts) != 2 {
			return "", errors.New("некорректный формат выражения")
		}

		a, b, isRoman, isMixed = parseNumber(
			parts[0]), 
			parseNumber(parts[1]), 
			isRomanNumber(parts[0], parts[1]), 
			isMixedNumber(parts[0], parts[1])
	} else if strings.Contains(input, "*") {
		operator = "*"
		parts := strings.Split(input, "*")

		if len(parts) != 2 {
			return "", errors.New("некорректный формат выражения")
		}

		a, b, isRoman, isMixed = parseNumber(
			parts[0]), 
			parseNumber(parts[1]), 
			isRomanNumber(parts[0], parts[1]), 
			isMixedNumber(parts[0], parts[1])
	} else if strings.Contains(input, "/") {
		operator = "/"
		parts := strings.Split(input, "/")

		if len(parts) != 2 {
			return "", errors.New("некорректный формат выражения")
		}

		a, b, isRoman, isMixed = parseNumber(
			parts[0]), 
			parseNumber(parts[1]), 
			isRomanNumber(parts[0], parts[1]), 
			isMixedNumber(parts[0], parts[1])
	} else {
		return "", errors.New("некорректный оператор")
	}

	if !isRoman && (a > 10 || b > 10) {
		return "", errors.New("числа вне допустимого диапазона для арабских чисел")
	}

	if isMixed {
		return "", errors.New("различные системы счисления в выражении")
	}

	var result int

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("деление на ноль")
		}
		result = a / b
	}

	if isRoman {
		if result < 1 {
			return "", errors.New("результат меньше единицы для римских чисел")
		}
		return arabicToRoman(result), nil
	}

	return strconv.Itoa(result), nil
}

var romanToArabic = map[string]int{
	"I":  1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V":  5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX":  9,
	"X":   10,
}

func parseNumber(s string) int {
	if val, ok := romanToArabic[s]; ok {
		return val
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		panic("некорректный формат числа")
	}

	if val < 1 || val > 10 {
		panic("число вне допустимого диапазона")
	}

	return val
}

func isRomanNumber(a, b string) bool {
	_, aIsRoman := romanToArabic[a]
	_, bIsRoman := romanToArabic[b]

	return aIsRoman && bIsRoman
}

func isMixedNumber(a, b string) bool {
	_, aIsRoman := romanToArabic[a]
	_, bIsRoman := romanToArabic[b]

	return (aIsRoman && !bIsRoman) || (!aIsRoman && bIsRoman)
}

func arabicToRoman(num int) string {
	var roman strings.Builder
	romanMap := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
		{1, "I"},
	}

	for _, rm := range romanMap {
		for num >= rm.Value {
			roman.WriteString(rm.Symbol)
			num -= rm.Value
		}
	}

	return roman.String()
}


