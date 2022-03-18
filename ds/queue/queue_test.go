package queue

import (
	"container/list"
	"fmt"
	"testing"
)

func TestFunctionList(t *testing.T) {
	l := list.New()
	l.PushBack("n1")
	l.PushBack("n2")
	node := l.Front()
	for node != nil {
		fmt.Printf("node.Value: %v\n", node.Value)
		node = node.Next()
	}
}

//测试push
//搞错了，新来的元素因该在左边，所以我们必须维护一个left节点的指针.用来push
func TestQueue_Push(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 5; i++ {
		_ = queue.Push(i)
	}
	//打印
	for !queue.IsEmpty() {
		fmt.Printf("queue.Pop()=%#v\n", queue.Pop())
	}
}

//测试peek 和 pop
func TestQueue_Peek(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 5; i++ {
		_ = queue.Push(i)
	}
	//打印
	for !queue.IsEmpty() {
		fmt.Printf("queue.Peek()=%#v\n", queue.Peek())
		fmt.Printf("queue.Pop()=%#v\n", queue.Pop())
	}
}
func TestQueue_Size(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 5; i++ {
		_ = queue.Push(i)
	}
	//打印
	for !queue.IsEmpty() {
		fmt.Printf("queue.Size()=%#v\n", queue.Size())
		_ = queue.Pop()
	}
}

//压测 1000w (1.38s)
func TestQueue_Press(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 1000_0000; i++ {
		_ = queue.Push(i)
	}
}
