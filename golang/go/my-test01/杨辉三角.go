package main
import "fmt"

func GetYangHuiTriangleNextLine(inArr []int) []int {
    var out []int
    var i int
    arrLen := len(inArr)
    fmt.Printf("获取inArr的长度：%d\n",arrLen)
    out = append(out, 1)
    if 0 == arrLen {
        return out
    }
    fmt.Println("获取的out前的值：",out)
    for i = 0; i < arrLen-1; i++ {
        out = append(out, inArr[i]+inArr[i+1])
    }
    out = append(out, 1)
    fmt.Println("获取的out后的值：",out)
    return out
}

func main() {
    nums := []int{}
    var i int
    for i = 0; i < 10; i++ {
        nums = GetYangHuiTriangleNextLine(nums)
        fmt.Println(nums)
    }
}