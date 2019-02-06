package main
import "fmt"


type Goods struct{
	Name string
	Price int
}


type Book struct {
	Goods //这里就是嵌套匿名结构体Goods
	Writer string
}

/*
	继承给编程带来的遍历
		代码的复用性提高了
		代码的扩展性和维护性提高了

		注：结构体可以使用嵌套匿名结构体所有的字段和方法，即： 首字母大写或者小写的字段，方法，都可以使用。
*/
type A struct {
	Name string
	age int
}

func (a A) SayOk(){
	fmt.Println("A SayOk",a.Name)
}

func (a *A) hello(){
	fmt.Println("A hello",a.Name)
}

type B struct {
	A
}

/*
	1.当我们直接通过 b 访问字段或方法，其执行流程如下 比如：b.Name
	2.编译器会先看b对应的类型有没有Name,如果有，则直接调用B类型的Name字段
	3.如果没有就去看B中嵌套的匿名结构体A 有没有声明Name字段，如果有就调用，如果没有继续查找..如果都找不到就报错。

	4.当结构体和匿名j结构体有相同的字段或则方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
	5.结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错。
	6.如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
	7.嵌套匿名结构体，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
*/
func ABtest(){
	var b B
	b.A.Name = "Tom"
	b.A.age = 19
	b.A.SayOk()
	b.A.hello()

	//如上写法也可以简化为如下写法
	b.Name = "smith"
	b.age = 20
	b.SayOk()
	b.hello()

}

type A1 struct {
	Name string
	age int
}

type B1 struct {
	Name string
	Score float64
}

type C struct {
	A1
	B1
}

func A1B1Test(){
	var c C
	//如果c 没有Name 字段，而 A 和 B 有Name，这时就必须通过指定的匿名结构体名字来区分
	//所以c.Name 就会包含编译错误，这个规则对方法也是一样的！
	// c.Name = "Tom"//报错
	c.A1.Name = "Tom" //正确
	fmt.Println(c)
}

type D struct{
	a A //有名结构体  组合关系
}

func DTest(){
	var d D
	d.a.Name = "jack"
}

// 7.嵌套匿名结构体，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
type Goods2 struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods2
	Brand
}

type TV2 struct {
	*Goods2
	*Brand
}

func Goods2TVTest(){
	tv := TV{ Goods2{"电视机001",5000.99},Brand{"海尔","三东"}}
	tv2 := TV{
		Goods2{
			Price : 5000.99,
			Name : "电视机002",
		},
		Brand{
			Name : "夏普",
			Address : "北京",
		},
	}
	fmt.Println("tv",tv)
	fmt.Println("tv2",tv2)


	tv3 := TV2{ &Goods2{"电视机003",7000.99},&Brand{"创维","河南"}}
	tv4 := TV2{
		&Goods2{
			Name : "电视机004",
			Price : 9000.99,
		},
		&Brand{
			Name : "长虹",
			Address : "四川",
		},
	}
}





func main(){
	ABtest()
}


