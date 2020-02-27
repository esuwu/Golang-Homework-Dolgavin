package main

import (
	myStack "./stack"
	"fmt"
	"os"
	"strconv"
)

func pushOutFromStackForPlusMinus(tmpNumberStr *string, item int32, stackOperands *myStack.Stack, stackOperations *myStack.Stack){
	if *tmpNumberStr != "" {
		stackOperands.Push(*tmpNumberStr)
		*tmpNumberStr = ""
	}
	if stackOperations.Peek() == "(" {
		stackOperations.Push(string(item))
	} else {
		stackOperands.Push(makeOperation(stackOperands.Pop(), stackOperands.Pop(), stackOperations.Pop()))
		stackOperations.Push(string(item))
	}
}

func pushOutFromStackForMulAndDiv(tmpNumberStr *string, item int32, stackOperands *myStack.Stack, stackOperations *myStack.Stack){
	if *tmpNumberStr != "" {
		stackOperands.Push(*tmpNumberStr)
		*tmpNumberStr = ""
	}
	if stackOperations.Peek() == "(" {
		stackOperations.Push(string(item))
	} else if stackOperations.Peek() == "*" || stackOperations.Peek() == "/"{
		stackOperands.Push(makeOperation(stackOperands.Pop(), stackOperands.Pop(), stackOperations.Pop()))
		stackOperations.Push(string(item))
	} else {
		stackOperations.Push(string(item))
	}
}

func makeOperation(sec string, fir string, operator string) string {
	first, _ := strconv.Atoi(fir)
	second, _ := strconv.Atoi(sec)
	var result string

	switch operator{
	case "+":
		result = strconv.Itoa(first + second)
	case "-":
		result = strconv.Itoa(first - second)
	case "/":
		result = strconv.Itoa(first / second)
	case "*":
		result = strconv.Itoa(first * second)
	}
	return result
}

func calculateExprInBrackets(tmpNumberStr *string, stackOperands *myStack.Stack, stackOperations *myStack.Stack){
	if *tmpNumberStr != "" {
		stackOperands.Push(*tmpNumberStr)
		*tmpNumberStr = ""
	}
	for stackOperations.Peek() != "("{
		stackOperands.Push(makeOperation(stackOperands.Pop(), stackOperands.Pop(), stackOperations.Pop()))
	}
	stackOperations.Pop()
}

func calculate(expression string) string {
	expression = "(" + expression + ")"
	stackOperations := myStack.New()
	stackOperands := myStack.New()
	var tmpNumberStr string
	for _, item := range expression{
		switch string(item) {
		case "(":
			stackOperations.Push(string(item))
		case "+":
			pushOutFromStackForPlusMinus(&tmpNumberStr, item, stackOperands, stackOperations)
		case "-":
			pushOutFromStackForPlusMinus(&tmpNumberStr, item, stackOperands, stackOperations)
		case "*":
			pushOutFromStackForMulAndDiv(&tmpNumberStr, item, stackOperands, stackOperations)
		case "/":
			pushOutFromStackForMulAndDiv(&tmpNumberStr, item, stackOperands, stackOperations)
		case ")":
			calculateExprInBrackets(&tmpNumberStr, stackOperands, stackOperations)
		default:
			tmpNumberStr += string(item)
		}
	}
	return stackOperands.Pop()
}

func main(){
	expression := os.Args[1]
	fmt.Print(calculate(expression))
}