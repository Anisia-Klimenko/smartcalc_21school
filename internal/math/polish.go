package math

import (
	"strconv"
	"strings"
)

// all operators and functions
var opfuncs = []string{
	"+", "-", "*", "/", "^", "mod", "e", "sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log",
}

// all operators
var operators = []string{
	"+", "-", "*", "/", "^", "mod", "e",
}

// all functions
var functions = []string{
	"sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log",
}

// Calculate calculates result of input equation with passing x value xVal
func Calculate(input string, xVal float64) (result string) {
	// Convert infix notation to postfix
	output := infixToPostfix(input)
	// Checks result of conversion (left for evaluation)
	//fmt.Println(output)

	// Evaluate expression
	res, err := evaluatePostfix(output, xVal)
	if err != nil {
		result = "error"
	} else {
		// Format result
		result = strconv.FormatFloat(res, 'f', -1, 64)
		ind := strings.Index(result, ".")
		if ind > -1 && len(result[ind:]) > 7 {
			result = result[:ind+8]
		}
	}
	return
}

// precedence returns precedence of operators, 3 means that this operation calculates
// earlier than 1, returns -1 in case of error
func precedence(s string) int {
	if s == "^" || isFunction(s) || s == "e" {
		return 3
	} else if s == "/" || s == "*" || s == "mod" {
		return 2
	} else if s == "+" || s == "-" {
		return 1
	} else {
		return -1
	}
}

// isOperator checks if input is a function
func isOperator(input string) bool {
	for _, f := range operators {
		if input == f {
			return true
		}
	}
	return false
}

// isFunction checks if input is a function
func isFunction(input string) bool {
	for _, f := range functions {
		if input == f {
			return true
		}
	}
	return false
}

// isDigit checks if input is a digit
func isDigit(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil || len(input) == 0 {
		return false
	}
	return true
}

// infixToPostfix convert infix notation to postfix using Dijkstra's algorithm
func infixToPostfix(infix string) (postfix Stack) {
	// Split input into lexemes
	inputLex := splitLex(infix)
	// sta is a buffer stack to save functions, operators, braces
	var sta Stack
	for _, s := range inputLex {
		if isDigit(s) {
			postfix.Push(s)
		} else if s == "x" || s == "-x" {
			postfix.Push(s)
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

// splitLex splits input string into lexemes and returns result as Stack
func splitLex(input string) (res Stack) {
	var digit string
	var ind = 0
	var sign = 1
	for range input {
		if ind >= len(input) {
			break
		}
		if input[ind] == '(' || input[ind] == ')' || input[ind] == 'x' || input[ind] == 'e' {
			// "(", ")", "x", "e" - separate symbols
			res.Push(string(input[ind]))
			ind++
		} else if len(input[ind:]) > 1 && input[ind:ind+2] == "-x" {
			res.Push("-x")
			ind += 2
		} else {
			// Check if digit
			_, err := strconv.Atoi(string(input[ind]))
			for err == nil || input[ind] == '.' {
				// Write whole number
				digit += string(input[ind])
				ind++
				if len(input[ind:]) > 0 {
					_, err = strconv.Atoi(string(input[ind]))
				} else {
					break
				}
			}
			// If input is digit
			if len(digit) > 0 {
				// If unary minus
				if sign < 0 {
					digit = "-" + digit
					sign = 1
				}
				res.Push(digit)
				digit = ""
			} else if sign < 0 {
				// There is no number before minus, so this operator is not unary
				res.Push("-")
				sign = 1
			}
			if ind < len(input) {
				// Check if input is operator or function
				for _, op := range opfuncs {
					if strings.HasPrefix(input[ind:], op) {
						// Operator is unary, if:
						// 1. it is the first symbol
						// 2. there is one more operator before
						// 3. there is "(" before
						if op == "-" && (res.IsEmpty() || isOperator(res.Top()) || res.Top() == "(") {
							// Change sign of digit if there is unary minus
							sign = -1
						} else if op == "+" && (res.IsEmpty() || isOperator(res.Top()) || res.Top() == "(") {
							// Do nothing if there is unary plus
						} else {
							// Add operator to result otherwise
							res.Push(op)
						}
						ind += len(op)
						break
					}
				}
			}
		}
	}
	return
}
