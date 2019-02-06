package main
import (
	"fmt"
)




// func main(){

// 	//方式一
// 	//求两个数的和，匿名函数的方式
// 	res:=func (n1 int,n2 int) int{
// 		return n1 + n2
// 	}(10,10)
	
// 	fmt.Println("两个数的和",res)

// 	//方式二
// 	//把匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
// 	a := func (n1 int,n2 int) int {
// 		return n1 + n2
// 	}
// 	res2 := a(10,20)
// 	fmt.Println("获取结果的值",res2)
// 	res3 := a(30,60)
// 	fmt.Println("获取结果的值",res3)


// }


//累加功能 闭包就是一个函数和与其相关的引用环境组合的一个整体（实体，）
func addUpper() func (int) int{
	var n int = 10
	return func (x int) int{
		n =n + x
		return n
	}
}

func main(){
	res:=addUpper()
	fmt.Println("获取的数据",res(1)) 
}