package arrays

import "fmt"

func StackOperation() {
	//intialise the array size
	arr := StackArray(5)
	arr.Push(10)
	arr.Push(20)
	arr.Push(30)
	arr.Pop()
	arr.Pop()
	arr.Pop()
	arr.Pop()

}

type Stack struct {
	data [5]int
	top  int
	cap  int
}

func StackArray(capacity int) *Stack {
	return &Stack{
		top: -1,
		cap: capacity,
	}
}

func (a *Stack) Push(element int) {
	if !a.isFull() {
		a.data[a.top+1] = element
		a.top++
		a.display()
	}

}

func (a *Stack) Pop() {
	if !a.isEmpty() {
		a.top--
		a.display()
	}
}

func (a *Stack) display() {
	fmt.Print("[")
	for i := a.top; i >= 0; i-- {
		fmt.Print(" ", a.data[i])
	}
	fmt.Print(" ]")
	fmt.Println()
}

func (a *Stack) isFull() bool {
	if a.top == a.cap-1 {
		fmt.Println("Stack is full")
		return true
	}
	return false
}

func (a *Stack) isEmpty() bool {
	if a.top == -1 {
		fmt.Println("Stack is empty")
		return true
	}
	return false
}

func (a *Stack) Peek() {
	if !a.isEmpty() {
		fmt.Println(a.data[a.top])
	}
}
