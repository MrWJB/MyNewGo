package main
import (
	"fmt"
)

/*
	切片介绍：
		英文名slice
		使用和数组类似
		可以动态变化的数组
		切片语法：
			var slices []int
*/
//方式一
func sliceTest(){
	//切片的基本使用
	var intArr [5]int = [...]int{1,22,33,44,99}
	//声明定义切片
	slice := intArr[1:3]
	fmt.Println("intArr=",intArr)
	fmt.Println("slice 的元素是：=",slice) // 22,33
	fmt.Println("slice 的元素个数=",len(slice)) //2
	fmt.Println("slice 的容量 =",cap(slice)) //切片的容量是可以动态变化的

}

/*
	切片的使用
		基本语法：var 切片名 []type = make([]type,len,[cap])
		参数说明：
			type: 就是数据类型
			len：大小
			cap: 指定切片容量，可选，如果你分配了cap,则要求 cap >= len.
*/
func sliceTest01(){
	var slice []float64 = make([]float64,5,10)
	slice[1] = 10
	slice[2] = 20
	//对于切片，必须make使用
	fmt.Println(slice)
	fmt.Println("slice 的size=",len(slice))
	fmt.Println("slice 的cap=",cap(slice))
}

/*
	定义一个切片，直接就指定具体数组，使用原理类似make的方式
*/
func directArray(){
	var strSlice []string = []string{"tom","jack","mary"}
	fmt.Println("strSlice=",strSlice)
	fmt.Println("strSlice size=",len(strSlice))
	fmt.Println("strSlice cap=",cap(strSlice))
}


/*
	切片的遍历和数组一样，也有两种方式
*/
//for 循环常规方式遍历
func arrayTraverse(){
	var arr [5]int = [...]int{10,20,30,40,50}
	slice:=arr[1:4] //20,30,40
	for i:=0;i<len(slice);i++{
		fmt.Printf("slice[%v]=%v\n",i,slice[i])
	}
	fmt.Println()
	for i,v := range slice{
		fmt.Printf("i=%v v=%v",i,v)
	}
}

/*
	用append内置函数，可以对切片进行动态追加
*/
func appendTest(){
	var slice3 []int = []int{100,200,300}
	//通过append直接给slice3追加具体的元素
	slice3=append(slice3,400,500,600)
	fmt.Println("slice3",slice3)

	//通过append将切片slice3追加给slice3
	slice3=append(slice3,slice3...)
	fmt.Println("slice3",slice3)
}


/*
	切片的拷贝操作
	copy(param1,param2) 参数是切片类型
*/
func sliceCopy(){
	//切片使用copy内置函数完成拷贝
	var slice []int=[]int {1,2,3,4,5}
	var slice2 = make([]int,10)

	copy(slice2,slice)
	fmt.Println("slice=",slice)
	fmt.Println("slice2=",slice2)
}

/*
	string 和 slice 
		1.string 底层是一个byte 数组，因此string也可以进行切片处理
*/
func strSlice(){
	str:="hello@atguigu"
	slice:=str[6:]
	fmt.Println("slice",slice)
}

/*
	如果需要修改字符串，可以先将string -> []byte / 或者 []rune -> 修改 -> 重写转成string
	// "hello@atguigu" => 改成 "zello@atguigu"

	
	细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
	原因是 []byte 字节来处理，而一个汉字，是3个字节，因此就会出现乱码
	解决办法：将 string 转成 []rune 即可，因为[]rune 是按字符处理的，兼容汉字
	
*/
func strUpdate(){
	//  var str string = "hello@atguigu"
	//  arr := []byte(str)
	//  arr[0] = 'z'
	//  str=string(arr)
	//  fmt.Println("str=",str)
	 
	 //使用  []rune
	 var str2 string = "hello@atguigu"
	 arr2 := []rune(str2)
	 arr2[0] = '北'
	 str2 = string(arr2)
	 fmt.Println("str2=",str2)
}







//for-range 结构遍历切片
func main(){
	// sliceTest()
	// sliceTest01()
	// arrayTraverse()
	// appendTest()
	// sliceCopy()
	// strSlice()
	strUpdate()
}