package main
import (
	"fmt";
	"unsafe"
)

func main(){
	var price float32 = 1000.23
	fmt.Printf("price 类型 %T price 字节大小 %d",price,unsafe.Sizeof(price))
	fmt.Println(" price=",price)

	var variable int = 112
	fmt.Printf("variable=%c\n",variable)
}