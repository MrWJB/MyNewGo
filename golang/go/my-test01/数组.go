package main
import (
	"fmt"
)

func myTest01(){
	//初始化数组
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance的值：",balance)
	//访问数组元素
	var salary float32 = balance[4]
	fmt.Println("获取的值是：",salary)
}


func main(){
	//定义一个数组
	var intArray [5]int
	intArray[0] = 12
	intArray[1] = 22
	intArray[2] = 32
	intArray[3] = 41
	fmt.Printf("获取的结果是：%v\n",intArray)
	//循环遍历数组
	for index,val := range intArray{
		fmt.Printf("intArray[%d]=%d\n",index,val)
	}
}