package main
import "fmt"




/*
	类型转换
		类型转换用于将一种数据类型的变量转换为另外一种类型的变量。
		type_name(expression)
		type_name 为类型，expression 为表达式。
*/
func typeTest(){
	var sum int = 13
	var count int = 5 
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为：%f\n",mean)
}


func main(){
	typeTest()
}