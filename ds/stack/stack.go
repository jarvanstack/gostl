package stack

import (
	"fmt"
	"math"
)

//这是一个工具类，封装需要用到的数据结构
//一个基本栈的实现.
//Pop()
//push()
//IsEmpty()
//Peek()只看不取出来.
//size()

type Stack struct {
	//存放数据的动态数组.
	data []interface{}
	//栈顶的指针
	topIndex int
	//最大长度容量，用来看看是否扩容.
	cap int
}

func NewStack() *Stack {
	stack := &Stack{}
	//初始化初始容量为8
	stack.cap = 8
	stack.topIndex = -1
	stack.data = make([]interface{}, 8)
	return stack
}

//入栈.push
func (this *Stack) Push(item interface{}) error {
	//判空.
	if this == nil {
		return fmt.Errorf("error:%s", "stack is null")
	}
	if this.cap >= math.MaxInt64-1 {
		return fmt.Errorf("error:%s", "index out of int64 of stack array")
	}
	//先判断是否需要扩容
	if this.topIndex >= this.cap-1 {
		if this.cap < 1024 {
			this.cap <<= 1
			temp := make([]interface{}, this.cap)
			copy(temp[:], this.data)
			this.data = temp
		} else {
			this.cap = this.cap + (this.cap >> 2)
			temp := make([]interface{}, this.cap)
			copy(temp[:], this.data)
			this.data = temp
		}
	}
	//赋值
	this.topIndex++
	this.data[this.topIndex] = item
	return nil
}

//出栈pop
func (this *Stack) Pop() interface{} {
	//判空.
	if this == nil {
		fmt.Printf("error:%s", "stack is null")
		return nil
	}
	//如果栈顶没有元素
	if this.topIndex == -1 {
		return nil
	}
	//出栈
	result := this.data[this.topIndex]
	this.data[this.topIndex] = nil
	this.topIndex--
	return result
}

//判断这个栈是否为空.
func (this *Stack) IsEmpty() bool {
	if this == nil || this.topIndex == -1 {
		return true
	}
	return false
}

//只是看看不删除栈顶元素.
func (this *Stack) Peek() interface{} {
	//判空.
	if this == nil {

		return nil
	}
	//如果栈顶没有元素
	if this.topIndex == -1 {
		return nil
	}
	return this.data[this.topIndex]
}

//返回this.topIndex.
func (this *Stack) Size() int {
	if this == nil {
		return -1
	}
	return this.topIndex + 1
}

//返回 data[:this.topIndex+1]) 数据拷贝
//  先进的在左边, 后进来的在右边
func (this *Stack) GetData() (data []interface{}) {
	data = make([]interface{}, this.topIndex+1)
	copy(data, this.data[:this.topIndex+1])
	return data
}
