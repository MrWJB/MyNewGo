package main
import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
)

/*
	File 文件操作
*/
func FileTest(){
	//打开文件
	//概念说明：file 的叫法
	//1. file 叫 file 对象
	//2. file 叫 file 指针
	//3. file 叫 file 文件句柄
	file,err := os.Open("G:/learn/XW00001NYYH_acct_20191207_001.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	//输出下文件，看看文件是什么，看出file 就是一个指针 *File
	fmt.Printf("file=%v",file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
}


/*
	1.读取文件的内容并显示在终端（带缓冲区的方式）, 使用os.Open(),file.Close,bufio.NewReader(),reader.ReadString 函数和方法
	2.读取文件的内容并显示在终端（使用ioutil一次将整个文件读入到内存中），这种方式适用于文件不大的情况，相关f方法和函数（ioutil.ReadFile）
*/
func fileRead(){
	//使用ioutil.ReadFile 一次性将文件读取到内存
	file := "d:/test.txt"
	content,err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v",err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v",content) // []byte
	fmt.Printf("%v",string(content))

	//我们没有显示的Open文件，因此也不需要显示的Close文件
	//因为，文件的Open和Close被封装到ReadFile 函数内部

}

/*
	基本应用实例 - 方式一
*/
func outFile(){
	//创建一个新文件，写入内容 5句 “hello,Gardon”
	//1.打开文件 d:/abc.txt
	filePath := "d:/abc.txt"
	// file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666)
	//打开一个存在的文件中，将原来的内容覆盖成新的内容10句 ， "你好，尚硅谷！"
	// file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_TRUNC,0666)
	//打开一个存在的文件，在原来的内容追加内容 'ABC!ENGLISH!'
	// file,err := os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND,0666)
	//打开一个已经存在的文件，将原来的内容读取出来并显示在终端，并且追加5句"hello,北京！
	file,err := os.OpenFile(filePath,os.O_RDWR | os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入5句 "hello , Gardon"
	str := "hello , Gardon"
	writer := bufio.NewWriter(file)
	for i := 0;i< 5;i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用writerString方法时，其实内容是先写入到缓存的，所以需要调用Flush方法，将缓冲的数据真正写入到文件中，否则文件中会没有数据！！！
	writer.Flush()
}


/*
	基本应用案例 - 方式二
		编写一个程序，将一个文件的内容，写入到另外一个文件。 注： 这两个文件已经存在了。
		说明： 使用ioutil.ReadFile / ioutil.WriteFile 完成写文件的任务
*/
func ReadWriter(){
	//将d:/abc.txt 文件内容导入到 e:/kkk.txt
	//1.首先将 d:/abc.txt 内容读取到内存
	//2.将读取到的内容写入 e:/kkk.txt
	filePath := "d:/abc.txt"
	file2Path := "e:/kkk.txt"
	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err=%v\n",err)
		return
	}
	err = ioutil.WriteFile(file2Path,data,0666)
	if err != nil {
		fmt.Printf("read file err=%v\n",err)
	}



}

/*
	判断文件是否存在
		golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断：
			1.如果返回的错误为nil,说明文件或文件夹存在
			2.如果返回的错误类型使用 os.IsNotExist() 判断为true，说明文件或文件夹不存在
			3.如果返回的错误为其它类型，则不确定是否在存在
*/
//自己写一个函数
func PathExists(path string)(bool,error){
	_,err := os.Stat(path)
	if err == nil {   //文件或者目录存在
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}















func main(){
	FileTest()
}