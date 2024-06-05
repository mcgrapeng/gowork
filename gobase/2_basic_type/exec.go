package basic_type

import (
	"fmt"
	"math"
)

func Exec() {
	/**
	整型
	*/
	fmt.Println(`
	1.【整型】
		(1).整型分类：
				按长度分为：int8、int16、int32、int64。对应的无符号整型：uint8、uint16、uint32、uint64
				其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

		(2).数值范围：
				uint8	无符号 8位整型 (0 到 255)
				uint16	无符号 16位整型 (0 到 65535)
				uint32	无符号 32位整型 (0 到 4294967295)
				uint64	无符号 64位整型 (0 到 18446744073709551615)
				int8	有符号 8位整型 (-128 到 127)
				int16	有符号 16位整型 (-32768 到 32767)
				int32	有符号 32位整型 (-2147483648 到 2147483647)
				int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)
		
		(3).特殊整型：
				uint	32位操作系统上就是uint32，64位操作系统上就是uint64
				int		32位操作系统上就是int32，64位操作系统上就是int64
				uintptr	无符号整型，用于存放一个指针

		**注意事项**：
			实际使用中，切片或 map 的元素数量等都可以用int来表示。
			在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint


		Go 语言的字符有以下两种：
			uint8类型，或者叫 byte 型，代表一个ASCII码字符。
			rune类型，代表一个 UTF-8字符。 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32
	`)

	/**
	浮点类型
	*/
	fmt.Println(`
	2.【浮点类型】
		(1).浮点型分类：float32和float64
	`)

	/**
	复数类型
	*/
	fmt.Println(`
	3.【复数类型】
		(1).复数型分类：complex64和complex128。
		(2).复数结构：复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	`)

	/**
	布尔类型
	*/
	fmt.Println(`
	4.【布尔类型】
		(1).布尔型分类：Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
		**注意事项**：
			1️⃣布尔类型变量的默认值为false。
			2️⃣Go 语言中不允许将整型强制转换为布尔型.
			3️⃣布尔型无法参与数值运算，也无法与其他类型进行转换。
	`)

	/**
	字符串类型
	*/
	fmt.Println(`
	5.【字符串类型】
		(1).定义：
			Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
			Go语言里的字符串的内部实现使用UTF-8编码。
			字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符。
		举个例子：
			s1 := "hello"
			s2 := "你好"
		(2).转义符：
			\r	回车符（返回行首）
			\n	换行符（直接跳到下一行的同列位置）
			\t	制表符
			\'	单引号
			\"	双引号
			\\	反斜杠
		(3).字符串常用操作：
			len(str)	求长度
			+或fmt.Sprintf	拼接字符串
			strings.Split	分割
			strings.contains	判断是否包含
			strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
			strings.Index(),strings.LastIndex()	子串出现的位置
			strings.Join(a[]string, sep string)	join操作
	`)

	/**
	byte和rune类型
	*/
	fmt.Printf(`
	6.【字符类型】
		(1).字符型分类：
			1️⃣uint8类型，或者叫 byte 型，代表一个ASCII码字符（主要用于显示英文）
			2️⃣rune类型，代表一个 UTF-8字符
		(2).应用：
			当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
		
		举个例子：
			func traversalString() {
				s := "hello沙河"
				for i := 0; i < len(s); i++ { //byte
					fmt.Printf("%v(%c) ", s[i], s[i])
				}
				fmt.Println()
				for _, r := range s { //rune
					fmt.Printf("%v(%c) ", r, r)
				}
				fmt.Println()
			}

		输出：
			104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153() 230(æ) 178(²) 179(³) 
			104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)

		原因：
			字符串是不能修改的，字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
			rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	`)

	/**
	类型转换
	*/
	fmt.Println(`
	7.【类型转换】
		(1).定义：
			Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
		(2).格式：
			T(表达式)，其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
		
		举个例子：
			func sqrtDemo() {
				var a, b = 3, 4
				var c int
				// math.Sqrt()接收的参数是float64类型，需要强制转换
				c = int(math.Sqrt(float64(a*a + b*b)))
				fmt.Println(c)
			}
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	/**
	字符串
	*/
	s1 := "hello"
	s2 := "world"
	fmt.Println(s1, s2)

	s3 := `
      	长风破浪会有时，直挂云帆济沧海。
	`
	fmt.Println(s3)
	fmt.Println(len(s3))

	//字符串拼接
	sprint := fmt.Sprint("a", "b")
	fmt.Println(sprint)

	/**
	类型转换
	*/
	var b float32 = 1.2
	var c int32 = int32(b)

	fmt.Println(c)
}

func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
