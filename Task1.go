package main

import (
	"fmt"
	"strconv"
)

type stack []string
type intStack int

func (st *stack) isempty() bool {
	return len(*st) == 0;
}

func (st *stack) push(str string) {
	*st = append(*st, str);
}

func (st *stack) pop() bool {
	if(st.isempty()){
		return false;
	} else {
		index := len(*st) - 1;
		*st = (*st)[:index];
		return true;
	}
}

func (st *stack) top() string {
	if( st.isempty() ) {
		return "";
	} else {
		index := len(*st) - 1;
		element := (*st)[index];
		return element;
	}
}

func precedence(str string) int {
	if(str == "^") {
		return 3;
	} else if(str == "/" || str == "*") {
		return 2;
	} else if(str == "+" || str == "-") {
		return 1;
	} else {
		return -1;
	}
}

func applyOp(a int, b int, c string) int {

	var val int;

	switch (c) {
		case "+" : val =  a + b;
		case "-" : val =  a - b;
		case "*" : val =  a * b;
		case "/" : val =  a / b;
	}

	return val;
}



func main() {
	var equ string;
	
	equ = "2+3*(2^3-5)^(2+1*2)-4";
	fmt.Print(equ);

	
	var operator stack;
	var operand stack;

	for i:=0; i < len(equ); i++ {
		
		opchar := string(equ[i]);

		if (opchar == "(") {
			fmt.Print(opchar);
			operator.push("(");
		} else if (opchar >= "0" && opchar <= "9") {
			fmt.Print(opchar);
			var val int = 0;
			
			for j := i; j < len(equ); j++ {
				
				var tmp int;
				opchar := string(equ[i]);
				strVar := opchar;
				if (opchar < "0" && opchar > "9") {
					break;
				}
				intVar, err := strconv.Atoi(strVar);
				fmt.Print(strVar);
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				tmp = intVar;

				val = (val * 10) + tmp;
			}

			stringVal := strconv.Itoa(val);
			operand.push(stringVal);
			i--;

		} else if (opchar == ")") {

			for !operator.isempty() && opchar != "(" {
				
				var tmp1 int;
				opchar1 := operand.top();
				operand.pop();
				strVar1 := opchar1;
				intVar1, err := strconv.Atoi(strVar1);
				fmt.Print(strVar1);
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				tmp1 = intVar1;

				var tmp2 int;
				opchar2 := operand.top();
				operand.pop();
				strVar2 := opchar2;
				intVar2, err := strconv.Atoi(strVar2);
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				tmp2 = intVar2;

				opchar3 := operator.top();
				operator.pop();

				var now int = applyOp(tmp1, tmp2, opchar3);
				stringVal := strconv.Itoa(now);
				operand.push(stringVal);
			}

			if (!operator.isempty()) {
				operator.pop();
			}
		} else {


			for !operator.isempty() && precedence(operator.top()) >= precedence(opchar) {

				var tmp1 int;
				opchar1 := operand.top();
				operand.pop();
				strVar1 := opchar1;
				intVar1, err := strconv.Atoi(strVar1);
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				tmp1 = intVar1;

				var tmp2 int;
				opchar2 := operand.top();
				operand.pop();
				strVar2 := opchar2;
				intVar2, err := strconv.Atoi(strVar2);
				if err != nil {
					fmt.Println("Error during conversion")
					return
				}
				tmp2 = intVar2;

				opchar3 := operator.top();
				operator.pop();

				var now int = applyOp(tmp1, tmp2, opchar3);
				stringVal := strconv.Itoa(now);
				operand.push(stringVal);
			}

			operator.push(opchar);



		}


		for !operator.isempty() && precedence(operator.top()) >= precedence(opchar) {

			var tmp1 int;
			opchar1 := operand.top();
			operand.pop();
			strVar1 := opchar1;
			intVar1, err := strconv.Atoi(strVar1);
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			tmp1 = intVar1;

			var tmp2 int;
			opchar2 := operand.top();
			operand.pop();
			strVar2 := opchar2;
			intVar2, err := strconv.Atoi(strVar2);
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			tmp2 = intVar2;

			opchar3 := operator.top();
			operator.pop();

			var now int = applyOp(tmp1, tmp2, opchar3);
			stringVal := strconv.Itoa(now);
			operand.push(stringVal);
		}


		fmt.Print(operand.top());
	}

}
