package main

import (
	"fmt"
)

type stack []string

// Function to check if the stack has elements of the stack is empty
func (st *stack) isempty() bool {
	return len(*st) == 0
}

// Remove top element of stack or flase if the stack is already empty
func (st *stack) pop() bool {
	if st.isempty() {
		return false
	} else {
		index := len(*st) - 1
		*st = (*st)[:index]
		return true
	}
}

// Return top element of stack or false if the stack is already empty
func (st *stack) top() string {
	if st.isempty() {
		return ""
	} else {
		index := len(*st) - 1
		element := (*st)[index]
		return element
	}
}

// Push a new value onto the stack
func (st *stack) push(str string) {
	*st = append(*st, str)
}

// Function for integer to string conversion
func int_to_string(number int) string {
	var string_representation string = fmt.Sprint(number)
	return string_representation
}

// Function for string to integer conversion
func string_to_int(string_representation string) int {
	var number int = 0
	for i := 0; i < len(string_representation); i++ {
		number *= 10
		var digit int = int(string_representation[i] - '0')
		number += digit
	}
	return number
}

// Function to return precedence of operators
func prec(s string) int {
	if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

// Function to convert infix equation to postfix equation
func convertToPostfix(infix string) string {
	var st stack
	var postfix string
	for _, char := range infix {
		opchar := string(char)
		// if scanned character is operand and add it to the postfix string
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			postfix = postfix + opchar
		} else if char == '(' { // if scanned character is open bracket push to the stack
			st.push(opchar)
		} else if char == ')' {
			for st.top() != "(" { // if scanned character is open closing bracket add all previous elements to the string
				postfix = postfix + st.top()
				st.pop()
			}
			st.pop()
		} else {
			for !st.isempty() && prec(opchar) <= prec(st.top()) { // While the stack is not empty add all previous elements to the string
				postfix = postfix + st.top()
				st.pop()
			}
			st.push(opchar)
		}
	}

	// Pop all the remaining elements from the stack
	for !st.isempty() {
		postfix = postfix + st.top()
		st.pop()
	}

	return postfix
}

// Function to find the result of the operation on two numbers
func applyOp(a int, b int, c string) int {

	var val int

	switch c {
	case "+":
		val = a + b
	case "-":
		val = a - b
	case "*":
		val = a * b
	case "/":
		val = a / b
	}

	return val
}

func evaluate(infix string) {

	var equ string

	equ = infix

	var operator stack
	var operand stack

	for i := 0; i < len(equ); i++ {

		opchar := string(equ[i])

		// Push the element into operand stack if it is a digit
		if opchar >= "0" && opchar <= "9" {

			for j := i; j < len(equ); j++ {

				opchar := string(equ[j])
				var val int
				val = string_to_int(opchar)
				if val < 0 || val > 9 {
					break
				}

				i++
				operand.push(opchar)
			}
			i--

		} else {
			// Push the element into operator stack if it is a digit
			operator.push(opchar)
			// Now if there is a operator then do the operation with the two operands
			for !operator.isempty() {
				var tmp1 int
				opchar1 := operand.top()
				operand.pop()
				tmp1 = string_to_int(opchar1)

				var tmp2 int
				opchar2 := operand.top()
				operand.pop()
				tmp2 = string_to_int(opchar2)

				opchar3 := operator.top()
				operator.pop()

				var now int = applyOp(tmp2, tmp1, opchar3)

				tmp3 := int_to_string(now)
				operand.push(tmp3)
			}
		}
	}

	fmt.Print("Ans: ", operand.top())
}

func main() {
	infix := "((1+2)*(4/2))"
	postfix := convertToPostfix(infix)
	evaluate(postfix)
}
