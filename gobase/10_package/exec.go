package packageZhang

import (
	"fmt"
)

func Exec() {

	fmt.Println(`
	1.【package】
		(1).定义：
			Go语言中支持模块化的开发理念，在Go语言中使用包（package）来支持代码模块化和代码复用。
			一个包是由一个或多个Go源码文件（.go结尾的文件）组成，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等。
		
		(2).格式：
			package packagename

			其中：
				package：声明包的关键字。
				packagename：包名，可以不与文件夹的名称一致，不能包含 - 符号，最好与其实现的功能相对应。

			另外：
				需要注意一个文件夹下面直接包含的文件只能归属一个包，同一个包的文件不能在多个文件夹下。
				包名为main的包是应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。


		(3).可见性：
			在同一个包内部声明的标识符（如变量、常量、类型、函数等）都位于同一个命名空间下，在不同的包内部声明的标识符就属于不同的命名空间。
			想要在包的外部使用其它包的标识符就需要添加包名前缀，例如fmt.Println("Hello world!")，就是指调用fmt包中的Println函数。

			如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，那么标识符必须是对外可见的（public）。
			在Go语言中，只有首字母为大写的标识符才是导出的（Exported），才能对包外的代码可见；如果首字母是小写的，那么就说明这个标识符仅限于在声明它的包内可见。。

			举个例子：
				package demo
				import "fmt"
			
				// num 定义一个全局整型变量
				// 首字母小写，对外不可见(只能在当前包内使用)
				var num = 100
				
				// Mode 定义一个常量
				// 首字母大写，对外可见(可在其它包中使用)
				const Mode = 1
				
				// person 定义一个代表人的结构体
				// 首字母小写，对外不可见(只能在当前包内使用)
				type person struct {
					name string
					Age  int
				}
				
				// Add 返回两个整数和的函数
				// 首字母大写，对外可见(可在其它包中使用)
				func Add(x, y int) int {
					return x + y
				}
				
				// sayHi 打招呼的函数
				// 首字母小写，对外不可见(只能在当前包内使用)
				func sayHi() {
					var myName = "zhangp" // 函数局部变量，只能在当前函数内使用
					fmt.Println(myName)
				}
		
				//同样的规则也适用于结构体，结构体中可导出字段的字段名称必须首字母大写。
				type Student struct {
					Name  string // 可在包外访问的方法
					class string // 仅限包内访问的字段
				}

		(4).包的引入：
			1️⃣定义：要在当前包中使用另外一个包的内容就需要使用import关键字引入这个包，并且import语句通常放在文件的开头，package声明语句的下方。
			
				格式：
					import importname "path/to/package"
	
				其中：
					1.importname：引入的包名，通常都省略。默认值为引入包的包名。
					2.path/to/package：引入包的路径名称，必须使用双引号包裹起来。
					3.Go语言中禁止循环导入包。
	
				举个例子：
					import "fmt"
					import "net/http"
					import "os"

					fmt.Println("Hello world!")
					
					或：
					import (
						"fmt"
						"net/http"
						"os"
					)
				
					fmt.Println("Hello world!")

				注意⚠️：
					1.import “fmt” 一行中“fmt”代表的是包的导入路径，它表示的是标准库下的fmt目录，整个import声明语句的含义是导入标准库fmt目录下的包；
					2.fmt.Println函数调用一行中的“fmt”代表的则是包名。
					通常导入路径的最后一个【分段名】与【包名】是相同的，这也很容易让人误解import声明语句中的“fmt”指的是包名，其实并不是这样的。

			2️⃣包名：
				当引入的多个包中存在相同的包名或者想自行为某个引入的包设置一个新包名时，都需要通过importname指定一个在当前文件中使用的新包名。
				
				举个例子：
					import f "fmt"
					f.Println("Hello world!")

			3️⃣匿名：
				如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。
				一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。 
				被匿名引入的包中的init函数将被执行并且仅执行一遍。
				匿名引入的包与其他方式导入的包一样都会被编译到可执行文件中。

				举个例子：
					import _ "github.com/go-sql-driver/mysql"

			4️⃣注意：
				1.Go语言中不允许引入包却不在代码中使用这个包的内容，如果引入了未使用的包则会触发编译错误。

				2.在Go语言中，main包是不可以像标准库fmt包那样被导入（Import）的，如果导入main包，
				在代码编译阶段你会收到一个Go编译器错误：import “xx/main” is a program, not an importable package。

		(5).init函数：
			在每一个Go源文件中，都可以定义任意个如下格式的特殊函数。
			这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。
			当程序启动的时候，init函数会按照它们声明的顺序自动执行。

			格式：
				func init(){
				  // ...
				}
	`)

	fmt.Println(`
	2.【go module】
		(1).定义：
			在Go语言的早期版本中，我们编写Go项目代码时所依赖的所有第三方包都需要保存在GOPATH这个目录下面。
			Go module 会把下载到本地的依赖包会以类似下面的形式保存在 $GOPATH/pkg/mod目录下，
			每个依赖包都会带有版本号进行区分，这样就允许在本地存在同一个包的多个不同版本。
			这样的依赖管理方式存在一个致命的缺陷，那就是不支持版本管理，同一个依赖包只能存在一个版本的代码。
			可是我们本地的多个项目完全可能分别依赖同一个第三方包的不同版本。


		(2).背景
			GOPATH：
				GOPATH 是 Go语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录。
				Go语言所依赖的所有的第三方库都放在 GOPATH 这个目录下面，
				这就导致了同一个库只能保存一个版本的代码。
				
				查看GOPATH目录：
				go env
				
				Go Path代码组织形式:
				bin：存放编译后生成的二进制可执行文件
				pkg：存放编译后生成的文件
				src：存放项目的源代码，可以是你自己写的代码，也可以是你 go get 下载的包
				将你的包或者别人的包全部放在 $GOPATH/src 目录下进行管理的方式，我们称之为 GOPATH 模式。

			GO vendor 模式的过渡：
				为了解决 GOPATH 方案下不同项目下无法使用多个版本库的问题，
				Go v1.5 开始支持 vendor。
 

			GO Mod 模式：
				go module 是Go语言从 1.11 版本之后官方推出的版本管理工具，
				并且从 Go1.13 版本开始，go module 成为了Go语言默认的依赖管理工具。
				不需要配置 GOPATH 代码可以放到任意目录
				1.16 默认开启GO111MODULE  官方推荐使用。

		(3).命令：
			go mod init			初始化项目依赖，生成go.mod文件
			go mod download		根据go.mod文件下载依赖
			go mod tidy			比对项目文件中引入的依赖与go.mod进行比对
			go mod graph		输出依赖关系图
			go mod edit			编辑go.mod文件
			go mod vendor		将项目的所有依赖导出至vendor目录
			go mod verify		检验一个依赖包是否被篡改过
			go mod why			解释为什么需要某个依赖

		(4).使用：
			1️⃣引入包
				$ mkdir holiday
				$ cd holiday

				第一步：初始化工程
					$ go mod init holiday
					go: creating new go.mod: module holiday
				第二步：下载依赖包
					方法一：
						holiday $ go get -u github.com/zhangp/hello
						go get: added github.com/zhangp/hello v0.1.1

						此时，我们打开go.mod文件就可以看到下载的依赖包及版本信息都已经被记录下来了。
						内容如下：
							module holiday
	
							go 1.16
							
							require github.com/zhangp/hello v0.1.1 // indirect

					方法二：
						直接编辑go.mod文件，将依赖包和版本信息写入该文件。
						然后，执行go mod download命令。

			2️⃣发布包
				一个设计完善的包应该包含开源许可证及文档等内容，并且我们还应该尽心维护并适时发布适当的版本。
				github 上发布版本号使用git tag为代码包打上标签即可。

				举个例子：
					hello $ git tag -a v0.1.0 -m "release version v0.1.0"
					hello $ git push origin v0.1.0

				经过上面的操作我们就发布了一个版本号为v0.1.0的版本。

			3️⃣废弃包
				如果某个发布的版本存在致命缺陷不再想让用户使用时，我们可以使用retract声明废弃的版本。

				举个例子：（我们在hello/go.mod文件中按如下方式声明即可对外废弃v0.1.2版本。）
					module github.com/zhangp/hello

					go 1.16
					
					retract v0.1.2
				
				用户使用go get下载v0.1.2版本时就会收到提示，催促其升级到其他版本。

	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	包介绍
	*/
	fmt.Println(`
		在工程化的Go语言开发项目中，Go语言的源码复用是建立在包（package）基础之上的。

		Go语言中支持模块化的开发理念，在Go语言中使用包（package）来支持代码模块化和代码复用。
		一个包是由一个或多个Go源码文件（.go结尾的文件）组成，是一种高级的代码复用方案，
		Go语言为我们提供了很多内置包，如fmt、os、io等。
	`)

	/**
	包定义
	格式：【package packagename】
	*/
	fmt.Println(`
		package：声明包的关键字
		packagename：包名，可以不与文件夹的名称一致，不能包含 - 符号，最好与其实现的功能相对应。
	`)

	/**
	标识符可见性
	*/
	fmt.Println(`
		在同一个包内部声明的标识符都位于同一个命名空间下，在不同的包内部声明的标识符就属于不同的命名空间。
		想要在包的外部使用包内部的标识符就需要添加包名前缀，例如fmt.Println("Hello world!")，就是指调用fmt包中的Println函数。

		如果想让一个包中的标识符（如变量、常量、类型、函数等）能被外部的包使用，那么标识符必须是对外可见的（public）。在Go语言中是通过标识符的首字母大/小写来控制标识符的对外可见（public）/不可见（private）的。在一个包内部只有首字母大写的标识符才是对外可见的。
		例如我们定义一个名为demo的包，在其中定义了若干标识符。在另外一个包中并不是所有的标识符都能通过demo.前缀访问到，因为只有那些首字母是大写的标识符才是对外可见的。
	`)
	//例如调用structZhang包下的Exec方法
	//structZhang.Exec()

	/**
	包的引入

	要在当前包中使用另外一个包的内容就需要使用import关键字引入这个包，
	并且import语句通常放在文件的开头，package声明语句的下方。
	完整的引入声明语句格式如下:
	【import importname "path/to/package"】

	其中：
		1.importname：引入的包名，通常都省略。默认值为引入包的包名。
		2.path/to/package：引入包的路径名称，必须使用双引号包裹起来。
		3.Go语言中禁止循环导入包。
	*/
	fmt.Println(`
		例如：
			import "fmt"
			import "net/http"
			import "os"

		也可以这么写：
			import (
				"fmt"
				"net/http"
				"os"
			)

		当引入的多个包中存在相同的包名或者想自行为某个引入的包设置一个新包名时，
		都需要通过importname指定一个在当前文件中使用的新包名。
		例如，在引入fmt包时为其指定一个新包名f。
		import f "fmt"

		这样在当前这个文件中就可以通过使用f来调用fmt包中的函数了。
		f.Println("Hello world!")
	`)

	/**
	包的匿名引入
	*/
	fmt.Println(`
		如果引入一个包的时候为其设置了一个特殊_作为包名，那么这个包的引入方式就称为匿名引入。
		一个包被匿名引入的目的主要是为了加载这个包，从而使得这个包中的资源得以初始化。 
		被匿名引入的包中的init函数将被执行并且仅执行一遍。
		例如：import _ "github.com/go-sql-driver/mysql"


		匿名引入的包与其他方式导入的包一样都会被编译到可执行文件中。
		需要注意的是，Go语言中不允许引入包却不在代码中使用这个包的内容，
		如果引入了未使用的包则会触发编译错误。
	`)

	/**
	init函数初始化

	在每一个Go源文件中，都可以定义任意个如下格式的特殊函数：

	格式：
	func init(){
	  // ...
	}

	这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。
	当程序启动的时候，init函数会按照它们声明的顺序自动执行。

	一个包的初始化过程是按照代码中引入的顺序来进行的，所有在该包中声明的init函数都将被串行调用并且仅调用执行一次。
	每一个包初始化的时候都是先执行依赖的包中声明的init函数再执行当前包中声明的init函数。
	确保在程序的main函数开始执行时所有的依赖包都已初始化完成。
	*/
	fmt.Println(`
		func init() {
			fmt.Println("init A")
		}
		
		func init() {
			fmt.Println("init B")
		}
	`)

	/**
	go module

	在Go语言的早期版本中，我们编写Go项目代码时所依赖的所有第三方包都需要保存在GOPATH这个目录下面。
	这样的依赖管理方式存在一个致命的缺陷，那就是不支持版本管理，同一个依赖包只能存在一个版本的代码。
	可是我们本地的多个项目完全可能分别依赖同一个第三方包的不同版本。
	*/
	fmt.Println(`
		Go module 是 Go1.11 版本发布的依赖管理方案，从 Go1.14 版本开始推荐在生产环境使用，
		于Go1.16版本默认开启。Go module 提供了以下命令供我们使用：

		
		go mod init		初始化项目依赖，生成go.mod文件
		go mod download	根据go.mod文件下载依赖
		go mod tidy		比对项目文件中引入的依赖与go.mod进行比对
		go mod graph	输出依赖关系图
		go mod edit		编辑go.mod文件
		go mod vendor	将项目的所有依赖导出至vendor目录
		go mod verify	检验一个依赖包是否被篡改过
		go mod why		解释为什么需要某个依赖

		Go语言在 go module 的过渡阶段提供了 GO111MODULE 这个环境变量来作为是否启用 go module 功能的开关，
		考虑到 Go1.16 之后 go module 已经默认开启，所以本书不再介绍该配置，对于刚接触Go语言的读者而言完全没有必要了解这个历史包袱。
	`)

	/**
	GOPRIVATE

	GOPRIVATE 的值也可以设置多个，多个地址之间使用英文逗号 “,” 分隔。
	我们通常会把自己公司内部的代码仓库设置到 GOPRIVATE 中，
	例如：
		$ go env -w GOPRIVATE="git.mycompany.com"
	*/
	fmt.Println(`
		设置了GOPROXY 之后，go 命令就会从配置的代理地址拉取和校验依赖包。
		当我们在项目中引入了非公开的包（公司内部git仓库或 github 私有仓库等），
		此时便无法正常从代理拉取到这些非公开的依赖包，这个时候就需要配置 GOPRIVATE 环境变量。
		GOPRIVATE用来告诉 go 命令哪些仓库属于私有仓库，不必通过代理服务器拉取和校验。
	`)

	/**
	go module引入包
	*/

	fmt.Println(`
		我们的pkg1项目现在需要引入一个第三方包github.com/zp/hello来实现一些必要的功能。
		类似这样的场景在我们的日常开发中是很常见的。
		我们需要先将依赖包下载到本地同时在go.mod中记录依赖信息，然后才能在我们的代码中引入并使用这个包。
		下载依赖包主要有两种方法。

		方法一：（-u参数，强制更新现有依赖）
			cd 10_package/pkg1
			go get -u github.com/zp/hello

		这样默认会下载最新的发布版本，你也可以指定想要下载指定的版本号的。
			go get -u github.com/zp/hello@v0.1.0

		如果依赖包没有发布任何版本则会拉取最新的提交，
		最终go.mod中的依赖信息会变成类似下面这种由默认v0.0.0的版本号和最新一次commit的时间和hash组成的版本格式：
		require github.com/zp/hello v0.0.0-20210218074646-139b0bcd549d

		如果想指定下载某个commit对应的代码，可以直接指定commit hash，不过没有必要写出完整的commit hash，一般前7位即可。
		例如：
			go get github.com/zp/hello@2ccfadd

		此时，我们打开go.mod文件就可以看到下载的依赖包及版本信息都已经被记录下来了。
		module pkg1

		go 1.22.3
		
		require github.com/zp/hello v0.1.0  indirect

		行尾的indirect表示该依赖包为间接依赖，说明在当前程序中的所有 import 语句中没有发现引入这个包。

		
		方法二：
			1.cd 10_package/pkg1
			2.直接编辑go.mod文件，将依赖包和版本信息写入该文件，如下：
				module pkg1
	
				go 1.22.3
				
				require github.com/zp/hello latest   //latest  表示当前项目需要使用github.com/zp/hello库的最新版本

			3.在项目目录下执行go mod download下载依赖包
			
		如果不输出其它提示信息就说明依赖已经下载成功，此时go.mod文件已经变成如下内容：
			module pkg1
	
			go 1.22.3
			
			require github.com/zp/hello v0.1.1  //从中我们可以知道最新的版本号是v0.1.1。如果事先知道依赖包的具体版本号，可以直接在go.mod中指定需要的版本然后再执行go mod download下载。

		最终在文件中使用引入的包，如下：

			import (
				"fmt"
				"github.com/zp/hello"
			)
			
			func main() {
				hello.SayHi() // 调用hello包的SayHi函数
			}


		当我们的项目功能越做越多，代码越来越多的时候，通常会选择在项目内部按功能或业务划分成多个不同包。
		Go语言支持在一个项目（project <==> module）下定义多个包（package）。
		

		声明依赖的格式如下：
			require module/path v1.2.3

		其中：
			require：声明依赖的关键字
			module/path：依赖包的引入路径
			v1.2.3：依赖包的版本号。支持以下几种格式：
				latest：最新版本
				v1.0.0：详细版本号
				commit hash：指定某次commit hash


		依赖保存的位置：
			Go module 会把下载到本地的依赖包会以类似下面的形式保存在 $GOPATH/pkg/mod目录下，
			每个依赖包都会带有版本号进行区分，这样就允许在本地存在同一个包的多个不同版本。

			mod
			├── cache
			├── cloud.google.com
			├── github.com
					└──zp
					  ├── hello@v0.0.0-20210218074646-139b0bcd549d
					  ├── hello@v0.1.1
					  └── hello@v0.1.0
			...

		如果想清除所有本地已缓存的依赖包数据，可以执行 go clean -modcache 命令。
	`)

	/**
	使用go module发布包
	*/
	fmt.Println(`
		当我们想要在社区发布一个自己编写的代码包或者在公司内部编写一个供内部使用的公用组件时，我们该怎么做呢？

		我们需要将我们编写的代码包发布到github.com仓库，就能够被全球的Go语言开发者使用。
		比如：我们将该项目的代码 push 到仓库的远端分支，这样就对外发布了一个Go包。
		一个设计完善的包应该包含开源许可证及文档等内容，并且我们还应该尽心维护并适时发布适当的版本。
		github 上发布版本号使用git tag为代码包打上标签即可。

		例如，发布一个版本号为v0.1.0的版本的包：
		git tag -a v0.1.0 -m "release version v0.1.0"
		git push origin v0.1.0
	
		其中：
			主版本号：发布了不兼容的版本迭代时递增（breaking changes）。  	第一位版本号
			次版本号：发布了功能性更新时递增。							第二位版本号
			修订号：发布了bug修复类更新时递增。							第三位版本号

		当要发布一个改动很大的版本时，比如要升级主版本号，在这种情况下，我们通常会修改当前包的引入路径，像下面的示例一样为引入路径添加版本后缀。

		//pkg1/go.mod

		module github.com/zp/hello/v2

	
		go 1.22.3

		
		打好 tag 推送到远程仓库：
		git tag -a v2.0.0 -m "release version v2.0.0"
		git push origin v2.0.0


		这样在不影响使用旧版本的用户的前提下，我们新的版本也发布出去了。想要使用v2版本的代码包的用户只需按修改后的引入路径下载即可。
		如：
			go get github.com/zp/hello/v2@v2.0.0
		
		在代码中使用的过程与之前类似，只是需要注意引入路径要添加 v2 版本后缀。

		package main

		import (
			"fmt"
		
			"github.com/zp/hello/v2" // 引入v2版本
		)
		
		func main() {
			fmt.Println("现在是假期时间...")
		
			hello.SayHi("张三") // v2版本的SayHi函数需要传入字符串参数
		}
	`)

	/**
	废弃已发布版本
	*/
	fmt.Println(`
		如果某个发布的版本存在致命缺陷不再想让用户使用时，我们可以使用retract声明废弃的版本。例如我们在hello/go.mod文件中按如下方式声明即可对外废弃v0.1.2版本。

		module github.com/zp/hello

	
		go 1.22.3

		retract v0.1.2
	`)
}

// 包级别标识符的可见性

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 100

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 1

// person 定义一个代表人的结构体
// 首字母小写，对外不可见(只能在当前包内使用)
type person struct {
	name string
	Age  int
}

// Add 返回两个整数和的函数
// 首字母大写，对外可见(可在其它包中使用)
func Add(x, y int) int {
	return x + y
}

// sayHi 打招呼的函数
// 首字母小写，对外不可见(只能在当前包内使用)
func sayHi() {
	var myName = "张三丰"
	fmt.Println(myName)
}

// Student 同样的规则也适用于结构体中的字段，
// 结构体中可导出字段的字段名称必须首字母大写。
type Student struct {
	Name string // 可在包外访问的方法
	id   int    // 仅限包内访问的字段
}

func init() {
	fmt.Println("init A")
}

func init() {
	fmt.Println("init B")
}
