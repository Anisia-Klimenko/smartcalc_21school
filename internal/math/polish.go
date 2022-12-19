package math

import (
	"fmt"
	"strconv"
	"strings"
)

var operators = []string{
	"++", "--", "+", "-", "*", "/", "^", "mod", "sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log",
}

var functions = []string{
	"sqrt", "sin", "asin", "cos", "acos", "tan", "atan", "ln", "log",
}

func Calculate(input string) (result string) {
	output := infixToPostfix(input)
	fmt.Println(output)
	for _, s := range output {
		result += s
	}
	return
}

// Function to return precedence of operators
func precedence(s string) int {
	if (s == "^") || (s == "sqrt") || (s == "sin") || (s == "asin") || (s == "cos") || (s == "acos") || (s == "tan") || (s == "atan") || (s == "ln") || (s == "log") {
		return 3
	} else if (s == "/") || (s == "*") || (s == "mod") {
		return 2
	} else if (s == "++") || (s == "--") || (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

func isOperator(input string) bool {
	for _, op := range operators {
		if strings.HasPrefix(input, op) {
			return true
		}
	}
	return false
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

func evaluatePostfix(exp string) (string, error) {
	operands := new(Stack)
	//chars := strings.Split(exp, " ")
	for ind, _ := range exp {
		if !isOperator(exp[ind:]) {
			//op, err := strconv.ParseFloat(char, 64)
			//if err != nil {
			//	return 0.0, err
			//}
			//		operands.Push(op)
		} else {
			//		operand2, err := operands.Top()
			//		if err != nil {
			//			return 0.0, err
			//		}
			//		operands.Pop()
			//		operand1, err := operands.Top()
			//		if err != nil {
			//			return 0.0, err
		}
		//		operands.Pop()
		//		calculated, err := calculate(char, operand1, operand2)
		//		if err != nil {
		//			return 0.0, err
		//		}
		//		operands.Push(calculated)
		//	}
	}
	result := operands.Top()
	//if err != nil {
	//	return 0.0, err
	//}
	return result, nil
}

// infixToPostfix convert infix notation to postfix
func infixToPostfix(infix string) (postfix Stack) {
	inputLex := splitLex(infix)
	var sta Stack
	for _, s := range inputLex {
		if isDigit(s) {
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
	//log.Println(postfix)
	return
}

func infixToPostfix1(infix string) string {
	var sta Stack
	var postfix string

	for _, char := range infix {
		opchar := string(char)
		// if scanned character is operand, add it to output string
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' {
			sta.Push(opchar)
		} else if char == ')' {
			for sta.Top() != "(" {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Pop()
		} else {
			for !sta.IsEmpty() && precedence(opchar) <= precedence(sta.Top()) {
				postfix = postfix + sta.Top()
				sta.Pop()
			}
			sta.Push(opchar)
		}
	}
	// Pop all the remaining elements from the stack
	for !sta.IsEmpty() {
		postfix = postfix + sta.Top()
		sta.Pop()
	}
	return postfix
}
func splitLex(input string) (res Stack) {
	var digit string
	var ind = 0
	for range input {
		if ind >= len(input) {
			break
		}
		if input[ind] == '(' || input[ind] == ')' {
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
	//log.Println(res)
	return
}
