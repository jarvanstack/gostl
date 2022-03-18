package stack

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	i := s.GetData()
	fmt.Printf("i: %v\n", i)
}

//结论是 cap 是动态扩容的.
//初始化一个动态数据的 len 和  cap 都是 0
//append 1 之后len and cap 都是1
//append 是追加，就是在之前的基础上最佳，如果你初始化在append是无用的，前面的数据都是初始值，不会冲已开始开始追加.
func TestName(t *testing.T) {
	data := make([]int, 0)
	fmt.Printf("len(data)=%#v\n", len(data))
	fmt.Printf("cap(data)=%#v\n", cap(data))
	data = append(data, 1)
	fmt.Printf("len(data)=%#v\n", len(data))
	fmt.Printf("cap(data)=%#v\n", cap(data))
	data = append(data, 1)

	fmt.Printf("len(data)=%#v\n", len(data))
	fmt.Printf("cap(data)=%#v\n", cap(data))
}

//go没有删除切片元素的方法，我们只有用 copy 赋值可以使用的部分来模拟删删除.
func TestRemove(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("data=%#v\n", data)
	fmt.Printf("%s\n", "头部删除")
	data = data[1:len(data)]
	fmt.Printf("data=%#v\n", data)
	fmt.Printf("%s\n", "尾部删除")
	data = data[:len(data)-1]
	fmt.Printf("data=%#v\n", data)
	fmt.Printf("%s\n", "中间删除，下标为1的元素")
	index := 1
	data = append(data[0:index], data[index+1:len(data)]...)
	fmt.Printf("data=%#v\n", data)
}

//所以需要用指针，这个不像java，自动指针，这个要传递值，就需要指针，而不是用值传递
//go中所有的方法的传递都是值传递，所以要用指针..
func TestStackChangeLengthTest(t *testing.T) {
	stack := NewStack()
	fmt.Printf("stack.topIndex=%d\n", stack.topIndex)
}

//如果 切片的容量大于 1024 是不是只是扩大 25%
//这玩意比java的stack慢，难受.
func TestCopy(t *testing.T) {
	var data []int
	data = make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	fmt.Printf("len(data)=%#v\n", len(data))
	dst := [200]int{}
	copy(dst[:], data[:])
	data = dst[:]
}
