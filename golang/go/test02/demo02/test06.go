package main
import (
	"fmt";
	"strconv";
	"strings"
)

func sum(num1 int ,num2 int) int {
	defer fmt.Println("ok1 n1=",num1)
	defer fmt.Println("ok2 n2=",num2)
	res:=num1 + num2
	fmt.Println("ok3 res=",res)
	return res
}

/*
	循环遍历字符串
*/
func strEach(){
	var str2= "hello 北京"
	r:=[]rune(str2)
	for i :=0;i<len(r);i++{
		fmt.Printf("字符=%c\n",r[i] )
	}
}
/*
	字符串转为整数
*/
func strToInt(){
	res,err:=strconv.Atoi("hello")
	if err!=nil{
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("转换结果",res)
	}
}

/*
	整数转为字符串
*/
func intToStr(){
	str:=strconv.Itoa(12345)
	fmt.Printf("str type %T str %v ",str,str)
}

/*
	字符串转byte字节
*/
func strToByte(){
	var bytes=[]byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)

}

/*
	byte字节转字符串
*/
func byteToStr(){
	str:=string([]byte("hello go"))
	fmt.Printf("str=%v\n",str)
}


/*
	10进制转2，8，16进制：str=strconv.FormatInt(123,2),返回对应的字符串
*/
func tenToTwo(){
	strs:=strconv.FormatInt(123,2)
	fmt.Printf("123对应的二进制是=%v\n",strs)
	
	strs2:=strconv.FormatInt(123,6)
	fmt.Printf("123对应的二进制是=%v\n",strs2)
}

/*
	查找子串是否在指定的字符串中：strings.Contains("seafood","foo")  //true
*/
func strContains(){
	b:=strings.Contains("seafood","mary")
	fmt.Println("b",b)
}

/*
	统计一个字符串有几个指定的子串： strings.Count("ceheese","e")
*/
func strCount(){
	num:=strings.Count("ceheese","e")
	fmt.Printf("num=%v\n",num)
}

/*
	不区分大小写的字符串比较（==是区分字母大小写的）；fmt.Println(strings.EqualFold("abc","Abc"))
*/
func equalNo(){
	b:=strings.EqualFold("abc","Abc")
	fmt.Println("b",b)
}

/*
	返回子串在字符串第一次出现的index值，如果没有返回-1：strings.Index("NLT_abc","abc")
*/
func strIndex(){
	num:=strings.Index("NLT_abc","abc")
	fmt.Println("结果：",num)
}

/*
	返回子串在字符串最后一次出现的index，如果没有返回-1  ； strings.LastIndex("go golang","go")
*/
func lastIndex(){
	num:=strings.LastIndex("go golang","go")
	fmt.Println("num",num)
}

/*
	将指定的子串替换成 另外一个子串： strings.Replace("go go helo","go","go 语言",n) n 可以指定你希望替换几个，如果是n=-1 表示全部替换
*/
func strReplace(){
	strs:=strings.Replace("go go hello","go","go 语言",1)
	fmt.Printf("strs type %T, strs %v",strs,strs)
}

/*
	按照指定的某个字符，为分隔标识，将一个字符串拆分成字符串数组： strings.Split("hello,wrold,ok",",")
*/
func strSplit(){
	str:=strings.Split("hello,world,ok",",")
	for index,val:=range str{
		fmt.Printf("strSplit str type %d str %v",index,val)
	}
}

/*
	将字符串的字母进行大小写的转换： strings.ToLower("Go") //go strings.ToUpper("Go") //GO
*/
func strToLowerOrUpper(){
	// strs:=strings.ToLower("Go")
	strs:=strings.ToUpper("go")
	fmt.Println("strToLowerOrUpper结果：",strs)
}

/*
	将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gohelper ntrn ")
*/
func strTrim(){
	strs:=strings.TrimSpace("  tn a lone gohelper ntrn  ")
	fmt.Println("strTrim获取的结果值:",strs)
}

/*
	将字符串左右两边指定的字符去掉： strings.Trim("! hello!","!")
*/
func strTrimChar(){
	str:=strings.Trim("!hello!","!")
	fmt.Println("strTrimChar的值：",str)
}

/*
	将字符串左边指定的字符去掉： strings.TrimLeft("!hello!","!")
*/
func strTrimLeft(){
	strs:=strings.TrimLeft("!hello!","!")
	fmt.Println("strTrimLeft的值：",strs)
}

/*
	将字符串右边指定的字符去掉; strings.TrimRight("!hello!","!")
*/
func strTrimRight(){
	strs:=strings.TrimRight("!hello!","!")
	fmt.Println("strTrimRight 结果值：",strs)
}

/*
	判断字符串是否以指定的字符串开头：strings.HasPrefix("fpt://192.168.11.100","ftp")
*/
func strHasPrefix(){
	str:=strings.HasPrefix("ftp://192.168.11.100","ftp")
	fmt.Println("strHasPrefix 结果的值：",str)
}

/*
	判断字符串是否以指定的字符串结尾 : strings.HasSuffix("NLT_abc.jpg","abc")
*/
func strHasSuffix(){
	strs:=strings.HasSuffix("NLT_abc.jpg","abc")
	fmt.Println("strHasSuffix 结果的值：",strs)
}



func main(){
	
	// res:=sum(10,20)
	// fmt.Println("res=",res)
	// strEach()
	// strToInt()
	// intToStr()
	// strToByte()
	// byteToStr()
	// tenToTwo()
	// strContains()
	// strCount()
	// equalNo()
	// strIndex()
	// lastIndex()
	// strReplace()
	// strSplit()
	// strToLowerOrUpper()
	// strTrim()
	// strTrimChar()
	// strTrimLeft()
	// strTrimRight()
	// strHasPrefix()
	strHasSuffix()

}