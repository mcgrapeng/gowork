package interfaceZhang

import "fmt"

func Exec() {

	fmt.Printf(`
	1.【接口】
		(1).定义：
			type 接口类型名 interface{
				方法名1( 参数列表1 ) 返回值列表1
				方法名2( 参数列表2 ) 返回值列表2
				…
			}

			其中：
				接口类型名：	Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。
							接口名最好要能突出该接口的类型含义。
				方法名：		当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
				参数列表、返回值列表：	参数列表和返回值列表中的参数变量名可以省略。


			举个例子：
				type Writer interface{
					Write([]byte) error
				}

		(2).实现：
			接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

			举个例子：
				type Singer interface {
					Sing()
				}
			
				type Bird struct {}

				func (b Bird) Sing() {
					fmt.Println("夹夹夹")
				}

		(3).变量；
			一个接口类型的变量能够存储所有实现了该接口的类型变量。

			举个例子：
				var x Sayer // 声明一个Sayer类型的变量x
				a := Cat{}  // 声明一个Cat类型变量a
				b := Dog{}  // 声明一个Dog类型变量b
				x = a       // 可以把Cat类型变量直接赋值给x
				x.Say()
				x = b       // 可以把Dog类型变量直接赋值给x
				x.Say()

		(4).接收者
			1️⃣值类型接收者：
				使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

				举个例子：
					type Mover interface {
						Move()
					}
					
					type Dog struct{}
					
					func (d Dog) Move() {
						fmt.Println("狗会动")
					}

					var x Mover    // 声明一个Mover类型的变量x
					var d1 = Dog{} // d1是Dog类型
					x = d1         // 可以将d1赋值给变量x
					x.Move()
					
					var d2 = &Dog{} // d2是Dog指针类型
					x = d2          // 也可以将d2赋值给变量x
					x.Move()
	
			2️⃣指针类型接收者
				使用指针接收者实现接口之后，只有结构体指针类型的变量都可以赋值给该接口变量。
			
				举个例子：
					type Cat struct{}
					
					func (c *Cat) Move() {
						fmt.Println("猫会动")
					}
				
					var c1 = &Cat{} // c1是*Cat类型
					x = c1          // 可以将c1当成Mover类型
					x.Move()

					// 下面的代码无法通过编译
					var c2 = Cat{} // c2是Cat类型
					x = c2         // 不能将c2当成Mover类型

		(5).空接口
			1️⃣定义：
				空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。
				也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
				
			2️⃣格式：
				通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}。
				var x interface{}  // 声明一个空接口类型变量x


				举个例子：
					type Dog struct{}

					var x interface{}

					x = "你好" // 字符串型
					fmt.Printf("type:%T value:%v\n", x, x)
					x = 100 // int型
					fmt.Printf("type:%T value:%v\n", x, x)
					x = true // 布尔型
					fmt.Printf("type:%T value:%v\n", x, x)
					x = Dog{} // 结构体类型
					fmt.Printf("type:%T value:%v\n", x, x)

			3️⃣应用：
					场景一：
						// 空接口作为函数参数
						func show(a interface{}) {
							fmt.Printf("type:%T value:%v\n", a, a)
						}
			
					场景二：
						// 空接口作为map值
						var studentInfo = make(map[string]interface{})
						studentInfo["name"] = "沙河娜扎"
						studentInfo["age"] = 18
						studentInfo["married"] = false
						fmt.Println(studentInfo)

		(6).接口值
			1️⃣定义：
			由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
			也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值。

				举个例子：
					type Mover interface {
						Move()
					}
					
					type Dog struct {
						Name string
					}
					
					func (d *Dog) Move() {
						fmt.Println("狗在跑~")
					}
					
					type Car struct {
						Brand string
					}
					
					func (c *Car) Move() {
						fmt.Println("汽车在跑~")
					}

					
					//此时，接口变量m是接口类型的零值，也就是它的类型和值部分都是nil
					var m Mover

					//此时，接口值m的动态类型会被设置为*Dog，动态值为结构体变量的拷贝
					m = &Dog{Name: "旺财"}

					//这一次，接口值m的动态类型为*Car，动态值为nil
					var c *Car
					m = c

		(7).类型断言
			1️⃣语法：
				x.(T)
				其中：
					x：表示接口类型的变量
					T：表示断言x可能是的类型。
				该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，
				若为true则表示断言成功，为false则表示断言失败。

					举个例子：
						var n Mover = &Dog{Name: "旺财"}
						v, ok := n.(*Dog)
						if ok {
							fmt.Println("类型断言成功")
							v.Name = "富贵" // 变量v是*Dog类型
						} else {
							fmt.Println("类型断言失败")
						}
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	//var inZ InterfaceZ
	//inZ = &reverseZ{}
	//
	//inZ1 := ReverseZ(inZ)
	//fmt.Println(inZ1.Less(1, 2))

	/**
	接口介绍
	*/
	fmt.Println(`
		在Go语言中接口（interface）是一种类型，一种抽象的类型。
		相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）更注重“我是谁”，接口类型更注重“我能做什么”的问题。
		接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。
	`)

	/**
	接口类型定义
	*/
	fmt.Println(`
		格式：
			type 接口类型名 interface{
			方法名1( 参数列表1 ) 返回值列表1
			方法名2( 参数列表2 ) 返回值列表2
			…
		}
	
		其中：
			接口类型名：Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有关闭操作的接口叫closer等。接口名最好要能突出该接口的类型含义。
			方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
			参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
	
		举例：
			type Writer interface{
				Write([]byte) error
			}
	`)

	/**
	实现接口的条件
	*/
	fmt.Println(`
		接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口。

		例如：
			type Singer interface {
				Sing()
				Move()
			}
			
			type Bird struct {
			}
			
			func (b Bird) Sing() {
				fmt.Println("夹夹夹")
			}
			
			func (b Bird) Move() {
				fmt.Println("退退退")
			}
	`)

	/**
	接口类型变量
	*/
	var s Singer = &Bird{}
	fmt.Println(s)

	s = Bird{}
	fmt.Println(s)

	/**
	值接收者实现接口
	*/
	var x Mover
	var d1 = Dog{}
	x = d1
	x.MoveZ()

	var d2 = &Dog{}
	x = d2
	d2.MoveZ()
	x.MoveZ()

	fmt.Println(`
		type Mover interface {
			MoveZ()
		}
		
		type Dog struct {
		}
		
		func (d Dog) MoveZ() {
			fmt.Println("走动")
		}

		从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。
	`)

	/**
	指针接收者实现接口
	*/
	fmt.Println(`
		type Mover interface {
			MoveZ()
		}
		
		type Dog struct {
		}
		
		func (d *Dog) MoveZ() {
		
		}

		使用指针接收者实现接口之后，只能接受对应的结构体指针类型的变量赋值给该接口变量。
	`)

	/**
	接口组合
	*/
	fmt.Println(`
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
		
		type Closer interface {
			Close() error
		}
		
		// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
		type ReadWriter interface {
			Reader
			Writer
		}
		
		// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
		type ReadCloser interface {
			Reader
			Closer
		}
		
		// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
		type WriteCloser interface {
			Writer
			Closer
		}

		对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。


		通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。

		例如：

			// LessZ 为InterStruct类型添加LessZ方法，重写原Inter接口类型的LessZ方法
			func (in InterStruct) LessZ(i, j int) bool {
				return in.Inter.LessZ(j, i)
			}

			func InterStruct(data Inter) Inter {
				return &InterStruct{data}
			}
	`)

	/**
	空接口
	空接口是指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。
	也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
	*/
	var az AnyZ
	az = "你好"
	az = 100
	az = true
	fmt.Println("空接口", az)

	/**
	空接口应用
	*/
	fmt.Printf(`
			// 空接口作为函数参数
			func show(a interface{}) {
				fmt.Printf("type:%T value:%v\n", a, a)
			}
			
			//空接口作为map的值
			var studentInfo = make(map[string]interface{})
			studentInfo["name"] = "沙河娜扎"
			studentInfo["age"] = 18
			studentInfo["married"] = false
			fmt.Println(studentInfo)
	
	`)

	/**
	类型断言
	*/
	var inr Inter
	inr = &InterStruct{}

	var inr1 Inter
	inr1 = InterStruct{}

	v, ok := inr.(InterStruct)
	if ok {
		v.Inter = &InterStruct{}
		fmt.Println("类型断言成功")
	} else {
		fmt.Println(v)
		fmt.Println("类型断言失败")
	}

	v1, ok1 := inr1.(InterStruct)
	if ok1 {
		v1.Inter = InterStruct{}
		fmt.Println("类型断言成功")
	} else {
		fmt.Println("类型断言失败")
	}

	var a11 int
	b11 := 10
	a11 = b11
	a11 -= 10
	fmt.Println(a11)
	fmt.Println(b11)

	c11 := &b11
	*c11 = 100
	fmt.Println(b11)

}

type AnyZ interface {
}

type Singer interface {
	Sing()
	Move()
}

type Bird struct {
}

func (b Bird) Sing() {
	fmt.Println("夹夹夹")
}

func (b Bird) Move() {
	fmt.Println("退退退")
}

type Mover interface {
	MoveZ()
}

type Dog struct {
}

//func (d *Dog) MoveZ() {
//
//}

func (d Dog) MoveZ() {
	fmt.Println("走动")
}

type Inter interface {
	LessZ(i, j int) bool
}

type InterStruct struct {
	Inter
}

func (in InterStruct) LessZ(i, j int) bool {
	return in.Inter.LessZ(j, i)
}

// src/sort/sort.go

// InterfaceZ Interface 定义通过索引对元素排序的接口类型
type InterfaceZ interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// reverse 结构体中嵌入了Interface接口
type reverseZ struct {
	InterfaceZ
}

// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverseZ) Less(i, j int) bool {
	return r.InterfaceZ.Less(j, i)
}

func ReverseZ(data InterfaceZ) InterfaceZ {
	return &reverseZ{data}
}
