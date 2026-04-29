package stack

import "container/list"

/* 基于链表实现的栈 */
type linkedListStack struct {
	// 使用内置包 list 来实现栈
	data *list.List
}

/* 初始化栈 */
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

/* 入栈 */
func (s *linkedListStack) push(value int) {
	s.data.PushBack(value)
}

/* 出栈 */
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

/* 访问栈顶元素 */
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

/* 获取栈的长度 */
func (s *linkedListStack) size() int {
	return s.data.Len()
}

/* 判断栈是否为空 */
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

/* 获取 List 用于打印 */
func (s *linkedListStack) toList() *list.List {
	return s.data
}

func LinkedListStackTest() {
	stack := newLinkedListStack()

	/* 元素入栈 */
	stack.push(1)
	stack.push(3)
	stack.push(2)
	stack.push(5)
	stack.push(4)

	/* 访问栈顶元素 */
	peek := stack.peek()
	/* 输出栈顶元素 */
	println(peek.(int)) // 输出: 4

	/* 元素出栈 */
	pop := stack.pop()
	/* 输出出栈元素 */
	println(pop.(int)) // 输出: 4

	/* 获取栈的长度 */
	size := stack.size()
	/* 输出栈的长度 */
	println(size) // 输出: 4

	/* 判断是否为空 */
	isEmpty := stack.isEmpty()
	/* 输出判断结果 */
	println(isEmpty) // 输出: false
}
