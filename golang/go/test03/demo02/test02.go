package main
import (
	"fmt"
	"sort"
)

/*
	二维数组的介绍
*/
func arrayTwo(){
	//定义/声明二维数组
	var arr [4][6]int
	//赋值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组，按照要求输出图形
	for i := 0;i<4;i++{
		for j:=0;j<6;j++{
			fmt.Println(arr[i][j]," ")
		}
		fmt.Println()
	}
}

/*
	二维数组遍历
*/
func arrayTraces(){
	//双层for循环完成遍历
	//for-range方式完成遍历
	var arr3 = [2][3]int{{1,2,3},{4,5,6}}
	//for 循环来遍历
	for i:=0;i<len(arr3);i++{
		for j:=0;j<len(arr3[i]);j++{
			fmt.Printf("%v\t",arr3[i][j])
		}
		fmt.Println()
	}


	//for-range来遍历二维数组
	for i,v:=range arr3{
		for j,v2:=range v{
			fmt.Printf("arr3[%v][%v]=%v \t",i,j,v2)
		}
		fmt.Println()
	}
}

/*
	map 的使用
*/
func mapTest(){
	//方式一
	//map的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make,make 的作用就是给map分配数据空间
	a=make(map[string]string,10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)

	//方式二
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//方式三
	heroes := map[string]string{
		"hero1" : "宋江",
		"hero2" : "卢俊义",
		"hero3" : "吴用",
	}
	heroes["hero4"] = "林冲"
	fmt.Println("heroes=",heroes)
}

/*
	课堂练习：演示一个key-value 的value是map的案例
	比如: 我们要存放3个学生信息，每个学生有 name 和 sex 信息
	思路：map[string]map[string]string
*/
func mapValue(){
	studentMap:= make(map[string]map[string]string)

	studentMap["stu01"] =make(map[string]string,3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string,3) // 这句话不能少
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄埔江~"
	
	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["address"] )

}

/*
	map 的增删改查操作
*/
func mapOperator(){
	//map["key"] = value //如果key还没有，就是增加，如果key存在就是修改
	cities :=make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	fmt.Println(cities)

	//因为no3 这个key 已经存在，因此下面的这句话j就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)
}

/*
	map 删除
*/
func mapDelete(){
	cities :=make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities,"no3")
	fmt.Println(cities)
}

/*
	map 查找
*/
func mapFind(){
	cities :=make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	val,ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n",val)
	}else{
		fmt.Printf("没有no1 key\n")
	}
}

/*
	map 遍历
*/
func mapTravers(){
	//使用for-range 遍历map
	//第二种方式
	cities :=make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	for k,v := range cities{
		fmt.Printf("k=%v v=%v\n",k,v)
	}

	//使用for-range 遍历一个结构比较复杂的map
	studentMap:= make(map[string]map[string]string)

	studentMap["stu01"] =make(map[string]string,3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string,3) // 这句话不能少
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄埔江~"

	for k1,v1 := range studentMap{
		fmt.Println("k1=",k1)
		for k2,v2 := range v1{
			fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		}
		fmt.Println()
	}

}


/*
	map 切片
	基本介绍
		切片的数据类型如果是 map,则我们称为slice of map,map 切片，这样使用则map个数就可以动态变化了。
*/
func mapSlice(){
	//声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string,2) //准备放入两个妖怪
	//增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string,2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] ==nil {
		monsters[1] = make(map[string]string,2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	//这里我们需要使用到切片的append函数，可以动态的增加monster
	//1.定义个monster 信息
	newMonster := map[string]string{
		"name" : "新的妖怪~火云邪神",
		"age" : "200",
	}
	monsters = append(monsters,newMonster)
	fmt.Println(monsters)
}

/*
	golang 中的map 默认是无序的，注意也不是按照添加的顺序存放的，你每次遍历，得到的输出可能不一样。
*/
func mainSort(){
	//map 的排序
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	//如果按照map的key的顺序进行排序输出
	//1.先将map的key 放入到切片中
	//2.对切片排序
	//3.遍历切片，然后按照key来输出map的值

	var keys []int
	for k,_ := range map1 {
		keys=append(keys,k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)
	for _,k := range keys{
		fmt.Printf("map1[%v]=%v \n",k,map1[k])
	}
}










func main(){
	// arrayTwo()
	// arrayTraces()
	// mapTest()
	// mapValue()
	// mapOperator()
	// mapDelete()
	// mapFind()
	// mapTravers()
	// mapSlice()
	mainSort()
}