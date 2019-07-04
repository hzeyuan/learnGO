package main

var a = "G"

func main() {
	// 全局变量
	n()
	m()
	n()
}
func n() {
	print(a)
}
func m() {
	a = "O"
	print(a)
}
