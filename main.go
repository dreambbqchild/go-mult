package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getNumber(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var result strings.Builder
	for _, c := range text {
		if unicode.IsDigit(c) {
			result.WriteRune(c)
		}
	}

	return result.String()
}

func getDigitAt(str string, i int) int {
	length := len(str)
	if i >= length {
		return 0
	}

	return int(str[length-i-1] - '0')
}

func multiply(multiplicand string, multiplier string) {
	if multiplicand == "0" || multiplier == "0" {
		fmt.Print(0)
		return
	}

	answer := []int{}
	carry := 0
	cols := len(multiplicand) + len(multiplier)

	for col := 0; col < cols; col++ {
		value := carry
		for i := 0; i <= col; i++ {
			multiplicandDigit := getDigitAt(multiplicand, col-i)
			multiplierDigit := getDigitAt(multiplier, i)
			value += multiplicandDigit * multiplierDigit
		}
		answer = append(answer, value%10)
		carry = value / 10
	}

	leastSignificantDigitIndex := len(answer) - 1
	for i := leastSignificantDigitIndex; i >= 0; i-- {
		if i == leastSignificantDigitIndex && answer[i] == 0 {
			continue
		}

		fmt.Print(answer[i])
	}

	fmt.Println()
}

func main() {
	multiplicand := getNumber("Multiplicand:")
	multiplier := getNumber("Multiplier:")

	multiply(multiplicand, multiplier)
}
