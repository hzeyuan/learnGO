package main

var a string

func main() {
	//	函数嵌套函数
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}
func f2() {
	print(a)
}
