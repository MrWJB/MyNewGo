package main
import "fmt"


type Book struct{
	title string
	author string
	subject string
	book_id int
}

/*
	结构体
		结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
*/
func structTest(){
	// 创建一个新的结构体
	fmt.Println(Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Book{title:"Go 语言",author:"www.runoob.com",subject:"Go 语言教程",book_id:64321})
	// 忽略的字段为 0 或 空
	fmt.Println(Book{title:"Go 语言",author:"www.runoob.com"})
}


func main(){
	structTest()
}