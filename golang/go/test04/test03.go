package main 
import "fmt"




/*
	接口
		1.接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态 和 高内聚低耦合 的思想
		2.golang 中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量j就实现这个接口。因此，golang中没有implement这样的关键字

		注意事项： 
			1.接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
			2.接口中所有的方法都没有方法体，即都是没有实现的方法
			3.在 golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口。
			4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
			5.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
			6.一个自定义类型可以实现多个接口
			7.golang接口中不能有任何变量
			8.一个接口（比如A接口）可以继承多个别的接口（比如B,C接口),这时如果要实现A接口，也必须将B,C接口的方法也全部实现。
			9.interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，那么会输出nil
			10.空接口interface{} 没有任何方法，所以所有类型都实现了该接口，即我们可以把任何一个变量赋给空接口。
*/
type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say(){
	fmt.Println("Stu Say()")
}

func SayTest(){
	var stu Stu  //结构体变量实现了Say()方法
	var a AInterface = stu
	a.Say()

}


/*
	类型断言
		由一个具体的需要，引出了类型断言
			类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言，具体的如下：
*/
func interfaceTest(){
	var x interface{}
	var b2 float32 = 1.1
	x = b2   //空接口，可以接收任意类型
	y := x.(float32)
	fmt.Printf("y 类型是 %T 值是=%v",y,y)

	//类型断言（带检测的）
	if y,ok := x.(float32);ok {
		fmt.Println("convert success")
		fmt.Printf("y 类型是 %T 值是=%v",y,y)
	}else{
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")
}


/*
	类型断言的最佳实践2
		写一个函数，循环判断传入参数的类型
*/

func TypeJudge(items... interface{}){
	for index,x := range items {
		switch x.(type) {
			case bool:
				fmt.Printf("第%v个参数是bool类型，值是%v\n",index,x)
			case float32:
				fmt.Printf("第%v个参数是float32类型，值是%v\n",index,x)
			case float64:
				fmt.Printf("第%v个参数是float64类型，值是%v\n",index,x)
			case int,int32,int64:
				fmt.Printf("第%v个参数是 整数 类型，值是%v\n",index,x)
			case string:
				fmt.Printf("第%v个参数是 string 类型，值是%v\n",index,x)
			default:
				fmt.Printf("第%v个参数是 类型 不确定，值是%v\n",index,x)
		}
	}
}
/*
	测试
*/
func TypeJudgeTest(){
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	TypeJudge(n1,n2,n3,name,address,n4)
}




func main(){
	// SayTest()
	TypeJudgeTest()
}












