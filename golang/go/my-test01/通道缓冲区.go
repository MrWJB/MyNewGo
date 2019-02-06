package main
import "fmt"


/*
	通道缓冲区
		通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
*/
func cacheTest(){
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	ch := make(chan int,2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	//获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


/*
	遍历通道与关闭通道
*/
func fibonacci(n int,c chan int){
	x,y := 0,1
	for i := 0;i<n;i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}

func fibonacciTest(){
	c := make(chan int,10)
	go fibonacci(cap(c),c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了.
	for i := range c{
		fmt.Println(i)
	}
}



func main(){
	// cacheTest()
	fibonacciTest()
}
