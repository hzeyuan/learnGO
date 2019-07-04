package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + a // 编译错误 int32 不能int类型相加
	b = b + 5 // 因为 5 是常量，所以可以通过编译
}
