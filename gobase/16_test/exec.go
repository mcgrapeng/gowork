package testZhang

import "fmt"

func Exec() {

	fmt.Printf(`
	1.【go test工具】
		(1).定义：
			Go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。
			go test命令是一个按照一定约定和组织的测试代码的驱动程序。
			在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。

		(2).分类：
			在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
			类型		格式					作用
			测试函数	函数名前缀为Test		测试程序的一些逻辑行为是否正确
			基准函数	函数名前缀为Benchmark	测试函数的性能
			示例函数	函数名前缀为Example	为文档提供示例文档

		(3).测试函数：
			1️⃣定义：
				测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头.
			2️⃣格式：
				func TestName(t *testing.T){
					// ...
				}

				举个例子：
					func TestAdd(t *testing.T){ ... }
					func TestSum(t *testing.T){ ... }
					func TestLog(t *testing.T){ ... }

			3️⃣测试组+子测试：
				func TestSplit(t *testing.T) {
					type test struct { // 定义test结构体
						input string
						sep   string
						want  []string
					}
					tests := map[string]test{ // 测试用例使用map存储
						"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
						"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
						"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
						"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
					}
					for name, tc := range tests {
						t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
							got := Split(tc.input, tc.sep)
							if !reflect.DeepEqual(got, tc.want) {
								t.Errorf("expected:%#v, got:%#v", tc.want, got)
							}
						})
					}
				}

			4️⃣覆盖率：
				测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。
				Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。
				
				举个例子：
					split $ go test -cover
					PASS
					coverage: 100.0% of statements
					ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s

				从上面的结果可以看到我们的测试用例覆盖了100%的代码。
				Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。
				
				举个例子：
					split $ go test -cover -coverprofile=c.out
					PASS
					coverage: 100.0% of statements
					ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s

					上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，然后我们执行go tool cover -html=c.out，
					使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。

		(4).基准函数：
		(5).示例函数：
	`)

	fmt.Println("============================================================以下是例子=========================================================")
}
