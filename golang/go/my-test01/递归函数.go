package main
import "fmt"



/*
	递归，就是在运行的过程中调用自己。
		Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

*/
func recursionTest(){
	var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

/*
	阶乘
*/
func Factorial(n uint64)(result uint64){
	if (n > 0){
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main(){
	recursionTest()
}