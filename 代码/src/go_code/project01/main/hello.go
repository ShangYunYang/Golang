// helo.go,输出 “hello word”
package main

//表示hello.go 文件所在的包时main，在go中每个文件都必须归属于一个包

import "fmt"

//表示引入一个包，包名fmt 引入该包后，可以使用fmt函数 比如fmt.Println

func main() {
	//func 是一个关键字 表示一个函数，main时一个函数名，是一个主函数，即程序入口
	fmt.Println("hello word")
	//表示调用fmt包的函数，Println输出hello word
}
