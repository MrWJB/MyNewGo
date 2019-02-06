package main
import (
	"fmt"
)

/*
	方法的声明和调用
*/
type Person2 struct{
	Name string
}

//给person类型绑定一个方法
/*
	test 方法和person类型绑定
	test方法只能通过person 类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
*/
func (p Person2) test(){
	fmt.Println("test() name=",p.Name)
}

/*
	方法快速入门
*/
func (p Person) jisuan() {
	res := 0
	for i:= 0;i <= 1000;i++ {
		res += i
	}
	fmt.Println(p.Name ,"计算的结果是=", res)
}

/*
	给Person 结构体添加getSum 方法，可以计算两个数的和，并返回结果
*/
func (p Person) getSum(n1 int,n2 int) int {
	return n1 + n2
}










func main(){
	// fmt.Println()
	var p Person2
	p.Name = "tom"
	p.test() //调用方法
}