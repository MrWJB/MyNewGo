package main
import "fmt"

// 定义一个 DivideError 结构
type DivideError struct{
	dividee int
	divider int
}

//实现error接口
func (divideError DivideError) Error() string{
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, divideError.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int,varDivider int) (result int,errorMsg string){
	if varDivider == 0 {
		dData := DivideError{
			dividee : varDividee,
			divider : varDivider,
		}
		errorMsg = dData.Error()
		return
	}else{
		return varDividee / varDivider, ""
	}
}

/*
	错误处理
*/
func errorTest(){
	//正常情况下
	if result,errorMsg := Divide(100,10); errorMsg == ""{
		fmt.Println("100/10 =",result)
	}

	// 当被除数为零的时候会返回错误信息
	if _,errorMsg := Divide(10,0);errorMsg != ""{
		fmt.Println("errorMsg is: ", errorMsg)
	}
}


/*
	1. 这里应该介绍一下 panic 与 recover,一个用于主动抛出错误，一个用于捕获panic抛出的错误。
	2. panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误，panic 用于主动抛出错误，
	   recover 用来捕获 panic 抛出的错误。
	3. 引发panic有两种情况，一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。
	4. 发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数。
	5. panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。
	6. recover用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，但是defer只有在后面的函数体内直接被掉用才能捕获panic来终止异常，否则返回nil，异常继续向外传递。
	

*/

func main(){

	errorTest()
}
