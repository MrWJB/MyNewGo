package main
import "fmt"


/*
	定义切片
*/
func sliceTest01(){
	//你可以声明一个未指定大小的数组来定义切片：
	//切片不需要说明长度。
	var slice []int
	fmt.Println("slice=",slice)

	//或使用make()函数来创建切片:
	slice2 := make([]string,10)
	slice2[0] = "zhangsan"
	fmt.Println("slice2=",slice2)

	//切片初始化
	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s := []int{1,2,3}
	fmt.Printf("s的值：%v\n",s)

	//定义一个数组arr
	var arr = [5]int{1,2,3,4,5}

	//初始化切片s1,是数组arr的引用
	s1 := arr[:]
	fmt.Printf("s1的值：%v\n",s1)

	//s := arr[startIndex:endIndex]  
	// 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	s2 := arr[0:3]
	fmt.Printf("s2的值：%v\n",s2)

	// 缺省endIndex时将表示一直到arr的最后一个元素
	// s := arr[startIndex:] 
	s3 := arr[2:]
	fmt.Printf("s3的值：%v\n",s3)

	// 缺省startIndex时将表示从arr的第一个元素开始
	// s := arr[:endIndex] 
	s4 := arr[:3]
	fmt.Printf("s4的值：%v\n",s4)

	// s1 := s[startIndex:endIndex] 
	// 通过切片s初始化切片s1
	s5 := s4[0:3]
	fmt.Printf("s5的值：%v\n",s5)

}

/*
	len() 和 cap() 函数
		切片是可索引的，并且可以由 len() 方法获取长度。
		切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/
func lenAndCapTest02(){
	var numbers = make([]int,5,10)
	fmt.Printf("获取的len长度 %d,cap容量为：%d",len(numbers),cap(numbers))
}


/*
	切片截取
		可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]
*/
func sliceTruncTest(){
	// 创建切片
	numbers := []int{1,2,3,4,5,6,7,8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] == ",numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	numbers2 := numbers1[:2]
	printSlice(numbers2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	numbers3 := numbers1[2:5]
	printSlice(numbers3)



}

/*
	append() 和 copy() 函数
		如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
*/
func appendOrCopyTest(){
	var numbers []int
	printSlice(numbers)
	
	/* 允许追加空切片 */
	numbers = append(numbers,0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers,1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers,2,3,4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int,len(numbers),(cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1,numbers)
	printSlice(numbers1)

	

}




func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}













func main(){
	sliceTest01()
	lenAndCapTest02()
}