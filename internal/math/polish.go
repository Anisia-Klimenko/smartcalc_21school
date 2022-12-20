package math

import (
	"fmt"
	"strconv"
	"strings"
)

var operators = []string{
	"+", "-", "*", "/", "^", "mod", "sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log", "e",
}

var functions = []string{
	"sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log",
}

func Calculate(input string, xVal float64) (result string) {
	output := infixToPostfix(input)
	fmt.Println(output)
	res, err := evaluatePostfix(output, xVal)
	if err != nil {
		result = "error"
	} else {
		result = strconv.FormatFloat(res, 'f', -1, 64)
		ind := strings.Index(result, ".")
		if ind > -1 && len(result[ind:]) > 7 {
			result = result[:ind+8]
		}
	}
	return
}

// Function to return precedence of operators
func precedence(s string) int {
	if s == "^" || isFunction(s) || s == "e" {
		return 3
	} else if s == "/" || s == "*" || s == "mod" {
		return 2
	} else if s == "++" || s == "--" || s == "+" || s == "-" {
		return 1
	} else {
		return -1
	}
}

func isFunction(input string) bool {
	for _, f := range functions {
		if input == f {
			return true
		}
	}
	return false
}

func isDigit(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return false
	}
	return true
}

// infixToPostfix convert infix notation to postfix
func infixToPostfix(infix string) (postfix Stack) {
	inputLex := splitLex(infix)
	var sta Stack
	for _, s := range inputLex {
		if isDigit(s) {
			postfix.Push(s)
		} else if s == "x" {
			postfix.Push("x")
		} else if s == "(" || isFunction(s) {
			sta.Push(s)
		} else if s == ")" {
			for sta.Top() != "(" {
				if sta.IsEmpty() {
					return Stack{"error"}
				}
				postfix.Push(sta.Top())
				sta.Pop()
			}
			sta.Pop()
			if isFunction(sta.Top()) {
				postfix.Push(sta.Top())
				sta.Pop()
			}
		} else {
			for !sta.IsEmpty() && precedence(s) <= precedence(sta.Top()) {
				postfix.Push(sta.Top())
				sta.Pop()
			}
			sta.Push(s)
		}
	}
	for !sta.IsEmpty() {
		if sta.Top() == "(" || sta.Top() == ")" {
			return Stack{"error"}
		}
		postfix.Push(sta.Top())
		sta.Pop()
	}
	return
}

func splitLex(input string) (res Stack) {
	var digit string
	var ind = 0
	for range input {
		if ind >= len(input) {
			break
		}
		if input[ind] == '(' || input[ind] == ')' || input[ind] == 'x' || input[ind] == 'e' {
			res.Push(string(input[ind]))
			ind++
		} else {
			_, err := strconv.Atoi(string(input[ind]))
			for err == nil || input[ind] == '.' {
				digit += string(input[ind])
				ind++
				if len(input[ind:]) > 0 {
					_, err = strconv.Atoi(string(input[ind]))
				} else {
					break
				}
			}
			if len(digit) > 0 {
				res.Push(digit)
				digit = ""
			}
			if ind < len(input) {
				for _, op := range operators {
					if strings.HasPrefix(input[ind:], op) {
						res.Push(op)
						ind += len(op)
						break
					}
				}
			}
		}
	}
	return
}
