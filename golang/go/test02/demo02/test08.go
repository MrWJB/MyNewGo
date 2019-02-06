package main
import (
	"fmt"
	"time"
	// "strconv"
	"rand"
	"errors"
)

/*
	内置函数
		1.举例说明 new的使用
*/
func builtFunc(){
	num1:=100
	fmt.Printf(" num1 的类型 %T ，num1 的值=%v,num1的地址%v\n ",num1,num1,&num1)

	num2:=new(int) //*int
	*num2 = 100
	fmt.Printf(" num2的类型 %T，num2的值%v,num2的地址%v：num2这个指针，指向的值=%v ",num2,num2,&num2,*num2)
}

/*
	错误处理机制  go 中引入的处理方式为：defer panic recover
	描述：
		1.Go 中可以抛出一个panic的异常，然后在 defer中通过recover 捕获这个异常，然后正常处理
	使用 defer + recover 来处理错误
*/
func deferAndRecover(){
	defer func(){
		err := recover() //recover() 内置函数，可以捕获到异常
		if err !=nil{ //不为空，说明捕获到异常
			fmt.Println("err=",err)
			//这里可以将错误信息发送给管理员。。。
			fmt.Println("发送邮件给admin@sohu.com ~")
		}
	}()
	num1:=10
	num2:=0
	res:=num1/num2
	fmt.Println("res=",res)
}

/*
	自定义错误
	函数去读取以配置文件init.conf的信息
	如果文件名传入不正确，我们就返回一个自定义的错误
*/
func readConf(name string)(err error){
	if name == "config.ini"{
		//读取。。。
		return nil
	}else{
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func testDemo(){
	err:=readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("testDemo()继续执行....")
}


/*
	使用数组的方式来解决问题
*/
func arrayTest(){
	//1.定义一个数组
	var hens [7]float64
	//2.给数组的每个元素赋值，元素的下标是从0开始的 0-5
	hens[0] = 3.0 //hens数组的第一个元素 hens[0]
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 150.0

	//3.遍历数组求出总体重
	totalWeight2:=0.0
	for i:=0;i<len(hens);i++{
		totalWeight2 +=hens[i]
	}

	//4.求出平均体重
	avgWeight2:=fmt.Sprintf("%.2f",totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v avgWeight2=%v ",totalWeight2,avgWeight2)
}


/*
	测试数组
*/
func arrayTest02(){
	var intArr [3]int
	//当我们定义完数组后，其实数组的各个元素都有一个默认的 0 
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30

	fmt.Println(intArr)
	fmt.Printf(" intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p ",&intArr,&intArr[0],&intArr[1],&intArr[2])
}

/*
	数组的使用
	从终端循环输入5个成绩，保存到float64数组，并输出
*/
func arrayUse(){
	var score [5]float64

	for i:=0;i<len(score);i++{
		fmt.Printf("请输入第%d个元素的值\n",i+1)
		fmt.Scanln(&score[i])
	}

	//变量数组打印
	for i:=0;i<len(score);i++{
		fmt.Printf("score[%d]=%v\n",i,score[i])
	}
}

/*
	四种初始化数组的方式
*/
func initMethod(){
	//方式一
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("numArr01=",numArr01)
	
	//方式二
	var numArr02 =[3]int{5,6,7}
	fmt.Println("numArr02=",numArr02)

	//方式三  这里的[...] 是规定的写法
	var numArr03 = [...]int{8,9,10}
	fmt.Println("numArr03=",numArr03)

	//根据索引
	var numArr04 = [...]int{1:800,0:900,2:1000}
	fmt.Println("numArr04=",numArr04)

	//类型推导
	strArr05 := [...]string{1:"tom",0:"jack",2:"mary"}
	fmt.Println("strArr05=",strArr05)
}


/*
	数组遍历
*/
func arrayTraverse(){
	//for-range 遍历数组
	heroes:=[...]string{"宋江","吴用","卢俊义"}
	for index,val:=range heroes{
		fmt.Printf("index=%v val=%v\n",index,val)
		fmt.Printf("heroes[%d]=%v\n",index,heroes[index])
	}

	for _,val:=range heroes{
		fmt.Printf("元素的值=%v\n",val)
	}
}


/*
	定义数组
*/
func defArray(){
	//数组是多个相同类型数据的组合，一个数组一旦声明，定义了，其长度是固定的，不能动态变化。
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	// arr01[2] = 1.2 //因为数组的类型是int ，就不能给1.2
	//其长度是固定的，不能动态变化，否则会报越界异常
	// arr01[3] = 890 //数组不能动态增长
	fmt.Println(arr01)
}


/*
	数组的应用案例
		条件：创建一个byte 类型的26个元素的数组，分别 放置‘A’-‘Z’ 。 使用for 循环访问所有元素并打印出来。 提示：字符数据运算 'A'+1->B
*/
func byteAdd(){
	//1.创建一个byte 类型的26个元素的数组，分别 放置‘A’-‘Z’ 
	//2. 使用for 循环访问所有元素并打印出来。 提示：字符数据运算 'A'+1->B

	var myChars [26]byte
	for i:=0;i<26;i++{
		myChars[i] = 'A' + byte(i) //注意需要将i => byte
	}

	for i:=0;i<26;i++{
		fmt.Printf("%c",myChars[i])
	}
}

/*
	请求出一个数组的最大值，并得到对应的下标
*/
func maxArray(){
	fmt.Println()
	var intArr [6]int = [...]int {1,-1,9,90,11,9000}
	maxVal:=intArr[0]
	maxValIndex:=0

	for i:=1;i<len(intArr);i++{
		//然后从第二个元素开始循环比较，如果发现又更大，则交换
		if maxVal < intArr[i]{
			maxVal= intArr[i]
			maxValIndex=i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v",maxVal,maxValIndex)
}


/*
	请求出一个数组的和和平均值。for-range
*/
func overAndAvg(){
	var intArr [5]int =[...]int {1,-1,9,90,12}
	sum:=0
	for _,val:=range intArr{
		//累加求和
		sum += val
	}

	//如何让平均值保留到小数
	fmt.Printf("sum=%v 平均值=%v",sum,float64(sum)/float64(len(intArr)))
}

/*
	要求随机生成五个数，并将其反转打印，复杂应用
		思路: 1.随机生成五个数，rand.Intn()函数
			 2.当我们得到随机数后，就放到一个数组 int数组
			 3.反转打印，交换的次数是： len /2 
*/
func randomInt(){
	var intArr3 [5]int 
	//为了每次生成的随机数不一样，我们需要给一个seed值
	len:= len(intArr3)

	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len;i++{
		intArr3[i] = rand.Intn(100) 0 <= n <100
	}
	fmt.Println("交换前~=",intArr3)
	//反转打印，交换的次数是  len/2
	//倒数第一个和第一个元素交换，倒数第二个和第二个元素交换
	temp := 0 //做一个临时的变更
	for i :=0; i<len/2; i++{
		temp = intArr3[len - 1 -i]
		intArr3[len -1 - i] = intArr3[i]
		intArr3[i]=temp
	}
	fmt.Println("交换后~=",intArr3)
}






/*
	主函数
*/
func main(){
	// builtFunc()

	/*
		测试 deferAndRecover()
	*/
	// deferAndRecover()
	// for{
	// 	fmt.Println("main()下面的代码...")
	// 	time.Sleep(time.Second)
	// }

	// readConf("conf")
	// testDemo()
	// arrayTest()
	// arrayTest02()
	// arrayUse()
	// initMethod()
	// arrayTraverse()
	// byteAdd()
	// maxArray()
	// overAndAvg()
	randomInt()
}