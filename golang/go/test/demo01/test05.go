package main
import (
	"fmt"
)
// func multi(num1 float64,num2 float64)float64{
// 	sum:=num1+num2
// 	return sum
// }

// func main(){
// 	var num1 float64 = 32.23
// 	var num2 float64 = 43.42
// 	sum:=multi(num1,num2)
// 	fmt.Println("计算后的结果值！sum=",sum)
// 	fmt.Println("获取的结果信息》",utils.Cal())
// }

func main(){
	var n int =32
	if n>20{
		goto lable1
	}
	fmt.Println("ok1")
	fmt.Println("ok2")
	lable1:
	fmt.Println("ok3")
	fmt.Println("ok4")
}
