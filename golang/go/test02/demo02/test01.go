package main
import 
(
	"fmt"
)
var sum int
var sub int

//计算两个数的值
func getSumAndSub(n1 int,n2 int)(int,int){
	sum = n1 + n2
	sub = n1 - n2
	return sum,sub
}

//使用递归的方式，求出斐波那契数 1,1,2,3,5,8,13...
func fbn(n1 int) int {
	if (n1==1||n1==2){
		return 1
	}else{
		return fbn(n1-1)+fbn(n1-2) 
	}
}

//值传递的例子
func test02(num1 int){
	num1= num1 + 10
	fmt.Println("获取test02结果的值",num1)
}

//引用传递的例子
func test03(num2 *int){
	*num2 = *num2 + 10
	fmt.Println("获取的test03的结果的值",*num2)
}

//定义一个函数
func getSum(var1 int ,var2 int) int{
	sumval:=var1 + var2
	return sumval
}


type myfun func(int,int) int

//函数作为参数进行传递，例子
//示例1：
func test04(funvar func(int,int) int,num3 int ,num4 int) int{
	return funvar(num3,num4)
}
//示例2：
func test05(funvar myfun,num3 int ,num4 int) int{
	return funvar(num3,num4)
}
//支持对函数返回值命名
func getSumAndDiv(n1 int,n2 int)(sum int,div int){
	sum = n1 + n2
	div = n1 / n2
	return
}


//案例演示： 编写一个函数sum，可以求出 1 到多个int的和
//可以参数的使用
func sums(n1 int,args... int)int {
	sum:=n1
	//遍历args
	for i:=0;i<len(args);i++{
		sum +=args[i] //args[0] 表示取出args切片的一个元素值，其它依次类推
	}
	return sum
}

// func main(){
// 	var num1 int = 22
// 	var num2 int = 12

// 	_,res2:=getSumAndSub(num1,num2)
// 	fmt.Printf("结果2的信息 %d ",res2)

// 	//调用斐波那契
// 	res:=fbn(11)
// 	fmt.Println("获取的结果信息是：",res)
// }


func main(){
	// var num1 int = 22
	// test02(num1)
	// fmt.Println("main 获取的结果的值信息",num1)

	// var num2 int = 22
	// test03(&num2)
	// fmt.Println("main 方法获取的结果值",num2)

	//函数作为参数进行传递
	// var num3 int = 22
	// var num4 int = 32
	// res:=test04(getSum,num3,num4)
	// fmt.Println("获取的结果值是",res)

	// var n1 int = 22
	// var n2 int = 33
	// res,div:=getSumAndDiv(n1,n2)
	// fmt.Printf("获取的sum结果的值 %d ,获取div结果的值 %d",res,div)

	res:=sums(10,0,20,30,12,32,100)
	fmt.Println("获取结果的值",res)
}

