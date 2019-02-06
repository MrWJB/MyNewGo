package main
import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

type Monster2 struct {
	Name string
	Age int
}


/*
	使用反射机制，编写函数的适配器，桥连接
	反射常见应用场景有以下两种
		1.不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这种需要对函数或方法反射。
*/
func reflectTest01(b interface{}){
	//通过反射获取的传入的变量的type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType",rTyp)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=",n2)
	//下面我们将rVal转成interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=",num2)
}

//专门演示反射【对j结构体的反射】
func reflectTest02(b interface{}){
	//通过反射获取的传入的变量的type,kind,值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=",rTyp)
	//2.获取到 reflect.Vlaue
	rVal := reflect.ValueOf(b)
	//下面我们将rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T\n",iV,iV)
	//将interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言
	stu,ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n",stu.Name)
	}
}

/*
	反射课堂练习

*/
func test01(){
	var str string = "tom"
	fs := reflect.ValueOf(&str)
	fmt.Println("fs=",fs)
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n",str)
}




func main(){
	test01()
}

/*
	Type 和 Kind 的区别
		 Type 是类型，Kind是类别，Type 和 Kind 可能是相同的，也可能是不同的
			 比如： var num int = 10 	num 的Type 是int,Kind 也是 int
			 比如： var stu Student 	stu 的Type 是 main.Student ,Kind 是 struct
*/


// func main(){
// 	//演示对（基本数据类型、interface{}、reflect.Value）进行反射的基本操作
// 	//1.先定义一个int
// 	// var num int= 100
// 	// reflectTest01(num)

// 	//2.定义一个Student 的实例
// 	stu := Student{
// 		Name : "tom",
// 		Age : 20,
// 	}
// 	reflectTest02(stu)
// }




/*
	使用select可以解决从管道取数据的阻塞问题
*/
// func main(){
// 	//使用select可以解决从管道取数据的阻塞问题
// 	//1.定义一个管道10个数据int
// 	var intChan chan int
// 	intChan = make(chan int,10)
// 	for i:=0;i<10;i++{
// 		intChan <- i
// 	}
// 	//2.定义一个管道5个数据string
// 	stringChan := make(chan string,5)
// 	for i:=0;i<5;i++{
// 		stringChan <- "hello" + fmt.Printf("%d",i)
// 	}
// 	//传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
// 	//问题，在实际开发中，可能我们不好确定什么关闭该管道
// 	//可以s使用select方式可以解决

// 	//4.gorountine 中使用recover，解决协程中出现panic，导致程序崩溃问题

// }