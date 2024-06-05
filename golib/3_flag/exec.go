package flagZhang

import (
	"flag"
	"fmt"
	"time"
)

func Exec() {
	fmt.Printf(`
	1.【flag】
		(1).定义：
			Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。

		(2).os.Args:
			如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用os.Args来获取命令行参数。
			os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。

			举个例子：
				//os.Args是一个[]string
				if len(os.Args) > 0 {
					for index, arg := range os.Args {
						fmt.Printf("args[%d]=%v\n", index, arg)
					}
				}

			编译：
				go build -o "args_demo"

			执行：
				./args_demo a b c d

			输出：
				args[0]=./args_demo
				args[1]=a
				args[2]=b
				args[3]=c
				args[4]=d

		(3).定义命令行：
			1️⃣flag.Type()
				格式：
					flag.Type(flag名, 默认值, 帮助信息)*Type
	
				举个例子：
					name := flag.String("name", "张三", "姓名")
					age := flag.Int("age", 18, "年龄")
					married := flag.Bool("married", false, "婚否")
					delay := flag.Duration("d", 0, "时间间隔")
	
				需要注意的是，此时name、age、married、delay均为对应类型的指针。

			2️⃣flag.TypeVar()
				格式：
					flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
			
				举个例子：
					var name string
					var age int
					var married bool
					var delay time.Duration
					flag.StringVar(&name, "name", "张三", "姓名")
					flag.IntVar(&age, "age", 18, "年龄")
					flag.BoolVar(&married, "married", false, "婚否")
					flag.DurationVar(&delay, "d", 0, "时间间隔")

		(4).解析命令行：
				通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
				
				支持的命令行参数格式有以下几种：

					-flag xxx （使用空格，一个-符号）
					--flag xxx （使用空格，两个-符号）
					-flag=xxx （使用等号，一个-符号）
					--flag=xxx （使用等号，两个-符号）
					其中，布尔类型的参数必须使用等号的方式指定。
					
					Flag解析在第一个非flag参数（单个"-“不是flag参数）之前停止，或者在终止符”–“之后停止。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	//if len(os.Args) > 0 {
	//	for index, arg := range os.Args {
	//		fmt.Printf("args[%d]=%v\n", index, arg)
	//	}
	//}

	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")
	//
	//fmt.Printf("%v\n", *name)
	//fmt.Printf("%v\n", *age)
	//fmt.Printf("%v\n", *married)
	//fmt.Printf("%v\n", *delay)

	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "时间间隔")

	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}
