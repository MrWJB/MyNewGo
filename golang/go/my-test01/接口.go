package main
import "fmt"


/*
	定义接口
*/
type Phone  interface{
	call()
}


type NokiaPhone struct {
}

/*
	实现call方法
*/
func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia,I can call you !")
}

type IPhone struct{
}

func (iphone IPhone) call(){
	fmt.Println("I am iPhone, I can call you!")
}


/*
	接口
*/
func interfaceTest(){

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
	
}

func main(){
	interfaceTest()
}