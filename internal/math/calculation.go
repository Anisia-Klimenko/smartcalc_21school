package math

import (
	"fmt"
	"math"
	"strconv"
)

func evaluatePostfix(exp Stack, xVal float64) (float64, error) {
	operands := new(StackFloat)
	var calculated float64
	var err error
	for _, elem := range exp {
		if isDigit(elem) {
			op, err := strconv.ParseFloat(elem, 64)
			if err != nil {
				return 0.0, err
			}
			operands.Push(op)
		} else if elem == "x" {
			operands.Push(xVal)
		} else {
			if isFunction(elem) {
				operand := operands.Top()
				operands.Pop()
				calculated, err = calculateFunc(elem, operand)
			} else {
				operand2 := operands.Top()
				operands.Pop()
				operand1 := operands.Top()
				operands.Pop()
				calculated, err = calculate(elem, operand1, operand2)
			}
			if err != nil {
				return 0.0, err
			}
			operands.Push(calculated)
		}
	}
	result := operands.Top()
	return result, nil
}

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
	}
	return result, nil
}

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
