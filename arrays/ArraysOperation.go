package arrays

import (
	"fmt"
)

type ArraysADT struct {
	data [5]int
	size int
	cap  int
}

func NewArray(capacity int) *ArraysADT {
	return &ArraysADT{
		size: 0,
		cap:  capacity,
	}
}
func ArraysOperation() {

	arr := NewArray(5)

	arr.insert(0, 5)
	arr.insert(1, 10)
	arr.insert(2, 7)
	arr.insert(1, -1)
	fmt.Println("Delete")
	arr.delete(1)
	arr.get(1)
	arr.update(0, -1)
	arr.search(10)
	arr.search(-11)
}

func (a *ArraysADT) insert(index, element int) {
	if a.size >= a.cap {
		fmt.Println("Array is full")
	}
	if index < 0 || index > a.size {
		fmt.Println("Index out of bound")
	}

	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = element
	a.size++
	a.display()
}

func (a *ArraysADT) delete(index int) {
	if index < 0 || index >= a.size {
		fmt.Println("Index out of bound")
	}
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	a.display()
}

func (a *ArraysADT) get(index int) {
	if index < 0 || index > a.cap {
		fmt.Println("Index out of bound")
	}
	fmt.Println(fmt.Sprintf("The element in the index: %d is %d", index, a.data[index]))
}

func (a *ArraysADT) Size() int {
	return a.size
}

func (a *ArraysADT) update(index, element int) {
	if index < 0 || index > a.cap {
		fmt.Println("Index out of bound")
	}
	a.data[index] = element
	a.display()
}

func (a *ArraysADT) display() {
	fmt.Println(a.data[:a.size])
}

func (a *ArraysADT) search(element int) {
	found := false
	for i := 0; i <= a.size; i++ {
		if a.data[i] == element {
			fmt.Println(fmt.Sprintf("The element found in the index %d", i))
			found = true
		}
	}
	if !found {
		fmt.Println("value not found")
	}
}
