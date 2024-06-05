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

			3️⃣子测试：
				常单元测试中需要多组测试数据保证测试的效果。
				Go1.7+中新增了子测试，支持在测试函数中使用t.Run执行一组测试用例，这样就不需要为不同的测试数据定义多个测试函数了。
					
				格式：
					func TestXXX(t *testing.T){
					  t.Run("case1", func(t *testing.T){...})
					  t.Run("case2", func(t *testing.T){...})
					  t.Run("case3", func(t *testing.T){...})
					}
			
			4️⃣表格驱动测试：
				1.定义：
					表格驱动测试不是工具、包或其他任何东西，它只是编写更清晰测试的一种方式和视角。
					编写好的测试并非易事，但在许多情况下，
					表格驱动测试可以涵盖很多方面：表格里的每一个条目都是一个完整的测试用例，包含输入和预期结果，有时还包含测试名称等附加信息，以使测试输出易于阅读。
					使用表格驱动测试能够很方便的维护多个测试用例，避免在编写单元测试时频繁的复制粘贴。
					表格驱动测试的步骤通常是定义一个测试用例表格，然后遍历表格，并使用t.Run对每个条目执行必要的测试。
				
					举个例子：
						func TestSplitAll(t *testing.T) {
							// 定义测试表格
							// 这里使用匿名结构体定义了若干个测试用例
							// 并且为每个测试用例设置了一个名称
							tests := []struct {
								name  string
								input string
								sep   string
								want  []string
							}{
								{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
								{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
								{"more sep", "abcd", "bc", []string{"a", "d"}},
								{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
							}
							// 遍历测试用例
							for _, tt := range tests {
								t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
									got := Split(tt.input, tt.sep)
									if !reflect.DeepEqual(got, tt.want) {
										t.Errorf("expected:%#v, got:%#v", tt.want, got)
									}
								})
							}
						}

					在终端执行go test -v，会得到如下测试输出结果：
						> go test -v
						=== RUN   TestSplit
						--- PASS: TestSplit (0.00s)
						=== RUN   TestSplitWithComplexSep
						--- PASS: TestSplitWithComplexSep (0.00s)
						=== RUN   TestSplitAll
						=== RUN   TestSplitAll/base_case
						=== RUN   TestSplitAll/wrong_sep
						=== RUN   TestSplitAll/more_sep
						=== RUN   TestSplitAll/leading_sep
						--- PASS: TestSplitAll (0.00s)
							--- PASS: TestSplitAll/base_case (0.00s)
							--- PASS: TestSplitAll/wrong_sep (0.00s)
							--- PASS: TestSplitAll/more_sep (0.00s)
							--- PASS: TestSplitAll/leading_sep (0.00s)
						PASS
						ok      gotest/1_base/example 0.010s

				2.并行：
					表格驱动测试中通常会定义比较多的测试用例，而Go语言又天生支持并发，所以很容易发挥自身并发优势将表格驱动测试并行化。 
					想要在单元测试过程中使用并行测试，可以像下面的代码示例中那样通过添加t.Parallel()来实现。

					举个例子：
						func TestSplitAll(t *testing.T) {
							t.Parallel()  // 将 TLog 标记为能够与其他测试并行运行
							// 定义测试表格
							// 这里使用匿名结构体定义了若干个测试用例
							// 并且为每个测试用例设置了一个名称
							tests := []struct {
								name  string
								input string
								sep   string
								want  []string
							}{
								{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
								{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
								{"more sep", "abcd", "bc", []string{"a", "d"}},
								{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
							}
							// 遍历测试用例
							for _, tt := range tests {
								tt := tt  // 注意这里重新声明tt变量（避免多个goroutine中使用了相同的变量）
								t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
									t.Parallel()  // 将每个测试用例标记为能够彼此并行运行
									got := Split(tt.input, tt.sep)
									if !reflect.DeepEqual(got, tt.want) {
										t.Errorf("expected:%#v, got:%#v", tt.want, got)
									}
								})
							}
						}

					执行go test -v的时候就会看到每个测试用例并不是按照我们定义的顺序执行，而是互相并行了。
				
				3.工具：
					社区里有很多自动生成表格驱动测试函数的工具，比如gotests等，很多编辑器如Goland也支持快速生成测试文件。这里简单演示一下gotests的使用。
					
					(1).安装：
						go get -u github.com/cweill/gotests/...

					(2).执行：
						gotests -all -w split.go
						上面的命令表示，为split.go文件的所有函数生成测试代码至split_test.go文件（目录下如果事先存在这个文件就不再生成）。

					生成代码大致如下：
						package example

						import (
							"reflect"
							"testing"
						)
						
						func TestSplit(t *testing.T) {
							type args struct {
								s   string
								sep string
							}
							tests := []struct {
								name       string
								args       args
								wantResult []string
							}{
								// TODO: Add test cases.
							}
							for _, tt := range tests {
								t.Run(tt.name, func(t *testing.T) {
									if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
										t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
									}
								})
							}
						}

					代码格式与我们上面的类似，只需要在TODO位置添加我们的测试逻辑就可以了。

			5️⃣覆盖率：
				测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。
				在公司内部一般会要求测试覆盖率达到80%左右。
				Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。
				
				举个例子：
					split $ go test -cover
					PASS
					coverage: 100.0% of statements
					ok      github.com/zhangp/gowork/gotest/example/split       0.005s

				从上面的结果可以看到我们的测试用例覆盖了100%的代码。
				Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。
				
				举个例子：
					split $ go test -cover -coverprofile=c.out
					PASS
					coverage: 100.0% of statements
					ok      github.com/zhangp/gowork/gotest/example/split       0.005s

				> tree .
				.
				├── c.out
				├── split.go
				└── split_test.go

				上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，然后我们执行go tool cover -html=c.out，
				使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。(用绿色标记的语句块表示被覆盖了，而红色的表示没有被覆盖。)
			
			6️⃣命令：
				go test
				解释：
					最基本的测试命令，适合只有一个测试用例的场景。(也就是一个测试文件(以_test.go为后缀名的源代码文件)只有一个测试方法)

					举个例子：
						> go test
						PASS
						ok      gotest/1_base/example       0.003s

				go test -v
				解释：
					当存在多个测试用例的时候，为了能更好的在输出结果中看到每个测试用例的执行情况，我们可以为go test命令添加-v参数，让它输出完整的测试结果。
					
					举个例子：
						> go test -v
						=== RUN   TestSplit
						--- PASS: TestSplit (0.00s)
						=== RUN   TestSplitWithComplexSep
							split_test.go:20: expected:[a d], got:[a cd]
						--- FAIL: TestSplitWithComplexSep (0.00s)
						FAIL
						exit status 1
						FAIL    gotest/1_base/example 0.009s

				go test -run
				解释：
					在执行go test命令的时候可以添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。

					举个例子：
						给go test添加-run=Sep参数来告诉它本次测试只运行TestSplitWithComplexSep这个测试用例
						> go test -run=Sep -v
						=== RUN   TestSplitWithComplexSep
						--- PASS: TestSplitWithComplexSep (0.00s)
						PASS
						ok      gotest/1_base/example 0.010s

			7️⃣回归测试：
				我们修改了代码之后仅仅执行那些失败的测试用例或新引入的测试用例是错误且危险的，
				正确的做法应该是完整运行所有的测试用例，保证不会因为修改代码而引入新的问题。

				举个例子：
					> go test -v
					=== RUN   TestSplit
					--- PASS: TestSplit (0.00s)
					=== RUN   TestSplitWithComplexSep
					--- PASS: TestSplitWithComplexSep (0.00s)
					PASS
					ok      gotest/1_base/example 0.011s

				以上测试结果表明我们的单元测试全部通过。
				通过这个示例我们可以看到，有了单元测试就能够在代码改动后快速进行回归测试，极大地提高开发效率并保证代码的质量。

			8️⃣跳过：
				为了节省时间支持在单元测试时跳过某些耗时的测试用例。

				举个例子：
					func TestTimeConsuming(t *testing.T) {
						if testing.Short() {
							t.Skip("short模式下会跳过该测试用例")
						}
						...
					}

				当执行go test -short时就不会执行上面的TestTimeConsuming测试用例。
			
			9️⃣testify/assert
				testify是一个社区非常流行的Go单元测试工具包，其中使用最多的功能就是它提供的断言工具——testify/assert或testify/require。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

}
