package main

import "fmt"

// 多变量声明
// func main() {
// 	// var n1,n2,n3 int
// 	// fmt.Println("n1=",n1,"n2 =",n2,"n3=",n3)
// 	var n4, name, n5 = 100, "tom", 888
// 	fmt.Println("n4=", n4, "name=", name, "n5=", n5)

// }

// 如何一次性声明多个全部变量
// 定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

func main() {
	fmt.Println(100, name, n2)
}
