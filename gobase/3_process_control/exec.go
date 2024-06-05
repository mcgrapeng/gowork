package process_control

import "fmt"

func Exec() {
	/**
	if else(分支结构)
	*/
	fmt.Println(`
	1.【if else】
		(1).标准写法：
			if 表达式1 {
				分支1
			} else if 表达式2 {
				分支2
			} else{
				分支3
			}
		
		举个例子：
			func ifDemo1() {
				score := 65
				if score >= 90 {
					fmt.Println("A")
				} else if score > 75 {
					fmt.Println("B")
				} else {
					fmt.Println("C")
				}
			}

		(2).特殊写法：
			可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断。
			
			举个例子：
				func ifDemo2() {
					if score := 65; score >= 90 {
						fmt.Println("A")
					} else if score > 75 {
						fmt.Println("B")
					} else {
						fmt.Println("C")
					}
				}
	`)

	/**
	for(循环结构)
	*/
	fmt.Println(`
	2.【for】
		(1).标准写法：
			for 初始语句;条件表达式;结束语句{
				循环体语句
			}
		
		举个例子：
			func forDemo() {
				for i := 0; i < 10; i++ {
					fmt.Println(i)
				}
			}

		(2).特殊写法：
			1️⃣for循环的初始语句可以被忽略，但是初始语句后的分号必须要写。
			举个例子：
				func forDemo2() {
					i := 0
					for ; i < 10; i++ {
						fmt.Println(i)
					}
				}
			2️⃣for循环的初始语句和结束语句都可以省略，这种写法类似于其他编程语言中的while。
			举个例子：
				func forDemo3() {
					i := 0
					for i < 10 {
						fmt.Println(i)
						i++
					}
				}

		(3).无限循环：
			for {
				循环体语句
			}

		(4).循环终止符
			for循环可以通过break、goto、return、panic语句强制退出循环。
	`)

	/**
	for range(键值循环)
	*/
	fmt.Println(`
	3.【for range】
		(1).作用：
			Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。
		(2).用法：
			1️⃣数组、切片、字符串返回索引和值。
			2️⃣map返回键和值。
			3️⃣通道（channel）只返回通道内的值。
	`)

	/**
	switch case
	*/
	fmt.Println(`
	4.【switch case】
		(1).标准写法：
			Go语言规定每个switch只能有一个default分支。

			switch 变量 {
				case 变量值1:
					fmt.Println("大拇指")
				case 变量值2:
					fmt.Println("食指")
				case 变量值3:
					fmt.Println("中指")
				.....

				default:
					fmt.Println("无效的输入！")
				}
		
		举个例子：
			func switchDemo1() {
				finger := 3
				switch finger {
				case 1:
					fmt.Println("大拇指")
				case 2:
					fmt.Println("食指")
				case 3:
					fmt.Println("中指")
				case 4:
					fmt.Println("无名指")
				case 5:
					fmt.Println("小拇指")
				default:
					fmt.Println("无效的输入！")
				}
			}

		(2).特殊写法：
			1️⃣一个分支可以有多个值，多个case值中间使用英文逗号分隔。
			举个例子：
				func testSwitch3() {
					switch n := 7; n {
					case 1, 3, 5, 7, 9:
						fmt.Println("奇数")
					case 2, 4, 6, 8:
						fmt.Println("偶数")
					default:
						fmt.Println(n)
					}
				}
			2️⃣分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。
			举个例子：
				func switchDemo4() {
					age := 30
					switch {
					case age < 25:
						fmt.Println("好好学习吧")
					case age > 25 && age < 35:
						fmt.Println("好好工作吧")
					case age > 60:
						fmt.Println("好好享受吧")
					default:
						fmt.Println("活着真好")
					}
				}
		(3).fallthrough
			定义：可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

			举个例子：
				func switchDemo5() {
					s := "a"
					switch {
					case s == "a":
						fmt.Println("a")
						fallthrough
					case s == "b":
						fmt.Println("b")
					case s == "c":
						fmt.Println("c")
					default:
						fmt.Println("...")
					}
				}
			输出：
				a
				b
	`)

	/**
	goto(跳转到指定标签)
	*/
	fmt.Println(`
	5.【goto】
		(1).作用：
			goto语句通过标签进行代码间的无条件跳转。
	`)

	/**
	break(跳出循环)
	*/
	fmt.Printf(`
	6.【break】
		(1).作用：
			break语句可以结束for、switch和select的代码块。
			break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。

		举个例子：
			func breakDemo1() {
			BREAKDEMO1:
				for i := 0; i < 10; i++ {
					for j := 0; j < 10; j++ {
						if j == 2 {
							break BREAKDEMO1
						}
						fmt.Printf("%v-%v\n", i, j)
					}
				}
				fmt.Println("...")
			}
	`)

	/**
	continue(继续下次循环)
	*/
	fmt.Printf(`
	7.【continue】
		(1).作用：continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。

		举个例子：
			func continueDemo() {
			forloop1:
				for i := 0; i < 5; i++ {
					// forloop2:
					for j := 0; j < 5; j++ {
						if i == 2 && j == 2 {
							continue forloop1
						}
						fmt.Printf("%v-%v\n", i, j)
					}
				}
			}
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	if else
	*/
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	/**
	for循环结构
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for i < 10 {
		fmt.Println(j)
		j++
	}

	/**
	switch case
	*/
	city := 5
	switch city {
	case 1:
		fmt.Println("南京")
	case 2:
		fmt.Println("苏州")
	case 3:
		fmt.Println("杭州")
	case 4:
		fmt.Println("成都")
	default:
		fmt.Println("西安")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	age := 30
	switch {
	case age < 25:
		fmt.Println("祖国花朵")
	case age > 25 && age < 35:
		fmt.Println("韭菜")
		fallthrough //fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	case age > 60:
		fmt.Println("老年打工人")
	default:
		fmt.Println("体制内")
	}

	/**
	break：break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上
	*/
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("...")

	/**
	goto：跳转到指定标签
	*/
	//传统写法
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Println(i, j)
		}

		if breakFlag {
			break
		}
	}

	//使用goto
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakFlag
			}
		}
	}
breakFlag:
	fmt.Println("结束for循环")

	/**
	continue
	*/
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
	}

	continueDemo()

}

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func breakDemo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
