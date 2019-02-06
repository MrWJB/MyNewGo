package main
import (
	"fmt"
	"time"
	"strconv"
)

func testTime(){
	//获取当前时间
	now:=time.Now()
	fmt.Println("获取的时间：",now)

	//通过now 可以获取到年月日，时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",now.Month())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("小时=%v\n",now.Hour())
	fmt.Printf("分钟=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	//格式化时间
	//方式一 使用Printf 或者 SPrintf
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

	dataStr:=fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dateStr=%v\n",dataStr)

	//方法二 使用time.Format() 方法完成
	fmt.Printf(now.Format("2008-01-12 15:23:03"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("16:23:01"))
	fmt.Println()
	
}

func testTime01(){
	i:=0
	for{
		i++
		fmt.Println(i)
		//休眠
		time.Sleep(time.Millisecond * 500)
		if i == 100{
			break
		}
	}
}


/*
	time 的Unix 和 UnixNano的方法
*/
func  unixOrUnixNano(){
	// fmt.Printf("unix时间戳=%v unixnano 时间戳=%v \n",now.Unix(),now.UnixNano())
}


/*
	时间和日期的课堂练习
*/
func test03(){
	str:=""
	for i:=0;i<100000;i++{
		str += "hello " + strconv.Itoa(i)
		// fmt.Println("获取的值：",str)
	}
}



func main(){
	// testTime()
	// testTime01()
	// unixOrUnixNano()

	/*
		测试:在执行test03前 ，先获取到当前的您unix时间戳
	*/
	start:=time.Now().Unix()
	test03();
	end:=time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n",end - start)
}