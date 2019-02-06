package main
import (
	"fmt";
	// "math/rand";
	// "time"
)

//加减乘除算法方法
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

	// str:="hello world! The world is very beautiful."
	// //1.使用 for-range方式
	// for index,val:=range str{
	// 	fmt.Println("index值 %d,val值 %c\n",index,val)
	// }
	// //使用原始方式
	// var str2 string ="hello world! The world is very beautiful."
	// for i:=0;i<len(str2);i++{
	// 	fmt.Println("%c \n",str[i])
	// }

	// 包含汉字的情况[]rune	
	// for {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n:=rand.Intn(100)+1
	// 	fmt.Println("n=",n)
	// 	count++
	// 	if(n==99){
	// 		break;
	// 	}
	// 	fmt.Printf("生成99 一共用了 %d 次\n",count)
	// }

	//--------------------------

	// var n int =20
	// if n > 20{
	// 	goto lable1
	// }
	// fmt.Println("ok1")
	// fmt.Println("ok2")
	// lable1:
	// fmt.Println("ok3")
	// fmt.Println("ok4")

	//模拟加减乘除算法	
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