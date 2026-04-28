package main

func main() {
	// 声明一个切片
	var s []int
	println(s)    // 输出: []
	print(cap(s)) // 输出: 0

	// 使用 make 函数创建一个切片，指定长度和容量
	s = make([]int, 0, 5)
	println(s)      // 输出: []
	println(cap(s)) // 输出: 5
	println(len(s)) // 输出: 0

	// 使用 make 函数创建一个切片
	s = make([]int, 5)
	println(s) // 输出: [0 0 0 0 0]

	// 使用字面量创建一个切片
	s = []int{1, 2, 3, 4, 5}
	println(s) // 输出: [1 2 3 4 5]

	// 切片的长度和容量
	println(len(s)) // 输出: 5
	println(cap(s)) // 输出: 5

	// 切片的切割
	subSlice := s[1:4]
	println(subSlice) // 输出: [2 3 4]
	for i := range subSlice {
		print(subSlice[i])
		print(" ") // 输出: 2 3 4
	}
	println()

	// 切片的追加
	s = append(s, 6)
	println(s) // 输出: [1 2 3 4 5 6]

	// 切片的复制
	copySlice := make([]int, len(s))
	copy(copySlice, s)
	println(copySlice) // 输出: [1 2 3 4 5 6]

	copySlice = append(copySlice, 7)
	println(copySlice) // 输出: [1 2 3 4 5 6 7]

	combineSlice := append(s, copySlice...)
	println(combineSlice) // 输出: [1 2 3 4 5 6 1 2 3 4 5 6 7]
	combineSlice = append(combineSlice, copySlice...)
	println(combineSlice) // 输出: [1 2 3 4 5 6 1 2 3 4 5 6 7 1 2 3 4 5 6 7]

	combineSlice = append(combineSlice, copySlice...)
	println(combineSlice) // 输出: [1 2 3 4 5 6 1 2 3 4 5 6 7 1 2 3 4 5 6 7 1 2 3 4 5 6 7]
}
