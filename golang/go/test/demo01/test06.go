package main
import (
	"fmt"
)

//计算
func cal(n1 float64,n2 float64,operator byte) float64{
	var res float64
	switch operator{
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符号错误！")
		}
	return res
}

func main(){
	var num1 float64
	var num2 float64
	var operator byte
	fmt.Println("请输入num1的值")
	fmt.Scanln(&num1)
	fmt.Println("请输入num2的值")
	fmt.Scanln(&num2)
	fmt.Println("请输入operator的值")
	fmt.Scanln(&operator) 
	var res=cal(num1,num2,operator)
	fmt.Println("结果值是：",res) 
}
