package errorZhang

import (
	"errors"
	"fmt"
)

func Exec() {

	fmt.Printf(`
	1.【Error】
		(1).定义：
			1️⃣背景：
				Go 语言中把错误当成一种特殊的值来处理，不支持其他语言中使用try/catch捕获异常的方式。

			2️⃣实现：
				Go 语言中使用一个名为 error 接口来表示错误类型。
				error 接口只包含一个方法——Error，这个函数需要返回一个描述错误信息的字符串。

				具体如下：
					type error interface {
						Error() string
					}
			
		(2).规则：
			1️⃣当一个函数或方法需要返回错误时，我们通常是把错误作为最后一个返回值。

				举个例子：
					func Open(name string) (*File, error) {
						return OpenFile(name, O_RDONLY, 0)
					}

			2️⃣由于 error 是一个接口类型，默认零值为nil。所以我们通常将调用函数返回的错误与nil进行比较，以此来判断函数是否返回错误。

				举个例子：
					file, err := os.Open("./xx.go")
					if err != nil {
						fmt.Println("打开文件失败,err:", err)
						return
					}
		
		(3).API：
			1️⃣创建错误
				方式一：func New(text string) error   它接收一个字符串参数返回包含该字符串的错误

				举个例子：
					func queryById(id int64) (*Info, error) {
						if id <= 0 {
							return nil, errors.New("无效的id")
						}
					
						// ...
					}

				方式二：fmt.Errorf  当我们需要传入格式化的错误描述信息时，使用fmt.Errorf是个更好的选择。
			
				举个例子：
					fmt.Errorf("查询数据库失败，err:%v", err)
				
				方式三：但是方式二会丢失原有的错误类型，只拿到错误描述的文本信息。
				为了不丢失函数调用的错误链，使用fmt.Errorf时搭配使用特殊的格式化动词%w，可以实现基于已有的错误再包装得到一个新的错误。
				
				举个例子：
					fmt.Errorf("查询数据库失败，err:%w", err)

			2️⃣其他操作
				func Unwrap(err error) error                 // 获得err包含下一层错误
				func Is(err, target error) bool              // 判断err是否包含target
				func As(err error, target interface{}) bool  // 判断err是否为target类型
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	var EOF = errors.New("EOF")
	//fmt.Printf("%T\n", EOF)
	//err := fmt.Errorf("查询数据库失败，err:%v", EOF)
	//fmt.Printf("%T\n", err)
	err := fmt.Errorf("查询数据库失败，err:%w", EOF)
	fmt.Printf("%T\n", err)

	fmt.Println(err)

	EOF1 := errors.Unwrap(err)
	fmt.Printf("%T\n", EOF1)
	fmt.Println(EOF1)

}
