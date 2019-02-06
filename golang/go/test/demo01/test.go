package main
import "fmt"
import "unsafe"

var age = 22
var sex = "女"
func main(){
	var i int64 = 100;
	// name:="zhangsan"
	fmt.Printf("i 的 类型 %T i占用的字节数是%d \n",i,unsafe.Sizeof(i))
	// fmt.Println(name+"\t"+age+"\t"+sex) 
	fmt.Println("hello world!\tmy china!")
	fmt.Println("This is my second case!\nDo you understand!")
}

