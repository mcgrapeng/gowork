package funcZhang

import (
	"errors"
	"fmt"
)

// 全局变量
var num int64 = 10

func Exec() {
	fmt.Println(`
	1.【函数】
		(1).地位：
			Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
		
		(2).格式：
			func 函数名(参数)(返回值){
				函数体
			}
			
			其中：
				函数名：	由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
				参数：	参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
				返回值：	返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
				函数体：	实现指定功能的代码块。

			

			举个例子：
				func intSum(x int, y int) int {
					return x + y
				}

				//函数的参数和返回值都是可选的，例如我们可以实现一个既不需要参数也没有返回值的函数：
				func sayHello() {
					fmt.Println("Hello 沙河")
				}

		(3).调用：
			可以通过【函数名()】的方式调用函数。
			
			举个例子：
				sayHello()
				ret := intSum(10, 20)
				fmt.Println(ret)

		(4).函数其他格式：
			1️⃣参数类型简写
				举个例子：
					func intSum(x, y int) int {
						return x + y
					}

			2️⃣可变参数（固定参数搭配可变参数使用时，可变参数要放在固定参数的后面）
				举个例子：
					func intSum3(x int, y ...int) int {
						fmt.Println(x, y)
						sum := x
						for _, v := range y {
							sum = sum + v
						}
						return sum
					}

			3️⃣多返回值
				举个例子：
					func calc(x, y int) (int, int) {
						sum := x + y
						sub := x - y
						return sum, sub
					}

			4️⃣返回值命名
				举个例子：
					func calc(x, y int) (sum, sub int) {
						sum = x + y
						sub = x - y
						return
					}

			5️⃣返回值nil
				举个例子：
					func someFunc(x string) []int {
						if x == "" {
							return nil // 没必要返回[]int{}
						}
						...
					}
	`)

	fmt.Println(`
	2.【函数进阶】
		(1).全局变量：
			全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

		(2).局部变量：
			函数内定义的变量无法在该函数外使用，如果局部变量和全局变量重名，优先访问局部变量。

		(3).定义函数类型：
			1️⃣格式：
				type 函数名 func(参数类型, 参数类型) 返回类型
			
			2️⃣使用：
				type calculation func(int, int) int

				func add(x, y int) int {
					return x + y
				}
				
				func sub(x, y int) int {
					return x - y
				}
				
				var c calculation
				c = add
			
			3️⃣作为参数：
				func add(x, y int) int {
					return x + y
				}
				func calc(x, y int, op func(int, int) int) int {
					return op(x, y)
				}
				ret2 := calc(10, 20, add)

			4️⃣作为返回值：
				func do(s string) (func(int, int) int, error) {
					switch s {
					case "+":
						return add, nil
					case "-":
						return sub, nil
					default:
						err := errors.New("无法识别的操作符")
						return nil, err
					}
				}
	`)

	fmt.Println(`
	3.【匿名函数和闭包】
		(1).匿名函数：
			1️⃣定义：
				在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数。
				匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数。

			2️⃣格式：
				func(参数)(返回值){
					函数体
				}
			
			举个例子：
				// 将匿名函数保存到变量
				add := func(x, y int) {
					fmt.Println(x + y)
				}
				add(10, 20) // 通过变量调用匿名函数
			
				//自执行函数：匿名函数定义完加()直接执行
				func(x, y int) {
					fmt.Println(x + y)
				}(10, 20)
			
			3️⃣用途：
				匿名函数多用于实现回调函数和闭包。

		(2).闭包：
			通过以下的例子可以看明白。

			func adder2(x int) func(int) int {
				return func(y int) int {
					x += y
					return x
				}
			}
			
			var f = adder2(10)
			fmt.Println(f(10)) //20
			fmt.Println(f(20)) //40
			fmt.Println(f(30)) //70
		
			f1 := adder2(20)
			fmt.Println(f1(40)) //60
			fmt.Println(f1(50)) //110
	`)

	fmt.Println(`
	4.【defer】
		(1).定义：
			Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
			在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

			举个例子：
				fmt.Println("start")
				defer fmt.Println(1)
				defer fmt.Println(2)
				defer fmt.Println(3)
				fmt.Println("end")

			输出：
				start
				end
				3
				2
				1

		(2).作用：
			由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
			比如：资源清理、文件关闭、解锁及记录时间等。

		(3).执行时机：
			在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
			而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。

	`)

	fmt.Println(`
	5.【内置函数】
		(1).介绍：
			close			主要用来关闭channel
			len				用来求长度，比如string、array、slice、map、channel
			new				用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
			make			用来分配内存，主要用来分配引用类型，比如chan、map、slice
			append			用来追加元素到数组、slice中
			panic和recover	用来做错误处理
	`)

	fmt.Println(`
	6.【异常处理】
		(1).机制：
			Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 
			panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

			举个例子：
				func funcB() {
					defer func() {
						err := recover()
						//如果程序出出现了panic错误,可以通过recover恢复过来
						if err != nil {
							fmt.Println("recover in B")
						}
					}()
					panic("panic in B")
				}

		(2).注意事项：
				1️⃣recover()必须搭配defer使用。
				2️⃣defer一定要在可能引发panic的语句之前定义。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	fmt.Println(num)
	//局部变量
	num := 100
	fmt.Println(num)

	/**
	函数调用
	*/
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)

	/**
	多返回值
	*/
	fmt.Println(`

		func calc(x, y int) (int, int) {
			sum := x + y
			sub := x - y
			return sum, sub
		}
		
		对返回值命名
		func calc(x, y int) (sum, sub int) {
			sum = x + y
			sub = x - y
			return
		}
		
		返回nil，相当于java的null
		func calc(x string) []int{
			if x == ""{
				return nil
			}
			return make([]int, 0)
		}
	`)

	/**
	定义函数类型：type 函数类型 func(int, int) int
	*/
	fmt.Println(`
		type calc func(int, int) int
	`)

	var c calc
	c = intSum
	fmt.Printf("type of c:%T\n", c)
	fmt.Println(c(1, 2))

	/**
	匿名函数
	func(参数)(返回值){
	    函数体
	}
	*/
	add := func(x, y int) {
		fmt.Println(x + y)
	}

	/**
	匿名函数调用
	*/
	add(10, 20)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	/**
	闭包
	*/
	fmt.Println(`
		func adder() func(int) int {
			var x int
			return func(y int) int {
				x += y
				return x
			}
		}
	`)

	/**
	defer
	Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
	在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
	也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

	由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
	比如：资源清理、文件关闭、解锁及记录时间等。
	*/
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

	f1()
	f2(10)
	f3()
	f4(10)

	/**
	panic/recover
	Go语言中目前没有异常机制，但是使用panic/recover模式来处理错误。
	panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

	注意：
		1.recover()必须搭配defer使用。
		2.defer一定要在可能引发panic的语句之前定义。
	*/
	funcA()
	funcB()
	funcC()

}

func sayHello() {
	fmt.Println("Hello 成都")
}
func intSum(x int, y int) int {
	return x + y
}

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

// 函数作为参数
func add2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func add3(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	default:
		err := errors.New("无法识别")
		return nil, err
	}
}

func add4(s string) (int, error) {
	switch s {
	case "A":
		return 0, nil
	default:
		err := errors.New("无法识别")
		return 1, err
	}
}

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
		fmt.Println("f1 start: ", x)
	}()
	fmt.Println("f1 end:", x)
	return x
}

func f2(x int) {
	defer func() {
		x++
		fmt.Println("f2 start: ", x)
	}()
	fmt.Println("f2 end:", x)
	return
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
		fmt.Println("f3 start: ", x)
	}()
	fmt.Println("f3 end: ", x)
	return x
}
func f4(x int) {
	defer func(x int) {
		x++
		fmt.Println("f4 start: ", x)
	}(x)
	fmt.Println("f4 end: ", x)
	return
}
func funcA() {
	fmt.Println("funcA")
}
func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic in B")
}
func funcC() {
	fmt.Println("funcC")
}
