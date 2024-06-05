package contextZhang

import "fmt"

func Exec() {

	fmt.Println(`
	1.【Context】
		(1).定义：
			Go1.7加入了一个新的标准库context，它定义了Context类型，
			专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

			1️⃣结构：
				type Context interface {
					Deadline() (deadline time.Time, ok bool)
					Done() <-chan struct{}
					Err() error
					Value(key interface{}) interface{}
				}

				其中：
					1.Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
					2.Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
					3.Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
						--如果当前Context被取消就会返回Canceled错误；
						--如果当前Context超时就会返回DeadlineExceeded错误；
					4.Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

			2️⃣特殊函数：
				Go内置两个函数：Background()和TODO()。
				Background()：主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。
				TODO()：它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。
				background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

			3️⃣With系列函数：
				1.WithCancel
					格式：
						func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
			
	`)
}
