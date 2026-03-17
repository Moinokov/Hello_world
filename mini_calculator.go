package main

import "fmt"

func intToWord(n int) string {
	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 20:
		return "twenty"
	case 21:
		return "twenty-one"
	case 22:
		return "twenty-two"
	case 23:
		return "twenty-three"
	case 24:
		return "twenty-four"
	case 25:
		return "twenty-five"
	case 26:
		return "twenty-six"
	case 27:
		return "twenty-seven"
	case 28:
		return "twenty-eight"
	case 29:
		return "twenty-nine"
	case 30:
		return "thirty"
	default:
		return "big number"
	}
}

func myCalculation(a int, b int, op string) string {
	var result int
	switch op {
	case "+":
		result = a + b
	case "*":
		result = a * b
	default:
		panic("неподдерживаемая операция: " + op)
	}

	aWord := intToWord(a)
	bWord := intToWord(b)
	resWord := intToWord(result)

	return fmt.Sprintf("%s %s %s = %s", aWord, op, bWord, resWord)
}

func main() {
	const operand = 3

	for i := 1; i <= 30; i++ {
		fmt.Println(myCalculation(i, operand, "+"))
		fmt.Println(myCalculation(i, operand, "*"))
	}
}
