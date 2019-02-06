package main
import (
	"fmt"
	"strconv"
	"time"
	"runtime"
)



/*
	编写一个函数，每隔1秒输出 "Hello world"
*/
func test(){
	for i := 0 ;i<10;i++ {
		fmt.Println("test() hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}



/*
	快速入门小结
		1.主线程是一个物理线程，直接作用在cpu上的。是重量级的，非常耗费cpu资源。
		2.协程从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小。
		3.Golang 的协程机制是重要的特点，可以轻松的开启上万个协程。其它编程语言的并发机制是一般基于线程的，开启过多的线程，资源耗费大，这里就突显Golang在并发上的优势了

		goroutine 的调度模型
			MPG模式基本介绍
*/
func goroutineTest(){
	//获取当前系统cpu的数量
	num := runtime.NumCPU()
	//我这里设置num-1 的cpu运行go程序
	runtime.GOMAXPROCS(num-1)
	fmt.Println("num=",num)
}


//map 应该做出一个全局的
var (
	myMap = make(map[int]int,10)
)

/*
	channel(管道) 
		需求： 现在计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中。最后显示出来。要求使用goroutine完成
*/
//test01 这个函数就是计算 n!,让将这个结果放入到myMap
func test01(n int){
	res := 1
	for i := 1; i<=n; i++{
		res*=i
	}
	//这里我们将res放入到myMap
	myMap[n] = res 
} 


var (
	myMap = make(map[int]int,10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包： sychornized 同步
	//Mutex: 是互斥
	lock sync.Mutex
)

//test 函数就是计算n!，让将这个结果放入到myMap
func test02(){
	res := 1
	for i := 1; i<=n; i++{
		res*=i
	}
	//这里我们将res 放入到myMap
	//枷锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}







/*
	使用全局变量加锁同步改进程序
*/
func 








func main(){
	// go test()

	// for i := 0 ;i<10;i++ {
	// 	fmt.Println("main() hello,golang " + strconv.Itoa(i))
	// 	time.Sleep(time.Second)
	// }
	// goroutineTest()

	//我们这里开启多个协程完成这个任务【200个】
	for i:=0;i<200;i++ {
		go test01(i)
	}
	// // fmt.Println(myMap)
	// //休眠10秒钟
	time.Sleep(time.Second * 10)
	// //这里我们输出结果，变量这个结果
	// for i,v:= range myMap {
	// 	fmt.Println("map[%d]=%d\n",i,v)
	// }


	//这里我们输出结果，变量这个结果
	lock.Lock()
	for i,v := range myMap {
		fmt.Println("map[%v]=%v",i,v)
	}
	lock.Unlock()

}