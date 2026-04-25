package array

func InitArray() {
	// 定义一个长度为5的整数数组
	var arr [5]int

	// 初始化数组元素
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	// 定义并初始化一个字符串数组
	strArr := [3]string{"Hello", "World", "Go"}

	// 输出数组元素
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	for i := 0; i < len(strArr); i++ {
		println(strArr[i])
	}
}
