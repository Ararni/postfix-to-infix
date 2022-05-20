package lab2

import (
	"errors"
	"fmt"
	"strings"
)

type StackNode struct {
	data string
	next *StackNode
}

func getStackNode(data string, top *StackNode) *StackNode {
	return &StackNode{
		data,
		top,
	}
}

type MyStack struct {
	top   *StackNode
	count int
}

func getMyStack() *MyStack {
	return &MyStack{
		nil,
		0,
	}
}

func (this MyStack) size() int {
	return this.count
}
func (this MyStack) isEmpty() bool {
	if this.size() > 0 {
		return false
	} else {
		return true
	}
}

func (this *MyStack) push(data string) {
	this.top = getStackNode(data, this.top)
	this.count++
}

func (this *MyStack) pop() string {
	var temp string = ""
	if this.isEmpty() == false {
		temp = this.top.data
		this.top = this.top.next
		this.count--
	}
	return temp
}

func (this MyStack) peek() string {
	if !this.isEmpty() {
		return this.top.data
	} else {
		return ""
	}
}

func isOperator(text string) bool {
	if text == "+" || text == "-" ||
		text == "*" || text == "/" ||
		text == "^" {
		return true
	}
	return false
}

func isOperands(text string) bool {
	return !isOperator(text)
}

func PostfixToInfix(input string) (string, error) {
	var inputStr = strings.Split(strings.TrimSpace(input), " ")
	var size int = len(inputStr)
	var s *MyStack = getMyStack()
	var auxiliary string = ""
	var op1 string = ""
	var op2 string = ""
	var notValid bool = true
	var notEmpty bool = true
	if size == 1 {
		notEmpty = false
	}
	fmt.Println(size)
	for i := 0; i < size && notValid && notEmpty; i++ {
		if isOperator(inputStr[i]) {
			if s.size() > 1 {
				op1 = s.pop()
				op2 = s.pop()
				auxiliary = "(" + op2 + string(inputStr[i]) + op1 + ")"
				s.push(auxiliary)
			} else {
				notValid = false
			}
		} else if isOperands(inputStr[i]) {
			auxiliary = string(inputStr[i])
			s.push(auxiliary)
		} else {
			notValid = false
		}
	}
	if notEmpty == false {
		return "", errors.New("Empty postfix")
	}
	if notValid == false {
		return "", errors.New("Invalid postfix")
	}
	return s.pop(), nil
}
