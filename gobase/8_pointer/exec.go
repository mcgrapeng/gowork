package pointerZhang

import "fmt"

func Exec() {

	fmt.Printf(`
	1.【指针】
		(1).定义：
			任何程序数据载入内存后，在内存都有他们的地址，这就是指针。
			而为了保存一个数据在内存中的地址，我们就需要指针变量。
		
		(2).声明格式：
			var 变量名 *变量类型
				举个例子：
					var a *int

		(3).变量取址：
			ptr := &v    // v的类型为T

			其中：
				v:代表被取地址的变量，类型为T
				ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。

			举个例子：
				a := 10
				b := &a
				fmt.Printf("a:%d ptr:%p\n", a, &a) 	// a:10 ptr:0xc00001a078
				fmt.Printf("b:%p type:%T\n", b, b) 	// b:0xc00001a078 type:*int
				fmt.Println(&b)                    	// 0xc00000e018

		(4).指针取值：
			ptr := *v
			
			举个例子：
				a := 10
				b := &a // 取变量a的地址，将指针保存到b中
				fmt.Printf("type of b:%T\n", b)
				c := *b // 指针取值（根据指针去内存取值）
				fmt.Printf("type of c:%T\n", c)
				fmt.Printf("value of c:%v\n", c)

			输出：
				type of b:*int
				type of c:int
				value of c:10

		(5).总结：
				取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

				变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
					1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
					2.指针变量的值是指针地址。
					3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	`)

	fmt.Println(`
	2.【new和make】
		(1).背景：
				在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
				而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
				要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
		
		(2).new：
			func new(Type) *Type

			举个例子：
				a := new(int)
				b := new(bool)
				fmt.Printf("%T\n", a) 	// *int
				fmt.Printf("%T\n", b) 	// *bool
				fmt.Println(*a)       	// 0
				fmt.Println(*b)       	// false

		(3).make：
			func make(t Type, size ...IntegerType) Type

			make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，
			而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

			举个例子：
				var b map[string]int
				b = make(map[string]int, 10)
				b["沙河娜扎"] = 100
				fmt.Println(b)

		(4).区别：
			二者都是用来做内存分配的。
			make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
			而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	Go语言的指针：任何程序数据载入内存后，在内存都有他们的地址，这就是指针。
	而为了保存一个数据在内存中的地址，我们就需要指针变量。

	Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
	Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。

	格式：ptr:= &v
	*/
	var v int = 10
	ptr := &v
	//ptr := v
	fmt.Printf("v:%d , ptr:%p\n", v, &v)
	fmt.Printf("ptr:%p type:%T\n", ptr, ptr)
	fmt.Println(&ptr)

	/**
	指针取值
	总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

	1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	2.指针变量的值是指针地址。
	3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	*/
	a := 10
	b := &a
	fmt.Printf("type of b:%T\n", b)
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	modity(a)
	fmt.Println(a)
	modity1(&a)
	fmt.Println(a)

	/**
	new和make
	*/
	fmt.Println(`
		var a1 *int
		*a1 = 100
		fmt.Println(*a1)
	
		var b1 map[string]int
		b1["南京"] = 100
		fmt.Println(b1)

		执行上面的代码会引发panic，为什么呢？ 
		在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
		而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
		要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存。
		
		new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。

		make也是用于内存分配的，区别于new，它只用于slice、map以及channel的内存创建，
		而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
		所以就没有必要返回他们的指针了。
	`)

	a2 := new(int)
	b2 := new(bool)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%T\n", b2)

	fmt.Println(*a2)
	fmt.Println(*b2)

	var a3 *int
	a3 = new(int)
	*a3 = 1000
	fmt.Println(*a3)

	fmt.Println(`
		var b map[string]int只是声明变量b是一个map类型的变量，
		需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值
		var b4 map[string]int
		b4 = make(map[string]int, 10)
		b4["苏州"] = 100
		fmt.Println(b4)
	`)

	fmt.Println(`
		new与make的区别
			二者都是用来做内存分配的。
			make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
			而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	`)
}

func modity(x int) {
	x = 100
}
func modity1(x *int) {
	*x = 100
}
