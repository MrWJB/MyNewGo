package main
import (

	"fmt"
)

func main(){
	var num1 int=99
	// var num2 float64=23.222
	// var bol bool=true
	// var char byte='h'
	var str string //空字符串

	str=fmt.Sprintf("%d",num1)
	fmt.Printf("str type %T str=%q\n",str,str)

	var i int = 10
	fmt.Println("i的地址=",&i)

	var prt *int =&i
	fmt.Println(*prt)

}

