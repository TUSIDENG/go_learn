package stack

func StackLearn() {
	/* 初始化栈 */
	// 在 Go 中，推荐将 Slice 当作栈来使用
	var stack []int

	/* 元素入栈 */
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)

	/* 访问栈顶元素 */
	peek := stack[len(stack)-1]
	/* 输出栈顶元素 */
	println(peek) // 输出: 4

	/* 元素出栈 */
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	/* 输出出栈元素 */
	println(pop) // 输出: 4

	/* 获取栈的长度 */
	size := len(stack)
	/* 输出栈的长度 */
	println(size) // 输出: 4

	/* 判断是否为空 */
	isEmpty := len(stack) == 0
	/* 输出判断结果 */
	println(isEmpty) // 输出: false
}
