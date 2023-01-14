package math

import (
	"fmt"
	"math"
	"strconv"
)

// evaluatePostfix calculates the value of an expression exp in postfix notation, and sets
// x value to xVal. Function is using README algorithm.
func evaluatePostfix(exp Stack, xVal float64) (float64, error) {
	// operands stack contains operands from exp, the number of operands and their values
	// changes during loop, the last remaining element in stack is the result of the expression
	operands := new(StackFloat)
	// calculated is current result of calculation operator or function
	var calculated float64
	var err error
	for _, elem := range exp {
		if isDigit(elem) {
			// If there is a number on the stack, move it to the operands
			op, _ := strconv.ParseFloat(elem, 64)
			operands.Push(op)
		} else if elem == "x" {
			// If there is x on the stack, move its value to the operands
			operands.Push(xVal)
		} else if elem == "-x" {
			// If there is negative x on the stack, move its negative value to the operands
			operands.Push(-xVal)
		} else {
			if isFunction(elem) {
				// If there is a function on the stack, take the last operand
				// from the stack as the function argument and calculate the value
				operand := operands.Top()
				operands.Pop()
				calculated, err = calculateFunc(elem, operand)
			} else {
				// If there is an operator on the stack, take the last 2 operands
				// from the stack and calculate the value
				operand2 := operands.Top()
				operands.Pop()
				operand1 := operands.Top()
				operands.Pop()
				calculated, err = calculate(elem, operand1, operand2)
			}
			if err != nil {
				return 0.0, err
			}
			// Save calculated result as last operand in stack
			operands.Push(calculated)
		}
	}
	// Result of the expression is the last operand in the stack
	result := operands.Top()
	return result, nil
}

// calculateFunc calculates the value of a function, substituting the operand as an argument
func calculateFunc(function string, operand float64) (float64, error) {
	result := 0.0
	switch function {
	case "sqrt":
		result = math.Sqrt(operand)
	case "sin":
		result = math.Sin(operand)
	case "asin":
		result = math.Asin(operand)
	case "cos":
		result = math.Cos(operand)
	case "acos":
		result = math.Acos(operand)
	case "tan":
		result = math.Tan(operand)
	case "atan":
		result = math.Atan(operand)
	case "ln":
		result = math.Log(operand)
	case "log":
		result = math.Log10(operand)
	default:
		return 0.0, fmt.Errorf("invalid operator")
	}
	return result, nil
}

// calculate calculates the result of the operator on the operand1 and operand2
func calculate(operator string, operand1, operand2 float64) (float64, error) {
	result := 0.0
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		result = operand1 / operand2
	case "^":
		result = math.Pow(operand1, operand2)
	case "mod":
		result = operand1 - (operand2 * float64(int(operand1)/int(operand2)))
	case "e":
		result = operand1 * math.Pow(10, operand2)
	default:
		return 0.0, fmt.Errorf("invalid operator")
	}
	return result, nil
}
