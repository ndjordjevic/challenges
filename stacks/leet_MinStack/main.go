package main

import "fmt"

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
)

type MinStack struct {
	data []int
	min  int
}

func Constructor() MinStack {
	return MinStack{
		[]int{}, MaxInt,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	l := len(this.data)
	this.data = this.data[:l-1]

	this.min = MaxInt
	for _, e := range this.data {
		if e < this.min {
			this.min = e
		}
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(3)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
