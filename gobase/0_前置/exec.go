package priorZhang

import "fmt"

func Exec() {

	fmt.Println(`
	1.【Go项目的布局结构】

		(1).逻辑结构：
			1️⃣关系：
				项目/仓库 --> module --> package -->源文件(.go文件)
			
			2️⃣说明：
				module是一组同属于一个版本管理单元的包的集合，
				Go支持在一个项目/仓库中存在多个module，
				一个package是由一个或多个Go源码文件（.go结尾的文件）组成，是一种高级的代码复用方案。

		(2).物理结构：
			1️⃣可执行项目布局：
				(1).一个项目中只有一个module

					图解：
						exe-layout
						├── cmd/
						│   ├── app1/
						│   │   └── main.go
						│   └── app2/
						│       └── main.go
						├── go.mod
						├── go.sum
						├── internal/
						│   ├── pkga/
						│   │   └── pkg_a.go
						│   └── pkgb/
						│       └── pkg_b.go
						├── pkg1/
						│   └── pkg1.go
						├── pkg2/
						│   └── pkg2.go
						└── vendor/

					说明：
						【cmd目录】就是存放项目要编译构建的可执行文件对应的main包的源文件。
						如果你的项目中有多个可执行文件需要构建，每个可执行文件的main包单独放在一个子目录中，
						比如图中的app1、app2，cmd目录下的各app的main包将整个项目的依赖连接在一起。
						另外，也有一些Go项目将cmd这个名字改为app或其他名字，但它的功能其实并没有变。
		
						【pkgN目录】是一个存放项目自身要使用、同样也是可执行文件对应main包所要依赖的库文件，
						同时这些目录下的包还可以被外部项目引用。
		
						【internal目录】一个Go项目里的internal目录下的Go包，只可以被本项目内部的包导入。
						项目外部是无法导入这个internal目录下面的包的。

					简化：
						如果Go可执行程序项目有一个且只有一个可执行程序要构建，我们可以将上面项目布局进行简化。

						single-exe-layout
						├── go.mod
						├── internal/
						├── main.go
						├── pkg1/
						├── pkg2/
						└── vendor/

				(2).一个项目中有多个module：
					Go 1.11引入的module是一组同属于一个版本管理单元的包的集合。
					并且Go支持在一个项目/仓库中存在多个module，但这种管理方式可能要比一定比例的代码重复引入更多的复杂性。 
					因此，如果项目结构中存在版本管理的“分歧”，
					比如：app1和app2的发布版本并不总是同步的，那么我建议你将项目拆分为多个项目（仓库），每个项目单独作为一个module进行单独的版本管理和演进。
					
					图解：
						multi-modules
						├── go.mod // mainmodule
						├── module1
						│   └── go.mod // module1
						└── module2
							└── go.mod // module2
	
					说明：
						我们可以通过git tag名字来区分不同module的版本。
						其中vX.Y.Z形式的tag名字用于代码仓库下的mainmodule；
						module1/vX.Y.Z形式的tag名字用于指示module1的版本；
						module2/vX.Y.Z形式的tag名字用于指示module2版本。

			2️⃣Go库项目布局：

				标准：
					lib-layout
					├── go.mod
					├── internal/
					│   ├── pkga/
					│   │   └── pkg_a.go
					│   └── pkgb/
					│       └── pkg_b.go
					├── pkg1/
					│   └── pkg1.go
					└── pkg2/
						└── pkg2.go

				简化：
					single-pkg-lib-layout
					├── feature1.go
					├── feature2.go
					├── go.mod
					└── internal/
	`)

	fmt.Println(`
	2.【Go Module发展史】
		(1).背景：
			Go程序由Go包组合而成的，Go程序的构建过程就是确定包版本、编译包以及将编译后得到的目标文件链接在一起的过程。
			Go语言的构建模式历经了三个迭代和演化过程，分别是最初期的GOPATH、1.5版本的Vendor机制，以及现在的Go Module。

		(2).迭代：
			1️⃣GOPATH的构建模式：
				Go语言在首次开源时，就内置了一种名为GOPATH的构建模式。
				在这种构建模式下，Go编译器可以在本地GOPATH环境变量配置的路径下，搜寻Go程序依赖的第三方包。
				如果存在，就使用这个本地包进行编译；如果不存在，就会报编译错误。

				举个例子：
					假定Go程序导入了github.com/user/repo这个包，我们也同时假定当前GOPATH环境变量配置的值为：
					export GOPATH=/usr/local/goprojects:/home/zhangp/go。

					那么在GOPATH构建模式下，Go编译器在编译Go程序时，就会在下面两个路径下搜索第三方依赖包是否存在：
					/usr/local/goprojects/src/github.com/user/repo
					/home/zhangp/go/src/github.com/user/repo

				注意⚠️：
					如果你没有显式设置GOPATH环境变量，Go会将GOPATH设置为默认值，
					不同操作系统下默认值的路径不同，在macOS或Linux上，它的默认值是$HOME/go。

				缺点：
					如果在GOPATH路径下没有搜索到第三方依赖包，这个时候就需要go get命令下载依赖包。
					
					不过，go get下载的包只是那个时刻各个依赖包的最新主线版本，这样会给后续Go程序的构建带来一些问题。
					比如，依赖包持续演进，可能会导致不同开发者在不同时间获取和编译同一个Go包时，得到不同的结果，
					也就是不能保证可重现的构建（Reproduceable Build）。
					又比如，如果依赖包引入了不兼容代码，程序将无法通过编译。
					如果依赖包因引入新代码而无法正常通过编译，并且该依赖包的作者又没用及时修复这个问题，这种错误也会传导到你的程序，
					导致你的程序无法通过编译。

					也就是说，在GOPATH构建模式下，Go编译器实质上并没有关注Go项目所依赖的第三方包的版本。
					但Go开发者希望自己的Go项目所依赖的第三方包版本能受到自己的控制，而不是随意变化。
					
	
			2️⃣vendor的构建模式：
				vendor机制本质上就是在Go项目的某个特定目录下，将项目的所有依赖包缓存起来，这个特定目录名就是vendor。
				Go编译器会优先感知和使用vendor目录下缓存的第三方包版本，而不是GOPATH环境变量所配置的路径下的第三方包版本。
				这样，无论第三方依赖包自己如何变化，无论GOPATH环境变量所配置的路径下的第三方包是否存在、版本是什么，都不会影响到Go程序的构建。
				因为，将vendor目录和项目源码一样提交到代码仓库，那么其他开发者下载你的项目后，就可以实现可重现的构建。
				因此，如果使用vendor机制管理第三方依赖包，最佳实践就是将vendor一并提交到代码仓库中。
				
				举个例子：
					.
					├── main.go
					└── vendor/
						├── github.com/
						│   └── sirupsen/
						│       └── logrus/
						└── golang.org/
							└── x/
								└── sys/
									└── unix/

				注意⚠️：
					要想开启vendor机制，你的Go项目必须位于GOPATH环境变量配置的某个路径的src目录下面。
					如果不满足这一路径要求，那么Go编译器是不会理会Go项目目录下的vendor目录的。

				缺点：
					vendor机制虽然一定程度解决了Go程序可重现构建的问题，但对开发者来说，它的体验却不那么好。
					一方面，Go项目必须放在GOPATH环境变量配置的路径下，庞大的vendor目录需要提交到代码仓库，不仅占用代码仓库空间，
					减慢仓库下载和更新的速度，而且还会干扰代码评审，对实施代码统计等开发者效能工具也有比较大影响。

					另外，你还需要手工管理vendor下面的Go依赖包，包括项目依赖包的分析、版本的记录、依赖包获取和存放，等等，最让开发者头疼的就是这一点。

			3️⃣Go Module构建模式：
				一个Go Module是一个Go包的集合。
				module是有版本的，所以module下的包也就有了版本属性。
				这个module与这些包会组成一个独立的版本单元，它们一起打版本、发布和分发。
				在Go Module模式下，通常一个代码仓库对应一个Go Module。
				一个Go Module的顶层目录下会放置一个go.mod文件，每个go.mod文件会定义唯一一个module，也就是说Go Module与go.mod是一一对应的。
				go.mod文件所在的顶层目录也被称为module的根目录，module根目录以及它子目录下的所有Go包均归属于这个Go Module，这个module也被称为main module。

				由go mod tidy下载的依赖module会被放置在本地的module缓存路径下，默认值为$GOPATH[0]/pkg/mod，
				Go 1.15及以后版本可以通过GOMODCACHE环境变量，自定义本地module的缓存路径。
	`)
}
