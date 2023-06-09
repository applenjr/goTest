package main

/* 匿名变量 */
import "fmt"

/*
在Go语言中，使用_（下划线）符号来声明匿名变量。
匿名变量的特点是在定义时不给变量命名，所以程序不会为它们分配内存空间，也就是说，匿名变量不占用内存。
在函数返回值中，如果不需要使用所有的返回值，可以使用匿名变量承接不需要的返回值。
*/
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

/*
在main函数中，我们只需要计算出商的值，而不需要计算余数的值。
因此，使用匿名变量_来接收余数的返回值，这样可以避免编译器报错。使用匿名变量可以让代码更简洁、更清晰。
需要注意的是，匿名变量只能出现在赋值语句的右边，也就是作为接收者使用。如果在赋值语句的左边使用匿名变量，会导致编译错误。
*/
func main() {
	x, _ := divide(10, 3)
	fmt.Println(x) // 输出 3
}
