package main
import (
   "strconv";
   "fmt"
)


func main(){
   var str string = "true"
	var b bool

   //1.strconv.ParseBool(str) 函数会返回两个值（value bool，err error）
   //2.因为我只想获取到value bool ,不想获取err 所以我是用_忽略
	b,_=strconv.ParseBool(str)
   fmt.Printf("b type %T b=%v\n",b,b)

   var str2 string = "1234567"
   var n1 int64
   var n2 int

   n1,_=strconv.ParseInt(str2,10,64)
   fmt.Printf("n1 type %T n1 =%v\n",n1,n1)
   n2=int(n1)
   fmt.Printf("n2 type %T n2 = %v\n",n2,n2)

   var str3 string = "123.456"
   var f1 float64
   f1,_=strconv.ParseFloat(str3,64)
   fmt.Printf("f1 type %T f1 = %v\n",f1,f1)
}