package queue

import "fmt"

//利用单向向左链表做一个简单的队列queue
//先进入的是左边，后进入的是右边.
//后进入的链表的左指针指向之前的top即可，我们保留的话就保留top(right)指针即可.
//push()
//pop()
//isEmpty()
//peek()
//size()

//向左的单向链表,
type linkedNode struct {
	data interface{}
	left *linkedNode
}

//队列的解构.
type Queue struct {
	//最右边的的元素相当于栈顶.
	right *linkedNode
	//最左边元素的指针，新增加的元素在左边.
	left *linkedNode
	//队列链表的长度.
	size int
}

//创建一个长度为0的队列.
func NewQueue() *Queue {
	return &Queue{right: nil, left: nil, size: 0}
}

//返回true如果队列长度为0
func (this *Queue) IsEmpty() bool {
	return this == nil || this.size <= 0
}

//Push 队列链表添加一个元素
//新添加的元素在最左边，所以要维护一个最左边元素的指针.
func (this *Queue) Push(in interface{}) error {
	//不允许添加空的内容.
	if in == nil {
		return fmt.Errorf("error:%s", "Push不能为空")
	}
	//插入的节点初始化
	node := &linkedNode{data: in}
	//如果队列为空
	if this.IsEmpty() {
		this.left = node
		this.right = node
	} else {
		//如果队列不为空
		//将插入节点插入到队列的最左边
		//左节点左指针指向插入节点
		this.left.left = node
		//插入节点替换有左节点的位置
		this.left = node
	}
	//长度 +1
	this.size++
	return nil
}

//Pop 对列链表中弹出一个元素
//最右节点开始弹出来.
func (this *Queue) Pop() interface{} {
	//队列中没有元素的时候.
	if this.size <= 0 {
		fmt.Printf("error:%s", "队列中的元素为空")
		return nil
	}
	//队列中有元素的时候
	//删除right元素
	right := this.right
	node := this.right.left
	//有后面一个元素继承right的位置
	this.right = node
	//长度 -1
	this.size--
	//返回的是元素，而不是，链表节点.
	return right.data
}

//Size()返回队列链表中的元素的数量.
func (this *Queue) Size() int {
	return this.size
}

//peek()
func (this *Queue) Peek() interface{} {
	//队列中没有元素的时候.
	if this.size <= 0 {
		fmt.Printf("error:%s", "队列中的元素为空")
		return nil
	}
	//队列中有元素的时候
	return this.right
}
