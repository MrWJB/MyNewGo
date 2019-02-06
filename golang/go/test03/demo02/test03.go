package main
import (
	"fmt"
)


/*
	map 是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map
*/
func modify(map1 map[int]int){
	map1[10] = 900
}



/*
	map 的容量达到后，再想map增加元素，会自动扩容，并不会发生panic,也就是说map能动态的增长 键值对（key-value）
	map 的value 也经常使用struct 类型，更适合管理复杂的数据（比前面value是一个map更好），比如value是Student 结构体【案例演示，因为还没有学习j结构体，体验一下即可】
*/
func mapStruct(){
	// students := make(map[string]Stu,10)
	// //创建两个学生
	// stu1 := Stu{"tom",18,"北京"}
	// stu2 := Stu{"mary",28,"上海"}

	// students["no1"] = stu1
	// students["no2"] = stu2

	// fmt.Println(students)

	// //遍历各个学生信息
	// for k,v := range students {
	// 	fmt.Printf("学生的编号是%v \n",k)
	// 	fmt.Printf("学生的名字是%v \n",v.Name)
	// 	fmt.Printf("学生的年龄是%v \n",v.Age)
	// 	fmt.Printf("学生的地址是%v \n",v.Address)

	// 	fmt.Println()
	// }
}

/*
	定义一个Cat结构体，将Cat 的各个字段/属性信息，放入到Cat结构体进行管理
*/
type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
}

/*
	面向对象编程 ： 结构体
*/
func structTest(){
	//创建一个Cat的变量
	var cat Cat
	cat.Name = "江小白"
	cat.Age = 10
	cat.Color = "白色"
	cat.Hobby = "吃<>"
	fmt.Println("cat=",cat)

	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=",cat.Name)
	fmt.Println("Age=",cat.Age)
	fmt.Println("color=",cat.Color)
	fmt.Println("hobby=",cat.Hobby)

}

/*
	如果结构体的字段类型是： 指针，slice 和map 的零值都是nil，即还没有分配空间
	如果需要使用这样的字段，需要先make,才能使用
*/
type Person struct{
	Name string
	Age int
	Scores [5]float64
	prt *int //指针
	slice []int //切片
	map1 map[string]string  //map

}

func PersonTest(){
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.prt == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用slice,再次说明，一定要make 
	p1.slice = make([]int,10)
	p1.slice[0] = 100

	//使用map ，一定要先make 
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"

	fmt.Println(p1)
}



// func main(){
// 	map1 := make(map[int]int)
// 	map1[1] = 90
// 	map1[2] = 88
// 	map1[10] = 1
// 	map1[20] = 2

// 	modify(map1)
// 	//看看结果，map1[10] = 900 ，说明map是引用类型
// 	fmt.Println(map1)
// }

/*
	创建结构体变量和访问结构体字段
*/
func PersonDef(){
	// //方式一 直接声明
	// var person Person
	// //方式二 
	// var person2 Person = Person{}
	// //方式三 返回的是 结构体指针
	// var person3 *Person = new(Person)
	// *person3.Name = "Smith"
	// person3.Name = "Smith"

	// *person3.Age = 30 
	// person3.Age = 100
	// fmt.Println(*person3)

	//返回的是结构体指针
	var person *Person = &Person{}


}

/*
	struct 的每个字段上，可以写上一个tag,该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
*/
type Monster struct{
	Name string `json:"name"` //`json:"name"` 就是struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func MonsterTest(){
	// //创建一个Monster变量
	// monster := Monster{"牛魔王",500,"芭蕉扇~"}
	// //将monster 变量序列化为json格式字符串
	// //json.Marshal 函数中使用反射，这个讲解反射时，我会详细介绍
	// jsonStr,err := json.Marshal(monster)
	// if err != nil {
	// 	fmt.Println("json 处理错误 ",err)
	// }
	// fmt.Println("jsonStr",string(jsonStr))
}




func main(){
	// structTest()
	PersonTest()
}