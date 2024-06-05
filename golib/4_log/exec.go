package logZhang

import (
	"fmt"
	"log"
)

func Exec() {
	fmt.Printf(`
	1.【log】
		(1).使用：
			Go语言内置的log包实现了简单的日志服务。

			举个例子：
				log.Println("这是一条很普通的日志。")
				v := "很普通的"
				log.Printf("这是一条%s日志。\n", v)
				log.Fatalln("这是一条会触发fatal的日志。")
				log.Panicln("这是一条会触发panic的日志。")
			
			输出：
				2024/06/04 14:04:17 这是一条很普通的日志。
				2024/06/04 14:04:17 这是一条很普通的日志。
				2024/06/04 14:04:17 这是一条会触发fatal的日志。

		(2).配置：
			1️⃣背景
				默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。
				log标准库中为我们提供了定制这些设置的方法。
				log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。
	
				func Flags() int
				func SetFlags(flag int)
			
			2️⃣选项
				log标准库提供了如下的flag选项，它们是一系列定义好的常量。

				const (
					// 控制输出日志信息的细节，不能控制输出的顺序和格式。
					// 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
					Ldate         = 1 << iota     // 日期：2009/01/23
					Ltime                         // 时间：01:23:23
					Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
					Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
					Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
					LUTC                          // 使用UTC时间
					LstdFlags     = Ldate | Ltime // 标准logger的初始值
				)
		
				举个例子：
					func main() {
						log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
						log.Println("这是一条很普通的日志。")
					}

				输出：
					2024/06/04 14:05:17.494943 .../log_demo/main.go:11: 这是一条很普通的日志。

			3️⃣前缀
				log标准库中还提供了关于日志信息前缀的两个方法Prefix()和SetPrefix().
				这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。
				
				func Prefix() string
				func SetPrefix(prefix string)

				举个例子：
					func main() {
						log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
						log.Println("这是一条很普通的日志。")
						log.SetPrefix("[项羽]")
						log.Println("这是一条很普通的日志。")
					}

				输出：
					[项羽]2024/06/04 14:05:57.940542 .../log_demo/main.go:13: 这是一条很普通的日志。

			4️⃣输出位置
				1.设置
					SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
					func SetOutput(w io.Writer)
	
						举个例子：
							func main() {
								logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
								if err != nil {
									fmt.Println("open log file failed, err:", err)
									return
								}
								log.SetOutput(logFile)
								log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
								log.Println("这是一条很普通的日志。")
								log.SetPrefix("[小王子]")
								log.Println("这是一条很普通的日志。")
							}
						
						以上例子会把日志输出到同目录下的xx.log文件中。
			
			5️⃣标准写法
				如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。

				举个例子：
					func init() {
						logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
						if err != nil {
							fmt.Println("open log file failed, err:", err)
							return
						}
						log.SetOutput(logFile)
						log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
					}

		(3).创建：
			log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。

			func New(out io.Writer, prefix string, flag int) *Logger

			举个例子：
				func main() {
					logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
					logger.Println("这是自定义的logger记录的日志。")
				}

			输出：
				<New>2024/06/04 14:06:51 main.go:34: 这是自定义的logger记录的日志。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
