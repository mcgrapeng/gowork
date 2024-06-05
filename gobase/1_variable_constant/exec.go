package variable_constant

import "fmt"

func Exec() {
	/**
	关键字
	*/
	fmt.Println(`
	1.【Go语言中有25个关键字】
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var`)

	/**
	保留字
	*/
	fmt.Println(`
	2.【Go语言中还有37个保留字】		
		Constants:	true  false  iota  nil
		Types:	int  int8  int16  int32  int64 uint  uint8  uint16  uint32  uint64  uintptr float32  float64  complex128  complex64 bool  byte  rune  string  error
		Functions:	make  len  cap  new  append copy  close delete  complex real  imag  panic  recover`)

	/**
	变量
	*/
	fmt.Println(`
	3.【变量】
		(1).变量声明种类：标准声明、批量声明
		(2).标准声明格式：var 变量名 变量类型
		(3).批量声明格式：var(
					变量名1 变量类型
					变量名2 变量类型
					变量名3 变量类型
				 )
		
			1️⃣变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号。 
			举个例子：
				var name string
				var age int
				var isOk bool
	
			2️⃣由于每声明一个变量就需要写var关键字会比较繁琐，go语言中还支持批量变量声明。
			举个例子：
				var (
					a1 string
					b1 int
					c1 bool
					d1 float32
				)

		(4).变量的初始化：Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。
				每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 
				字符串变量的默认值为空字符串。 布尔型变量默认为false。 
				切片、函数、指针变量的默认为nil。
		(5).变量初始化格式：
			1️⃣【var 变量 变量类型 = 表达式】
			2️⃣【var 变量 = 表达式】
			3️⃣【变量 := 表达式】（:=不能使用在函数外）
			举个例子：
				var name1 string = "zhangp"
				var age1 int = 18
				var name, age = "zhangp", 20
		(6).匿名变量：_多用于占位，表示忽略值。匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。
			举个例子：
				x, _ := foo(1, 2)
				func foo(a int, b int) (int, string) {
					return 0, "zhangp"
				}
		**注意**：
			1️⃣函数外的每个语句都必须以关键字开始（var、const、func等）。
			2️⃣:=不能使用在函数外。
			3️⃣_多用于占位，表示忽略值。
	`)

	/**
	常量
	*/
	fmt.Println(`
	4.【常量】
		(1).常量声明种类：标准声明、批量声明
		(2).标准声明格式：const 变量名 变量类型
		(3).批量声明格式：const(
					常量名1 常量类型
					常量名2 常量类型
					常量名3 常量类型
				 )
		举个例子：
			const name = 3.14
			const e = 1.2
			
			const (
				PI = 3.14
				A  = 2
			)
		(4).常量特殊声明格式：const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。
		举个例子：
			const (
				N = 3.14
				D
				E
			)
	`)

	/**
	iota
	*/
	fmt.Println(`
	5.【iota】
		(1).定义：iota是go语言的常量计数器，只能在常量的表达式中使用。
		(2).应用：iota在const关键字出现时将被重置为0。
			const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 
			使用iota能简化定义，在定义枚举时很有用。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	//变量声明
	var name111 string
	var age111 int
	var isOk bool

	var (
		a1 string
		b1 int
		c1 bool
		d1 float32
	)

	fmt.Println(name111, age111, isOk, a1, b1, c1, d1)

	//变量初始化
	var name11 string = "zhangp"
	var age1 int = 18
	var name12, age2 = "zhangp", 20
	fmt.Println(name11, name12, age2, age1)

	//匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	//常量声明
	const name = "zhangp"
	fmt.Println(name)
	const (
		cname string = "zhang"
		age   int    = 18
	)

	//iota
	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)

}

func foo() (int, string) {
	return 10, "zhangp"
}
