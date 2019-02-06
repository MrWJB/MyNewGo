package main
import (
	"fmt"
)




/*
	channel的基本介绍
		1.channel 本质就是一个数据结构-队列
		2.数据是先进先出
		3.线程安全，多goroutine 访问时，不需要加锁，就是说channel本身就是线程安全的
		4.channel有类型的，一个string的channel只能存放string类型的数据
	定义/声明channel
		var 变量名 chan 数据类型
		举例：
			var intChan chan int //intChan 用于存放int数据
		说明：
			channel是引用类型
			channel必须初始化才能写入数据，及make后才能使用
			管道是有类型的，intChan 只能写入 整数 int
	channel使用的注意事项
		1.channel中只能存放指定类型的数据类型
		2.channel的数据放满之后，就不能再放入了
		3.如果从channel取出数据后，可以继续放入
		4.在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock
*/
func ChannelTest01(){
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)
	//2.intChan是什么
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n",intChan,&intChan)
	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//intChan <- 98 //注意点，当我们给管道写入数据时，不能超过其容量

	//4.看看管道的长度 和 cap(容量)
	fmt.Printf("channel len=%v cap=%v \n",len(intChan),cap(intChan))

	//5.从管道中读取数据
	var num2 int
	num2 =<-intChan
	fmt.Printf("num2=",num2)
	fmt.Printf("channel len=%v cap=%v\n",len(intChan),cap(intChan))

	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	num5 := <-intChan

	fmt.Printf("num3=",num3,"num4=",num4,"num5=",num5)

}


func ChannelTest02(){
	var intChan chan int
	intChan = make(chan int,3)

	//赋值
	intChan <- 10
	intChan <- 20
	intChan <- 30
	//因为intChan 的容量为3，再存放会报告deadlock
	
	//取数据
	num1 := <- intChan
	num2 := <- intChan
	num3 := <- intChan

	//因为intChan 这时已经没有数据了，再取就会报告deadlock
	fmt.Printf("num1=%v num2=%v num3=%v ",num1,num2,num3)

}

/*
	创建一个mapChan,最多可以存放10个map[string]string的key-val ，演示写入和读取
*/
func ChannelTest03(){
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)
	m1 := make(map[string]string,20)
	m1["city1"] = "北京"
	m1["city2"] = "天津"

	m2 := make(map[string]string,20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	//写入
	mapChan <- m1
	mapChan <- m2

	//读取
	maps1 := <- mapChan
	maps2 := <- mapChan

	for i1,v1 := range maps1 {
		fmt.Printf("maps1[%v]=%v",i1,v1)
	}
	for i2,v2 := range maps2 {
		fmt.Printf("maps2[%v]=%v",i2,v2)
	}
}

type Cat struct {
	Name string
	Age int
}


/*
	创建一个catChan,最多可以存放10个Cat结构体变量，演示写入和读取的用法
*/
func ChannelTest04(){
	var catChan chan Cat
	catChan = make(chan Cat,10)

	cat1 := Cat{Name:"tom~",Age:20,}
	cat2 := Cat{Name:"tom",Age:18,}

	catChan <- cat1
	catChan <- cat2

	//取出
	cat11 := <- catChan
	cat22 := <- catChan

	fmt.Println(cat11,cat22)

}


/*
	创建一个catChan2,最多可以存放10个*Cat 变量，演示写入和读取的用法
*/
func ChannelTest05(){
	var catChan chan *Cat
	catChan = make(chan *Cat,10)

	cat1 := Cat{Name:"tom~",Age:20,}
	cat2 := Cat{Name:"tom",Age:18,}

	catChan <- &cat1
	catChan <- &cat2

	//取出
	cat11 := <- catChan
	cat22 := <- catChan

	fmt.Println(*cat11,*cat22)
}


/*
	创建一个allChan,最多可以存放10个任意数据类型变量，演示写入和读取的用法
*/
func ChannelTest06(){
	var allChan chan interface{}
	allChan = make(chan interface{},10)

	cat1 := Cat{Name:"tom~",Age:20,}
	cat2 := Cat{Name:"tom",Age:18,}

	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"

	//取出
	cat11 := <- allChan
	cat22 := <- allChan
	v1 := <- allChan
	v2 := <- allChan

	fmt.Println(cat11,cat22,v1,v2)
}

/*
	channel的遍历和关闭
		channel的关闭
			使用内置函数close 可以关闭channel,当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据
*/
func ChannelTest07(){
	intChan := make(chan int,3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	//这里不能够再写入数道channel中
	//intChan <- 300
	fmt.Println("okook~")
	//当管道关闭后，读取数据时可以的
	n1 := <- intChan
	fmt.Printf("n1=%v ",n1)
}

/*
	channel的遍历
		channel 支持 for-range 的方式进行遍历，请注意两个细节
			1.在遍历时，如果channel没有关闭，则会出现deadlock的错误
			2.在遍历时如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
*/
func ChannelTest08(){
	intChan := make(chan int,100)
	for i:=0;i<100;i++{
		intChan <- i *2  //放入一百个数据道管道里面
	}
	//遍历管道不能使用普通的for循环

	//在遍历的时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历是，如果channel已经关闭则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan)
	for v := range intChan{
		fmt.Println("v=%v",v)
	}
}

/*
	channel 使用细节和注意事项
		1.channel 可以声明为只读，或者只写性质
		2.
*/
func ChannelTest09(){
	//管道可以声明为只读或者只写
	//1.在默认的情况下，管道是双向 var chan1 chan int  	//可读可写

	//2.声明为只写
	var chan2 chan<- int

	chan2 = make(chan int,3)
	chan2 <- 20
	fmt.Printf("chan2=",chan2)

	//声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	fmt.Printf("num2",num2)
}








func main(){
	// ChannelTest01()
	// ChannelTest02()
	// ChannelTest03()
	// ChannelTest04()
	// ChannelTest05()
	// ChannelTest06()
	// ChannelTest07()
	ChannelTest08()
}





