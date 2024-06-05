package fileZhang

import (
	"fmt"
)

func Exec() {

	fmt.Println(`
	1.【文件】
		(1).打开和关闭：
			os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件。

			举个例子：
				func main() {
					// 只读方式打开当前目录下的main.go文件
					file, err := os.Open("./main.go")
					if err != nil {
						fmt.Println("open file failed!, err:", err)
						return
					}
					// 关闭文件
					file.Close()
				}

			为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。

		(2).读取文件
			func (f *File) Read(b []byte) (n int, err error)
			
	`)

	//// 只读方式打开当前目录下的main.go文件
	//file, err := os.Open("./ma.go")
	//if err != nil {
	//	fmt.Println("open file failed!, err:", err)
	//	return
	//}
	//// 关闭文件
	//file.Close()
}
