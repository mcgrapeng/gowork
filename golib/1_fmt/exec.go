package fmtZhang

import "fmt"

func Exec() {

	fmt.Printf(`
	1.【fmt】
		(1).Print：
			Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，
			Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
			
			func Print(a ...interface{}) (n int, err error)
			func Printf(format string, a ...interface{}) (n int, err error)
			func Println(a ...interface{}) (n int, err error)
		
			举个例子：
				func main() {
					fmt.Print("在终端打印该信息。")
					name := "zhangp"
					fmt.Printf("我是：%s\n", name)
					fmt.Println("在终端打印单独一行显示")
				}
			输出：
				在终端打印该信息。我是：zhangp
				在终端打印单独一行显示

		(2).Fprint：
			Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。

			func Fprint(w io.Writer, a ...interface{}) (n int, err error)
			func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
			func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

			举个例子：
				// 向标准输出写入内容
				fmt.Fprintln(os.Stdout, "向标准输出写入内容")

				fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("打开文件出错，err:", err)
					return
				}
				name := "zhangp"
				// 向打开的文件句柄中写入内容
				fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

		(3).Sprint：
			Sprint系列函数会把传入的数据生成并返回一个字符串。
			
			func Sprint(a ...interface{}) string
			func Sprintf(format string, a ...interface{}) string
			func Sprintln(a ...interface{}) string

			举个例子：
				s1 := fmt.Sprint("zhangp")
				name := "zhangp"
				age := 18
				s2 := fmt.Sprintf("name:%s,age:%d", name, age)
				s3 := fmt.Sprintln("zhangp")
				fmt.Println(s1, s2, s3)

		(4).Errorf：
			Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
			
			func Errorf(format string, a ...interface{}) error

			Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
			举个例子：
				e := errors.New("原始错误e")
				w := fmt.Errorf("Wrap了一个错误%w", e)
	`)

	fmt.Printf(`
	2.【格式化占位符】
		(1).定义：
			*printf系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

		(2).通用占位符：
			占位符	说明
			%v		值的默认格式表示
			%+v		类似%v，但输出结构体时会添加字段名
			%#v		值的Go语法表示
			%T	打印值的类型
			%%	百分号
		
			举个例子：
				fmt.Printf("%v\n", 100)
				fmt.Printf("%v\n", false)
				o := struct{ name string }{"小王子"}
				fmt.Printf("%v\n", o)
				fmt.Printf("%#v\n", o)
				fmt.Printf("%T\n", o)
				fmt.Printf("100%%\n")
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	s1 := fmt.Sprint("zhangp")
	name := "zhangp"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("zhangp")
	fmt.Println(s1, s2, s3)

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
