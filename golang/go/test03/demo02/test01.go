package main
import (
	"fmt"
)

/*
	斐波那契数列	
*/
func fbn(n int)([]uint64){
	//声明一个切片，切片大小 n
	fbnSlice := make([]uint64,n)
	//第一个数和第二个数的斐波那契 为 1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//进行for 循环来存放斐波那契的数列
	for i:=2;i<n;i++{
		fbnSlice[i] = fbnSlice[i - 1] +fbnSlice[i - 2]
	}
	return fbnSlice
}

/*
	排序
		描述： 排序是将一组数据，依指定的顺序进行排列的过程。
		排序的分类：
			1.内部排序
				指将需要处理的所有数据都加载到内部存储器中j进行排序。包括（交换式排序法、选择式排序法和插入式排序法）
			2.外部排序
				数据量太大，无法将全部加载到内存中，需要借助外部存储进行排序。包括（合并排序法和直接合并排序法）
*/
//冒泡排序实现
func BubbleSort(arr *[5]int){
	fmt.Println("排序前arr=",(*arr))
	temp:=0 //临时变量（用于做交换）

	//冒泡排序
	for i:=0; i<len(*arr) -1; i++ {
		for j:=0; j<len(*arr) - 1 - i; j++ {
			if(*arr)[j] > (*arr)[j + 1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=",(*arr))
}


/*
	查找
	 查找分类：
		 1.顺序查找
		 2.二分查找（该数组是有序）
*/
func findTest(){
	//代码
	names:= [4]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroName)

	//顺序查找：第一种方式
	for i:=0; i<len(names);i++{
		if heroName == names[i]{
			fmt.Printf("找到了%v,下标%v \n",heroName,i)
			break
		}else if i == (len(names) -1) {
			fmt.Printf("没有找到%v\n",heroName)
		}
	}

	//顺序查找： 第二种方式（推荐...）
	index:=-1

	for i :=0;i<len(names);i++{
		if heroName == names[i]{
			index = i  //将找到的值对应的下标赋给index
			break
		}
	}
	if index != -1{
		fmt.Printf("找到了%v,下标%v \n",heroName,index)
	}else{
		fmt.Printf("没有找到%v\n",heroName)
	}
}


/*
	二分查找法
		二分查询的代码实现
*/





func main(){
	// fnbSlice:=fbn(20)
	// fmt.Println("fnbSlice=",fnbSlice)

	//定义数组
	// var arr = [5]int{24,69,80,57,13}
	// //将数组传递给一个函数
	// BubbleSort(&arr)
	// fmt.Println("main arr = ",arr)


	//查找
	findTest()

}



