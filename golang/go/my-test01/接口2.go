package main
import "fmt"


type Man interface {
	name() string;
	age() int;
}

type Woman struct{
}

/*
	实现name方法
*/
func (woman Woman) name() string{
	return "Jin Yawei"
}

/*
	实现age方法
*/
func (woman Woman) age() int{
	return 23
}

/*
*/
type Men struct{
}

func (men Men) name() string{
	return "liweibin"
}

func (men Men) age() int{
	return 23
}


func main(){
	var man Man
	man = new(Woman)
	fmt.Println(man.name())
	fmt.Println(man.age())
	man = new(Men)
	fmt.Println(man.name())
	fmt.Println(man.age())

}
