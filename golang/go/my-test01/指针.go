package main
import "fmt"

/*
	指针的使用
		变量是一种使用方便的占位符，用于引用计算机内存地址。
		Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	什么是指针
		var var_name *var-type
		var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
*/
func ponitTest(){
	// var a int = 10
	// fmt.Printf("变量的地址：%x\n",&a)
	//一个指针变量指向了一个值的内存地址。
	/* 声明实际变量 */
	var a int = 20
	/* 声明指针变量 */
	var ip *int
	/* 指针变量的存储地址 */
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a  )
	/* 指针变量的存储地址 */	
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	/*
		空指针
			当一个指针被定义后没有分配到任何变量时，它的值为 nil。
			nil 指针也称为空指针。
			nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
			一个指针变量通常缩写为 ptr。
	*/
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )
	
	


}


func main(){
	ponitTest()
}