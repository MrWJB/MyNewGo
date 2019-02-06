package main
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

/*
	单元测试综合案例要求：
		1.编写一个Monster 结构体，字段Name,Age,Skill
		2.给Monster绑定方法Store，可以将一个Monster变量(对象)，序列化后保存到文件中
		3.给Monster绑定方法ReStore,可以将一个序列化的Monster，从文件中读取，并反序列化为Monster对象，检查反序列化，名字正确
		4.编程测试用例文件store_test.go,编写测试用例函数TestStore 和 TestRestore 进行测试
*/
type Monster struct {
	Name string
	Age int
	Skill string
}


//给Monster 绑定方法Store,可以将一个Monster变量（对象）,序列化后保存到文件中
func (this *Monster) Store() bool {
	//先序列化
	data,err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err=",err)
		return false
	}
	//保存到文件
	filePath := "d:/monster.ser"
	err = ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Println("write file err=",err)
		return false
	}
	return true
}

/*
	给Monster绑定方法Restore,可以将一个序列化的Monster，从文件中读取，并反序列化Monster对象，检查反序列化，名字正确
*/
func (this *Monster) ReStore() bool {
	//1.先从文件中，读取序列化的字符串
	filePath := "d:/monster.ser"
	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err=",err)
		return false
	}
	//进行反序列
	err = json.Unmarshal(data,this)
	if err != nil {
		fmt.Println("UnMarshal err=",err)
		return false
	}
	return true
}













