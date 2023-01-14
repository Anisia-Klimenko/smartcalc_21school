package math

import (
	"math"
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		input string
		xVal  float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "empty", args: args{input: "", xVal: 0}, wantResult: "0"},
		{name: "error )", args: args{input: "sqrt(4", xVal: 0}, wantResult: "error"},
		{name: "1/3", args: args{input: "1/3", xVal: 0}, wantResult: "0.3333333"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Calculate(tt.args.input, tt.args.xVal); gotResult != tt.wantResult {
				t.Errorf("Calculate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestStackFloat_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		st   StackFloat
		want bool
	}{
		{name: "empty StackFloat", st: StackFloat{}, want: true},
		{name: "not empty StackFloat", st: StackFloat{1, 2, 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackFloat_Pop(t *testing.T) {
	tests := []struct {
		name string
		st   StackFloat
		want bool
	}{
		{name: "empty StackFloat", st: StackFloat{}, want: false},
		{name: "not empty StackFloat", st: StackFloat{1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackFloat_Push(t *testing.T) {
	type args struct {
		str float64
	}
	tests := []struct {
		name string
		st   StackFloat
		args args
		want StackFloat
	}{
		{name: "empty StackFloat", st: StackFloat{}, args: args{str: 1}, want: StackFloat{1}},
		{name: "not empty StackFloat", st: StackFloat{1}, args: args{str: 2}, want: StackFloat{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.st.Push(tt.args.str)
		})
		if len(tt.st) != len(tt.want) || tt.st[len(tt.st)-1] != tt.args.str {
			t.Errorf("Push() = %v, want %v", tt.st, tt.want)
		}
	}
}

func TestStackFloat_Top(t *testing.T) {
	tests := []struct {
		name string
		st   StackFloat
		want float64
	}{
		{name: "empty StackFloat", st: StackFloat{}, want: 0.0},
		{name: "not empty StackFloat", st: StackFloat{1, 2, 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		st   Stack
		want bool
	}{
		{name: "empty Stack", st: Stack{}, want: true},
		{name: "not empty Stack", st: Stack{"a", "b", "c"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		st   Stack
		want bool
	}{
		{name: "empty Stack", st: Stack{}, want: false},
		{name: "not empty Stack", st: Stack{"a", "b", "c"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		st   Stack
		args args
		want Stack
	}{
		{name: "empty Stack", st: Stack{}, args: args{str: "a"}, want: Stack{"a"}},
		{name: "not empty Stack", st: Stack{"a"}, args: args{str: "b"}, want: Stack{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.st.Push(tt.args.str)
		})
		if len(tt.st) != len(tt.want) || tt.st[len(tt.st)-1] != tt.args.str {
			t.Errorf("Push() = %v, want %v", tt.st, tt.want)
		}
	}
}

func TestStack_Top(t *testing.T) {
	tests := []struct {
		name string
		st   Stack
		want string
	}{
		{name: "empty Stack", st: Stack{}, want: ""},
		{name: "not empty Stack", st: Stack{"a", "b", "c"}, want: "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate(t *testing.T) {
	type args struct {
		operator string
		operand1 float64
		operand2 float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "1+2", args: args{operator: "+", operand1: 1, operand2: 2}, want: 3, wantErr: false},
		{name: "1-2", args: args{operator: "-", operand1: 1, operand2: 2}, want: -1, wantErr: false},
		{name: "1*2", args: args{operator: "*", operand1: 1, operand2: 2}, want: 2, wantErr: false},
		{name: "1*0", args: args{operator: "*", operand1: 1, operand2: 0}, want: 0, wantErr: false},
		{name: "1/2", args: args{operator: "/", operand1: 1, operand2: 2}, want: 0.5, wantErr: false},
		{name: "1/0", args: args{operator: "/", operand1: 1, operand2: 0}, want: +math.Inf(1), wantErr: false},
		{name: "2^3", args: args{operator: "^", operand1: 2, operand2: 3}, want: 8, wantErr: false},
		{name: "-2^3", args: args{operator: "^", operand1: -2, operand2: 3}, want: -8, wantErr: false},
		{name: "5^-2", args: args{operator: "^", operand1: 5, operand2: -2}, want: 0.04, wantErr: false},
		{name: "4^0.5", args: args{operator: "^", operand1: 4, operand2: 0.5}, want: 2, wantErr: false},
		{name: "1mod2", args: args{operator: "mod", operand1: 1, operand2: 2}, want: 1, wantErr: false},
		{name: "4mod2", args: args{operator: "mod", operand1: 4, operand2: 2}, want: 0, wantErr: false},
		{name: "3mod2", args: args{operator: "mod", operand1: 3, operand2: 2}, want: 1, wantErr: false},
		{name: "3mod-2", args: args{operator: "mod", operand1: 3, operand2: -2}, want: 1, wantErr: false},
		{name: "-3mod2", args: args{operator: "mod", operand1: -3, operand2: 2}, want: -1, wantErr: false},
		{name: "-3mod-2", args: args{operator: "mod", operand1: -3, operand2: -2}, want: -1, wantErr: false},
		{name: "10e2", args: args{operator: "e", operand1: 10, operand2: 2}, want: 1000, wantErr: false},
		{name: "-10e2", args: args{operator: "e", operand1: 10, operand2: 2}, want: 1000, wantErr: false},
		{name: "-100e-2", args: args{operator: "e", operand1: -100, operand2: -2}, want: -1, wantErr: false},
		{name: "10e0", args: args{operator: "e", operand1: 10, operand2: 0}, want: 10, wantErr: false},
		{name: "-10k2", args: args{operator: "k", operand1: 10, operand2: 2}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculate(tt.args.operator, tt.args.operand1, tt.args.operand2)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFunc(t *testing.T) {
	type args struct {
		function string
		operand  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "sqrt(4)", args: args{function: "sqrt", operand: 4}, want: 2, wantErr: false},
		{name: "sqrt(0)", args: args{function: "sqrt", operand: 0}, want: 0, wantErr: false},
		{name: "sqrt(-4)", args: args{function: "sqrt", operand: -4}, want: math.NaN(), wantErr: false},
		{name: "sin(0)", args: args{function: "sin", operand: 0}, want: 0, wantErr: false},
		{name: "asin(0)", args: args{function: "asin", operand: 0}, want: 0, wantErr: false},
		{name: "cos(0)", args: args{function: "cos", operand: 0}, want: 1, wantErr: false},
		{name: "acos(1)", args: args{function: "acos", operand: 1}, want: 0, wantErr: false},
		{name: "tan(0)", args: args{function: "tan", operand: 0}, want: 0, wantErr: false},
		{name: "atan(0)", args: args{function: "atan", operand: 0}, want: 0, wantErr: false},
		{name: "ln(1)", args: args{function: "ln", operand: 1}, want: 0, wantErr: false},
		{name: "ln(-1)", args: args{function: "ln", operand: -1}, want: math.NaN(), wantErr: false},
		{name: "log(10)", args: args{function: "log", operand: 10}, want: 1, wantErr: false},
		{name: "log(1)", args: args{function: "log", operand: 1}, want: 0, wantErr: false},
		{name: "log(100)", args: args{function: "log", operand: 100}, want: 2, wantErr: false},
		{name: "log(-10)", args: args{function: "log", operand: -10}, want: math.NaN(), wantErr: false},
		{name: "abs(4)", args: args{function: "abs", operand: 4}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateFunc(tt.args.function, tt.args.operand)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if math.IsNaN(tt.want) != math.IsNaN(got) && math.IsNaN(got) {
				t.Errorf("calculateFunc() got = %v, want %v", got, tt.want)
			} else if got != tt.want && !math.IsNaN(got) {
				t.Errorf("calculateFunc() got = %v, want %v", got, tt.want)

			}
		})
	}
}

func Test_evaluatePostfix(t *testing.T) {
	type args struct {
		exp  Stack
		xVal float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "empty0", args: args{exp: Stack{}, xVal: 0}, want: 0, wantErr: false},
		{name: "empty1", args: args{exp: Stack{}, xVal: 1}, want: 0, wantErr: false},
		{name: "1", args: args{exp: Stack{"1"}, xVal: 0}, want: 1, wantErr: false},
		{name: "-1", args: args{exp: Stack{"-1"}, xVal: 0}, want: -1, wantErr: false},
		{name: "+1", args: args{exp: Stack{"1"}, xVal: 0}, want: 1, wantErr: false},
		{name: "1+2", args: args{exp: Stack{"1", "2", "+"}, xVal: 0}, want: 3, wantErr: false},
		{name: "1+-2", args: args{exp: Stack{"1", "-2", "+"}, xVal: 0}, want: -1, wantErr: false},
		{name: "1-+2", args: args{exp: Stack{"1", "2", "-"}, xVal: 0}, want: -1, wantErr: false},
		{name: "sin(x)", args: args{exp: Stack{"x", "sin"}, xVal: 0}, want: 0, wantErr: false},
		{name: "-cos(-x)", args: args{exp: Stack{"-x", "cos", "-"}, xVal: 0}, want: -1, wantErr: false},
		{name: "-co(-x)", args: args{exp: Stack{"-x", "co", "-"}, xVal: 0}, want: -1, wantErr: true},
		{name: "-10e2*sin(2/sqrt(4-2)*log(3e2+3^4-6*(4+-5)))", args: args{exp: Stack{"-10", "2", "e", "2", "4", "2", "-", "sqrt", "/", "2", "3", "e", "3", "4", "^", "+", "6", "4", "-5", "+", "*", "-", "log", "*", "sin", "*"}, xVal: 0}, want: 999.8402408588436, wantErr: false},
		{name: "-10ex*sin(x/sqrt(4-x)*log(3ex+3^4-6*(4+-5)))", args: args{exp: Stack{"-10", "2", "e", "2", "4", "2", "-", "sqrt", "/", "2", "3", "e", "3", "4", "^", "+", "6", "4", "-5", "+", "*", "-", "log", "*", "sin", "*"}, xVal: 2}, want: 999.8402408588436, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluatePostfix(tt.args.exp, tt.args.xVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluatePostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.wantErr {
				t.Errorf("evaluatePostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infixToPostfix(t *testing.T) {
	type args struct {
		infix string
	}
	tests := []struct {
		name        string
		args        args
		wantPostfix Stack
	}{
		{name: "empty", args: args{infix: ""}, wantPostfix: Stack{}},
		{name: "1", args: args{infix: "1"}, wantPostfix: Stack{"1"}},
		{name: "-1", args: args{infix: "-1"}, wantPostfix: Stack{"-1"}},
		{name: "+1", args: args{infix: "+1"}, wantPostfix: Stack{"1"}},
		{name: "1+2", args: args{infix: "1+2"}, wantPostfix: Stack{"1", "2", "+"}},
		{name: "1+-2", args: args{infix: "1+-2"}, wantPostfix: Stack{"1", "-2", "+"}},
		{name: "1-+2", args: args{infix: "1-+2"}, wantPostfix: Stack{"1", "2", "-"}},
		{name: "sin(1)", args: args{infix: "sin(1)"}, wantPostfix: Stack{"1", "sin"}},
		{name: "sin(-1)", args: args{infix: "sin(-1)"}, wantPostfix: Stack{"-1", "sin"}},
		{name: "-sin(-1)", args: args{infix: "-sin(-1)"}, wantPostfix: Stack{"-1", "sin", "-"}},
		{name: "sin(-x)", args: args{infix: "sin(-x)"}, wantPostfix: Stack{"-x", "sin"}},
		{name: "1+2)", args: args{infix: "1+2)"}, wantPostfix: Stack{"error"}},
		{name: "-10e2*sin(2/sqrt(4-2)*log(2e3+3^4-6*(4+-5)))",
			args: args{infix: "-10e2*sin(2/sqrt(4-2)*log(2e3+3^4-6*(4+-5)))"},
			wantPostfix: Stack{"-10", "2", "e", "2", "4", "2", "-", "sqrt", "/", "2", "3", "e", "3",
				"4", "^", "+", "6", "4", "-5", "+", "*", "-", "log", "*", "sin", "*"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPostfix := infixToPostfix(tt.args.infix)
			if gotPostfix.IsEmpty() && tt.wantPostfix.IsEmpty() {
				return
			}
			if !reflect.DeepEqual(gotPostfix, tt.wantPostfix) {
				t.Errorf("infixToPostfix() = %v, want %v", gotPostfix, tt.wantPostfix)
			}
		})
	}
}

func Test_isDigit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "5", args: args{input: "5"}, want: true},
		{name: "5.2", args: args{input: "5.2"}, want: true},
		{name: "-5", args: args{input: "-5"}, want: true},
		{name: "+5", args: args{input: "+5"}, want: true},
		{name: "a", args: args{input: "a"}, want: false},
		{name: "5a", args: args{input: "5a"}, want: false},
		{name: "log(5)", args: args{input: "log(5)"}, want: false},
		{name: "/", args: args{input: "/"}, want: false},
		{name: "-", args: args{input: "-"}, want: false},
		{name: "empty", args: args{input: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.input); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFunction(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "sqrt", args: args{input: "sqrt"}, want: true},
		{name: "sin", args: args{input: "sin"}, want: true},
		{name: "asin", args: args{input: "asin"}, want: true},
		{name: "cos", args: args{input: "cos"}, want: true},
		{name: "acos", args: args{input: "acos"}, want: true},
		{name: "tan", args: args{input: "tan"}, want: true},
		{name: "atan", args: args{input: "atan"}, want: true},
		{name: "log", args: args{input: "log"}, want: true},
		{name: "ln", args: args{input: "ln"}, want: true},
		{name: "e", args: args{input: "e"}, want: false},
		{name: "mod", args: args{input: "mod"}, want: false},
		{name: "abs", args: args{input: "abs"}, want: false},
		{name: "empty", args: args{input: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFunction(tt.args.input); got != tt.want {
				t.Errorf("isFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_precedence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "+", args: args{s: "+"}, want: 1},
		{name: "-", args: args{s: "-"}, want: 1},
		{name: "*", args: args{s: "*"}, want: 2},
		{name: "/", args: args{s: "/"}, want: 2},
		{name: "mod", args: args{s: "mod"}, want: 2},
		{name: "^", args: args{s: "^"}, want: 3},
		{name: "e", args: args{s: "e"}, want: 3},
		{name: "sin", args: args{s: "sin"}, want: 3},
		{name: "log", args: args{s: "log"}, want: 3},
		{name: "_", args: args{s: "_"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := precedence(tt.args.s); got != tt.want {
				t.Errorf("precedence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitLex(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantRes Stack
	}{
		{name: "1+2", args: args{input: "1+2"}, wantRes: Stack{"1", "+", "2"}},
		{name: "1+-2", args: args{input: "1+-2"}, wantRes: Stack{"1", "+", "-2"}},
		{name: "1-+2", args: args{input: "1-+2"}, wantRes: Stack{"1", "-", "2"}},
		{name: "-1", args: args{input: "-1"}, wantRes: Stack{"-1"}},
		{name: "+1", args: args{input: "+1"}, wantRes: Stack{"1"}},
		{name: "sin(1)", args: args{input: "sin(1)"}, wantRes: Stack{"sin", "(", "1", ")"}},
		{name: "sin(-1)", args: args{input: "sin(-1)"}, wantRes: Stack{"sin", "(", "-1", ")"}},
		{name: "-sin(-1)", args: args{input: "-sin(-1)"}, wantRes: Stack{"-", "sin", "(", "-1", ")"}},
		{name: "sin(-x)", args: args{input: "sin(-x)"}, wantRes: Stack{"sin", "(", "-x", ")"}},
		{name: "", args: args{input: ""}, wantRes: Stack{}},
		{name: "-10e2*sin(2/sqrt(4-2)*log(2e3+3^4-6*(4+-5)))", args: args{input: "-10e2*sin(2/sqrt(4-2)*log(2e3+3^4-6*(4+-5)))"},
			wantRes: Stack{"-10", "e", "2", "*", "sin", "(", "2", "/", "sqrt", "(", "4", "-", "2", ")", "*", "log", "(", "2",
				"e", "3", "+", "3", "^", "4", "-", "6", "*", "(", "4", "+", "-5", ")", ")", ")"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := splitLex(tt.args.input)
			if gotRes.IsEmpty() && tt.wantRes.IsEmpty() {
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("splitLex() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
