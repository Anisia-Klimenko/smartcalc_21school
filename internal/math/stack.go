package math

type Stack []string

// IsEmpty checks if stack is empty
func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

// Push pushes a new value onto the stack
func (st *Stack) Push(str string) {
	*st = append(*st, str)
}

// Pop removes top element of stack. Returns false if stack is empty.
func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1 // Get the index of top most element.
		*st = (*st)[:index]   // Remove it from the stack by slicing it off.
		return true
	}
}

// Top returns top element of stack. Returns false if stack is empty.
func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	} else {
		index := len(*st) - 1   // Get the index of top most element.
		element := (*st)[index] // Index onto the slice and obtain the element.
		return element
	}
}

type StackFloat []float64

// IsEmpty checks if stack is empty
func (st *StackFloat) IsEmpty() bool {
	return len(*st) == 0
}

// Push pushes a new value onto the stack
func (st *StackFloat) Push(str float64) {
	*st = append(*st, str)
}

// Pop removes top element of stack. Returns false if stack is empty.
func (st *StackFloat) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1 // Get the index of top most element.
		*st = (*st)[:index]   // Remove it from the stack by slicing it off.
		return true
	}
}

// Top returns top element of stack. Returns false if stack is empty.
func (st *StackFloat) Top() float64 {
	if st.IsEmpty() {
		return 0.0
	} else {
		index := len(*st) - 1   // Get the index of top most element.
		element := (*st)[index] // Index onto the slice and obtain the element.
		return element
	}
}
