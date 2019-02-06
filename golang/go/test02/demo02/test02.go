package main
import (
	"fmt"
)

//全局变量定义
var age=test01()

//test01 函数
func test01() int {
	fmt.Println("test01()函数执行")
	return 90
}

//初始化函数，会在main函数调用之前调用
func init(){
	fmt.Println("init函数执行")
}


func main(){
	fmt.Println("main函数执行")
}



