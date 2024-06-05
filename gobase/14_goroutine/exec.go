package goroutineZhang

import (
	"fmt"
	"sync"
	"time"
)

/*
*
并发安全和锁
*/
var (
	x int64

	wg sync.WaitGroup //等待组

	m sync.Mutex // 互斥锁

	rwMutex sync.RWMutex //读写锁

	mp = make(map[string]int)
)

func Exec() {

	fmt.Println(`
	1.【goroutine】
		(1).定义：
			Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，
			一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。
			
			操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，
			一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。
			并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。

			区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。
			例如Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。
			Goroutine 是 Go 程序中最基本的并发执行单元。

			每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。
			在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，当你需要让某个任务并发执行的时候，
			你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。
			一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。

		(2).go关键字:
			1️⃣定义：
			Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，
			从而让该函数或方法在新创建的 goroutine 中执行。

				举个例子：
					go f()  // 创建一个新的 goroutine 运行函数f
	
				匿名函数也支持使用go关键字创建 goroutine 去执行：
					go func(){
					  // ...
					}()

			2️⃣生命周期：
			当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，
			所有由 main goroutine 创建的 goroutine 也会一同退出。
			
		(3).调度：
			1️⃣定义：
				区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，
				是完全由 Go 语言本身实现的一套调度系统——go scheduler。它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。
				
				在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型。
	
				其中：
					【G】：	表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息。
	
					【全局队列】：存放等待运行的 G。
					
					【P】：	表示 goroutine 执行所需的资源，内部包含一个队列，最多有 GOMAXPROCS 个。
							同全局队列类似，存放的也是等待运行的 G，存的数量有限，不超过256个。
							新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列。

					【GOMAXPROCS】：	Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。
									默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。
					
					【M】：	Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，
							操作系统调度器负责把内核线程分配到 CPU 的核上执行。

			2️⃣流程：
				线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。
				M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。

			3️⃣优势：
				单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的， 
				goroutine 则是由Go运行时（runtime）自己的调度器调度的，完全是在用户态下完成的， 
				不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 
				不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 
				另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 
				再加上本身 goroutine 的超轻量级，以上种种特性保证了 goroutine 调度方面的性能。
	`)

	fmt.Printf(`
	2.【channel】
		(1).定义：
			单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
			虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。
			为了保证数据交换的正确性，很多并发模型中必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

			于是，Go语言使用一种新的方式解决以上问题，即channel。

			如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。
			channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
			Go 语言中的通道（channel）是一种特殊的类型。
			通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
			每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
		
		(2).声明：
			channel是 Go 语言中一种特有的类型。
			var 变量名称 chan 元素类型

			其中：
				chan：是关键字
				元素类型：是指通道中传递元素的类型

			举个例子：
				var ch1 chan int   // 声明一个传递整型的通道
				var ch2 chan bool  // 声明一个传递布尔型的通道
				var ch3 chan []int // 声明一个传递int切片的通道

		(3).初始化：
			声明的通道类型变量需要使用内置的make函数初始化之后才能使用。
			1️⃣格式：
				make(chan 元素类型, [缓冲大小])

				其中：
					channel的缓冲大小是可选的。

				举个例子：
					ch4 := make(chan int)  //无缓冲通道，详情请往后看
					ch5 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道

			2️⃣默认：
				未初始化的通道类型变量其默认零值是nil。
			
				举个例子：
					var ch chan int
					fmt.Println(ch) // <nil>

		(4).操作：
			通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
			1️⃣发送：
				将一个值发送到通道中。

				举个例子：
					ch <- 10 // 把10发送到ch中

			2️⃣接收：
				从一个通道中接收值。

				举个例子：
					x := <- ch // 从ch中接收值并赋值给变量x
					<-ch       // 从ch中接收值，忽略结果
				
			3️⃣关闭：
				我们通过调用内置的close函数来关闭通道。

				举个例子：
					close(ch)
			
			4️⃣注意：
				1.对一个关闭的通道再发送值就会导致 panic。
				2.对一个关闭的通道进行接收会一直获取值直到通道为空。
				3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
				4.关闭一个已经关闭的通道会导致 panic。

		(5).分类：
			1️⃣无缓冲的通道：
				使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道或者阻塞的通道。
				无缓冲的通道只有在有另一个goroutine接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
				同理，如果对一个无缓冲通道执行接收操作时，没有任何goroutine向通道中发送值的操作那么也会导致接收操作阻塞。

				举个例子：
					func main() {
						ch := make(chan int)
						ch <- 10
						fmt.Println("发送成功")
					}

				输出：
					fatal error: all goroutines are asleep - deadlock!
				
				解释：
					deadlock表示我们程序中的 goroutine 都被挂起导致程序死锁了。

				解决：
					func recv(c chan int) {
						ret := <-c
						fmt.Println("接收成功", ret)
					}
					
					func main() {
						ch := make(chan int)
						go recv(ch) // 创建一个 goroutine 从通道接收值
						ch <- 10
						fmt.Println("发送成功")
					}

			2️⃣有缓冲的通道：
				只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
				当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。

				举个例子：
					func main() {
						ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
						ch <- 10
						fmt.Println("发送成功")
					}

		(6).通道接收处理：
			1️⃣空值判断
				当向通道中发送完数据时，我们可以通过close函数来关闭通道。
				当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。
				通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。
				那我们如何判断一个通道是否被关闭了呢？那就是对一个通道执行接收操作时支持使用如下多返回值模式。

				格式：
					value, ok := <- ch
					
					其中：
						value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
						ok：通道ch关闭时返回 false，否则返回 true。
	
					举个例子：
						func f2(ch chan int) {
							for {
								v, ok := <-ch
								if !ok {
									fmt.Println("通道已关闭")
									break
								}
								fmt.Printf("v:%#v ok:%#v\n", v, ok)
							}
						}
						
						func main() {
							ch := make(chan int, 2)
							ch <- 1
							ch <- 2
							close(ch)
							f2(ch)
						}

			2️⃣for range接收值
				通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
				
				举个例子：
					func f3(ch chan int) {
						for v := range ch {
							fmt.Println(v)
						}
					}

			3️⃣注意⚠️
				**目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。
				不能简单的通过len(ch)操作来判断通道是否被关闭。
		
		(7).单向通道：
			在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，
			通常我们会选择在不同的任务函数中对通道的使用进行限制，比如限制通道在某个函数中只能执行发送或只能执行接收操作。

			格式：
				<- chan int // 只接收通道，只能接收不能发送
				chan <- int // 只发送通道，只能发送不能接收

			注意⚠️：
				另外对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。

			举个例子：
				// Producer2 返回一个接收通道
				func Producer2() <-chan int {
					ch := make(chan int, 2)
					// 创建一个新的goroutine执行发送数据的任务
					go func() {
						for i := 0; i < 10; i++ {
							if i%2 == 1 {
								ch <- i
							}
						}
						close(ch) // 任务完成后关闭通道
					}()
				
					return ch
				}
				
				// Consumer2 参数为接收通道
				func Consumer2(ch <-chan int) int {
					sum := 0
					for v := range ch {
						sum += v
					}
					return sum
				}
				
				func main() {
					ch2 := Producer2()
				  
					res2 := Consumer2(ch2)
					fmt.Println(res2) // 25
				}

			调用：
				var ch3 = make(chan int, 1)
				ch3 <- 10
				close(ch3)
				Consumer2(ch3) // 函数传参时将ch3转为单向通道
				
				var ch4 = make(chan int, 1)
				ch4 <- 10
				var ch5 <-chan int // 声明一个只接收通道ch5
				ch5 = ch4          // 变量赋值时将ch4转为单向通道
				<-ch5
	`)

	fmt.Println(`
	3.【select】
		(1).定义：
			在某些场景下我们可能需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以被接收那么当前 goroutine 将会发生阻塞。
			Go 语言内置了select关键字，使用它可以同时响应多个通道的操作。

			1️⃣流程：
				Select 的使用方式类似于之前学到的 switch 语句，它也有一系列 case 分支和一个默认的分支。
				每个 case 分支会对应一个通道的通信（接收或发送）过程。
				select 会一直等待，直到其中的某个 case 的通信操作完成时，就会执行该 case 分支对应的语句。

			2️⃣格式：
				select {
					case <-ch1:
						//...
					case data := <-ch2:
						//...
					case ch3 <- 10:
						//...
					default:
						//默认操作
				}

			3️⃣特点：
				1.可处理一个或多个 channel 的发送/接收操作。
				2.如果多个 case 同时满足，select 会随机选择一个执行。
				3.对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

			举个例子：
				func main() {
					ch := make(chan int, 1)
					for i := 1; i <= 10; i++ {
						select {
						case x := <-ch:
							fmt.Println(x)
						case ch <- i:
						}
					}
				}

			输出：
				1
				3
				5
				7
				9
	`)

	fmt.Println(`
	4.【并发安全和锁】
		(1).互斥锁：
			Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。

			sync.Mutex提供了两个方法：

			方法名						功能
			func (m *Mutex) Lock()		获取互斥锁
			func (m *Mutex) Unlock()	释放互斥锁

			举个例子：
				var m sync.Mutex // 互斥锁

				func add() {
					for i := 0; i < 5000; i++ {
						m.Lock() // 修改x前加锁
						x = x + 1
						m.Unlock() // 改完解锁
					}
				}

		(2).读写互斥锁：
			互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，
			当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。
			读写锁在 Go 语言中使用sync包中的RWMutex类型。

			sync.RWMutex提供了以下5个方法：
			
			方法名								功能
			func (rw *RWMutex) Lock()			获取写锁
			func (rw *RWMutex) Unlock()			释放写锁
			func (rw *RWMutex) RLock()			获取读锁
			func (rw *RWMutex) RUnlock()		释放读锁
			func (rw *RWMutex) RLocker() Locker	返回一个实现Locker接口的读写锁

		(3).任务同步:
			在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。
			sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
			例如当我们启动了 N 个并发任务时，就将计数器值增加N。
			每个任务完成时通过调用 Done 方法将计数器减1。
			通过调用 Wait 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。

			sync.WaitGroup有以下几个方法：

			方法名									功能
			func (wg * WaitGroup) Add(delta int)	计数器+delta
			(wg *WaitGroup) Done()					计数器-1
			(wg *WaitGroup) Wait()					阻塞直到计数器变为0

			举个例子：
				var wg sync.WaitGroup
				func hello() {
					defer wg.Done()
					fmt.Println("Hello Goroutine!")
				}
				func main() {
					wg.Add(1)
					go hello() // 启动另外一个goroutine去执行hello函数
					fmt.Println("main goroutine done!")
					wg.Wait()
				}
	`)

	/**
	go关键字
	*/
	go sayHello()
	fmt.Println("say main")
	time.Sleep(time.Second)

	/**
	goroutine调度
	*/
	fmt.Println(`
		操作系统内核在调度时会挂起当前正在执行的线程并将寄存器中的内容保存到内存中，
		然后选出接下来要执行的线程并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。
		从一个线程切换到另一个线程需要完整的上下文切换。
		因为可能需要多次内存访问，索引这个切换上下文的操作开销较大，会增加运行的cpu周期。

		区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，
		是完全由 Go 语言本身实现的一套调度系统——go scheduler。
		它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。

		在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型。

		其中：
			G：表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息。
			全局队列（Global Queue）：存放等待运行的 G。

			P：表示 goroutine 执行所需的资源，最多有 GOMAXPROCS 个。
			P 的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列。

			M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。

			Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，操作系统调度器负责把内核线程分配到 CPU 的核上执行。

			单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的， goroutine 则是由Go运行时（runtime）自己的调度器调度的，
			完全是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 
			不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，
			近似的把若干goroutine均分在物理线程上， 再加上本身 goroutine 的超轻量级，以上种种特性保证了 goroutine 调度方面的性能。
	`)
	/**
	GOMAXPROCS
	*/
	fmt.Println(`
		Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。
		默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。
		Go语言中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的 CPU逻辑核心数。（Go1.5版本之前，默认使用的是单核心执行。Go1.5 版本之后，默认使用全部的CPU 逻辑核心数。）
	`)

	/**
	channel
	*/
	fmt.Println(`
		Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
		如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。
		channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
		Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
		每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
	`)

	/**
	channel类型
	格式：var 变量名称 chan 元素类型
	其中：
		chan：是关键字
		元素类型：是指通道中传递元素的类型
	*/
	var ch1 chan int           // 声明一个传递整型的通道
	var ch2 chan bool          // 声明一个传递布尔型的通道
	var ch3 chan []int         // 声明一个传递int切片的通道
	fmt.Println(ch1, ch2, ch3) //未初始化的通道类型变量其默认零值是nil。

	/**
	初始化channel
	格式：make(chan 元素类型, [缓冲大小])

	其中：
	channel的缓冲大小是可选的。

	通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
	*/
	ch4 := make(chan int)
	ch5 := make(chan bool, 1)
	fmt.Println(ch4, ch5)

	//发送
	ch4 <- 10 //把10发送到ch4里

	//接收
	x := <-ch4 // 从ch4中接收值并赋值给变量x
	<-ch4      // 从ch4中接收值，忽略结果
	fmt.Println(x)

	//关闭通道
	close(ch4)

	fmt.Println(`
		**注意：**一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
		它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
		关闭后的通道有以下特点：
			对一个关闭的通道再发送值就会导致 panic。
			对一个关闭的通道进行接收会一直获取值直到通道为空。
			对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
			关闭一个已经关闭的通道会导致 panic。
	`)

	/**
	无缓冲管道
	无缓冲的通道又称为阻塞的通道。
	*/

	fmt.Println(`
		ch6 := make(chan int)
		ch6 <- 10
		fmt.Println(ch6)

		上面这段代码能够通过编译，但是执行的时候会出现以下错误：
		fatal error: all goroutines are asleep - deadlock!
		goroutine 1 [chan send]:
		main.main()
				.../main.go:8 +0x54

		deadlock表示我们程序中的 goroutine 都被挂起导致程序死锁了。为什么会出现deadlock错误呢？
		因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，
		否则会一直处于等待发送的阶段。同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
		就像田径比赛中的4x100接力赛，想要完成交棒必须有一个能够接棒的运动员，否则只能等待。简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。

		上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？
		其中一种可行的方法是创建一个 goroutine 去接收值，例如：
		
		ch7 := make(chan int)
		go recv(ch7)
		ch7 <- 10
		fmt.Println("发送成功")

		func recv(c chan int) {
			ret := <-c
			fmt.Println("接收成功", ret)
		}
	
		首先无缓冲通道ch上的发送操作会阻塞，直到另一个 goroutine 在该通道上执行接收操作，这时数字10才能发送成功，两个 goroutine 将继续执行。
		相反，如果接收操作先执行，接收方所在的 goroutine 将阻塞，直到 main goroutine 中向该通道发送数字10。
		使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。
	`)

	/**
	有缓冲通道

	只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
	当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
	就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
	我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
	*/
	ch8 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch8 <- 10
	fmt.Println("发送成功")

	fmt.Println(`
		当向通道中发送完数据时，我们可以通过close函数来关闭通道。
		当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。
		通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。那我们如何判断一个通道是否被关闭了呢？

		采用多返回值模式：
		value, ok := <- ch
		其中：
			value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
			ok：通道ch关闭时返回 false，否则返回 true。
	`)
	ch9 := make(chan int, 2)
	ch9 <- 1
	ch9 <- 2
	close(ch9)
	f2(ch9)

	/**
	for range接收值

	通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。上面那个示例我们使用for range改写后会很简洁。
	*/
	ch10 := make(chan int, 2)
	ch10 <- 1
	ch10 <- 2
	f3(ch10)

	fmt.Println(`
		**注意：**目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。
		不能简单的通过len(ch)操作来判断通道是否被关闭。
	`)

	/**
	单向通道
	*/
	fmt.Println(`
		在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，
		通常我们会选择在不同的任务函数中对通道的使用进行限制，
		比如限制通道在某个函数中只能执行发送或只能执行接收操作。
	`)

	ch11 := Producer()
	res := Consumer(ch11)
	fmt.Println(res)

	/**
	Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况。

	<- chan int // 只接收通道，只能接收不能发送
	chan <- int // 只发送通道，只能发送不能接收
	*/
	ch12 := Producer1()
	res2 := Consumer1(ch12)
	fmt.Println(res2)

	fmt.Println(`
		在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。
	`)

	var ch13 = make(chan int, 1)
	ch13 <- 10
	close(ch13)

	Consumer1(ch13)

	var ch14 = make(chan int, 1)
	ch14 <- 10
	var ch15 <-chan int // 声明一个只接收通道ch15
	ch15 = ch14         // 变量赋值时将ch14转为单向通道
	<-ch15

	fmt.Println(`**注意：**对已经关闭的通道再执行 close 也会引发 panic。`)

	/**
	select多路复用
	格式：
		select {
			case <-ch1:
				//...
			case data := <-ch2:
				//...
			case ch3 <- 10:
				//...
			default:
				//默认操作
		}

	Select 语句具有以下特点。
		1.可处理一个或多个 channel 的发送/接收操作。
		2.如果多个 case 同时满足，select 会随机选择一个执行。
		3.对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
	*/
	ch16 := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch16:
			fmt.Println(x)
		case ch16 <- i:
		}
	}

	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)

	/**
	互斥锁

	var m sync.Mutex
	m.Lock()
	m.Unlock()
	*/
	fmt.Println(`
		使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，
		其他的 goroutine 则在等待锁；当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，
		多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。
	`)

	/**
	读写互斥锁

		互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型。
	*/

	/**
	sync.Once

		在某些场景下我们需要确保某些操作即使在高并发的场景下也只会被执行一次，例如只加载一次配置文件等。
		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案——sync.Once，sync.Once只有一个Do方法，
	其签名如下：
		func (o *Once) Do(f func())
	*/

	/**
	sync.Map
	*/
	var mp1 = sync.Map{}
	mp1.Store("a1", 1)
	value2, _ := mp1.Load("a1")
	fmt.Println(value2)
}

func sayHello() {
	fmt.Println("say hello")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func Producer() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func Producer1() <-chan int {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer1(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// demo2 通道误用导致的bug
func demo2() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}

func add() {
	for i := 0; i < 10; i++ {
		m.Lock()
		x += 1
		m.Unlock()
	}
	wg.Done()
}

func writeWithLock() {
	m.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	m.Unlock()
	wg.Done()
}

func readWithLock() {
	m.Lock()
	time.Sleep(time.Millisecond)
	m.Unlock()
	wg.Done()
}

func writeWithRWLock() {
	rwMutex.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	rwMutex.Unlock()
	wg.Done()
}

func readWithRWLock() {
	rwMutex.RLock()
	time.Sleep(time.Millisecond)
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()

	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)
}

func get(key string) int {
	return mp[key]
}

func set(key string, value int) {
	mp[key] = value
}
