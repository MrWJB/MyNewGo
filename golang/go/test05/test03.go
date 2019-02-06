// package main
// import (
// 	"fmt"
// 	"encoding/json"
// )





// /*
// 	JSON (JavaScrip Object Notation)  是一种轻量级的数据交换格式。 易于人阅读和编写。同时也易于机器解析和生成。key-val
// 	json 是在2001年开始推广使用的数据格式，目前已经成为主流的数据格式。
// 	json 易于机器解析和生成，并有效的地提升网络传输效率，通常程序在网络传输时会先将	数据（结构体，map等）序列化成json字符串，到接收方得到json字符串时，在反序列化恢复到原来的数据类型（结构体，map等）。这种方式已然成为各个语言的标准

// 			   序列化				  网络传输			 反序列化
// 	golang -----------> json字符串 ----------> 程序 ------------> 其它语言

// */

// /*
// 	struct 结构体序列化
// */
// //定义一个结构体
// type Monster struct {
// 	Name string `json:"monster_name"`
// 	Age int `json:"monster_age"`
// 	Birthday string	`json:"monster_birthday"`
// 	Sal float64	`json:"monster_sal"`
// 	Skill string	`json:"monster_skill"`
// }

// func testStruct(){

// 	//演示
// 	monster := Monster{
// 		Name : "牛魔王",
// 		Age : 300,
// 		Birthday : "2011-10-12",
// 		Sal : 8000.0,
// 		Skill : "牛魔拳",
// 	}

// 	//将monster 序列化
// 	data,err := json.Marshal(&monster)
// 	if err != nil {
// 		fmt.Printf("序列号错误 err=%v\n",err)
// 	}
// 	//输出序列化后的结果
// 	fmt.Printf("monster 序列化后=%v\n",data)
// }


// /*
// 	将map进行序列化
// */
// func testMap(){
// 	//定义一个map
// 	var a map[string]interface{}
// 	//使用map需要make
// 	a = make(map[string]interface{})
// 	a["name"] = "红孩儿"
// 	a["age"] = 30
// 	a["address"] = "洪崖洞"

// 	//将a这个map进行序列化
// 	//将monster序列化
// 	data,err := json.Marshal(a)
// 	if err != nil {
// 		fmt.Printf("序列号错误 err=%v\n",err)
// 	}
// 	//输出序列化后的结果
// 	fmt.Printf("map 序列化后=%v\n",data)
// }

// /*
// 	对切片j进行序列化，我们这个切片 []map[string]interface{}
// */
// func testSlice(){
// 	var slice []map[string]interface{}
// 	var m1 map[string]interface{}
// 	var m2 map[string]interface{}

// 	//使用map前需要make
// 	m1 = make(map[string]interface{})
// 	m1["name"] = "Jack"
// 	m1["age"] = "7"
// 	m1["address"] = "北京"
// 	slice = append(slice,m1)

// 	m2 = make(map[string]interface{})
// 	m2["name"] = "tom"
// 	m2["age"] = "20"
// 	m2["address"] = [2]string{"墨西哥","夏威夷"}
// 	slice = append(slice,m2)

// 	//将切片进行序列化
// 	data,err := json.Marshal(slice)
// 	if err != nil {
// 		fmt.Printf("序列号错误 err=%v\n",err)
// 	}
// 	//输出序列化后的结果
// 	fmt.Printf("slice 序列化后=%v\n",data)
 
// }


// /*
// 	对基本数据类型进行序列化意义不大
// */
// func testFloat64(){
// 	var ft float64 = 32.123

// 	data,err := json.Marshal(ft)
// 	if err != nil {
// 		fmt.Printf("序列号错误 err=%v\n",err)
// 	}
// 	//输出序列化后的结果
// 	fmt.Printf("基本数据类型 序列化后=%v\n",data)
// }

// /*								------------------------									*/
// //将json字符串，反序列化成struct
// func unmarshalStruct(){
// 	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2018-01-21\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
// 	//定义一个monster实例
// 	var monster Monster
// 	err := json.Unmarshal([]byte(str),&monster)
// 	if err != nil {
// 		fmt.Printf("unmarshal err=%v\n",err)
// 	}
// 	fmt.Printf("反序列化 monster=%v monster.Name=%v",monster,monster.Name)
// }

// /*
// 	单元测试
// 		go 语言中自带有一个轻量级的测试框架testing 和自带的 go test 命令来实现单元测试和性能测试，testing 框架和其它语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于框架写相应的压力测试用例。通过单元测试，可以解决如下问题：
// 	快速入门：	
// */
// func 







// func main(){
// 	// testStruct()
// 	// testMap()
// 	// testSlice()
// 	// testFloat64()
// 	unmarshalStruct()
// }