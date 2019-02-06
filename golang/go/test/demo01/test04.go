package main
import (
	"fmt";
	"math/rand";
	"time"
)

func main(){

	//1.使用 for-range方式
	str:="hello world,北京"
	for index,val:=range str{
		fmt.Println("index=%d, val=%c \n",index,val)
	}
	//2.原始方式
	var str2 string = "hello world!"
	for i:=0;i<len(str2);i++{
		fmt.Println("%c \n",str[i])
	}
	//3.包含汉字的情况[]rune
	// var str3 string = "hello world,北京"
	// str4:=[]rune(str3)
	// for i:=0;i<len(str4);i++{
	// 	fmt.Println("%c \n",str[i])
	// }

	var count int = 0
	for{
		rand.Seed(time.Now().UnixNano())
		n:=rand.Intn(100)+1
		fmt.Println("n=",n)
		count++
		if(n==99){
			break;
		}
		fmt.Printf("生成99 一共用了 %d 次\n",count)
	}
}


