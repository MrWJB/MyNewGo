// package main
// import (
// 	"fmt"
// 	"bufio"
// 	"os"
// 	"io/ioutil"
// )

// /*
// 	定义一个结构体，用于保存统计结果
// */
// type CharCount struct {
// 	ChCount int //记录英文个数
// 	NumCount int //记录数字的个数
// 	SpaceCount int //记录空格的个数
// 	OtherCount int //记录其它字符的个数
// }

// /*
// 	统计英文、数字、空格和其它字符数量
// 		说明：
// 			统计一个文件中含有的英文、数字、空格及其它字符数量
// */
// func statisticsCount(){
// 	//思路： 打开一个文件，创一个Reader
// 	//每读取一行，就去统计该行有多少个 英文、数字、空格和其它字符
// 	//然后将结果保存到一个结构体
// 	fileName := "e:/abc.txt"
// 	file,err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Printf("open file err=%v\n",err)
// 		return
// 	}
// 	defer file.Close()
// 	//定义个CharCount实例
// 	var count CharCount
// 	//创建一个Reader
// 	reader := bufio.NewReader(file)
// 	//开始循环的读取fileName的内容
// 	for {
// 		str ,err := reader.ReadString('\n')
// 		//读取文件末尾就退出
// 		if err == io.EOF { 
// 			break
// 		}
// 		//为了兼容中文字符，可以将str转成 []rune
// 		str = []rune(str)
// 		//遍历str,进行统计
// 		for _,v := range str {
// 			switch {
// 			case v >= 'a' && v <= 'z':
// 				fallthrough  //穿透
// 			case v >= 'A' && v <= 'Z':
// 				count.ChCount ++
// 			case v == '' || v == '\t':
// 				count.SpaceCount ++
// 			case v >= '0' && v <= '9':
// 				count.NumCount ++
// 			default:
// 				count.OtherCount ++
// 			}
// 		}
// 	}
// 	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
// }


package main 
import (
	"fmt"
	"flag"
)



func main(){
	//定义几个变量用于接收命令行的参数值
	var user string
	var pwd string
	var host string 
	var port int

	/*
		&user 就是接收用户命令行中输入的 -u 后面的参数值
		"u"， 就是 -u 指定参数
		"" ，默认值
		"用户名，默认为空" 说明
	*/
	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","密码，默认为空")
	flag.StringVar(&host,"h","localhost","密码，默认为空")
	flag.IntVar(&port,"port",3306,"端口号，默认是3306")
	//这里有一个非常重要的操作，转换，必须调用该方法
	flag.Parse()
	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
}