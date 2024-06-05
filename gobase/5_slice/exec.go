package sliceZhang

import "fmt"

func Exec() {
	/**
	切片
	*/
	fmt.Printf(`
	1.【切片】
		(1).定义：
			因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。
			切片（Slice）是一个拥有相同类型元素的可变长度的序列。
			它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
			切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

		(2).声明格式：
			var name []T
			其中，
				1️⃣name:表示变量名
				2️⃣T:表示切片中的元素类型

				举个例子：
					var a []string              //声明一个字符串切片（注意：未初始化）
					var b = []int{}             //声明一个整型切片并初始化
					var c = []bool{false, true} //声明一个布尔切片并初始化
					var d = []bool{false, true} //声明一个布尔切片并初始化
					fmt.Println(a)              //[]
					fmt.Println(b)              //[]
					fmt.Println(c)              //[false true]
					fmt.Println(a == nil)       //true
					fmt.Println(b == nil)       //false
					fmt.Println(c == nil)       //false
					// fmt.Println(c == d)   	//切片是引用类型，不支持直接比较，只能和nil比较

		(3).结构：
			切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
			切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
			
		(4).初始化：
			1️⃣b := []T{}
			2️⃣c := []T{元素1, 元素2.....}
			3️⃣d := make([]T, size, cap)
			其中：
				T:切片的元素类型
				size:切片中元素的数量
				cap:切片的容量

			举个例子：
				a := make([]int, 2, 10)
				fmt.Println(a)      	//[0 0]
				fmt.Println(len(a)) 	//2
				fmt.Println(cap(a)) 	//10

		(5).切片的注意事项：
			1️⃣切片的判断
				要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
			2️⃣切片的比较
				切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 
				切片唯一合法的比较操作是和nil比较。 
				一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
				但是我们不能说一个长度和容量都是0的切片一定是nil。
			
				举个例子：
					var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
					s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
					s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
			3️⃣切片的拷贝
				如果两个切片共享底层数组，对一个切片的修改就会影响另一个切片的内容，这点需要特别注意。
			
				举个例子：
					s1 := make([]int, 3) 	//[0 0 0]
					s2 := s1             	//将s1直接赋值给s2，s1和s2共用一个底层数组
					s2[0] = 100
					fmt.Println(s1) 		//[100 0 0]
					fmt.Println(s2) 		//[100 0 0]
		
		(6).切片表达式：
			切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 

			1️⃣简单切片表达式
				格式：
					s := a[low:high]
				解释：
					切片表达式中的low和high表示一个索引范围（左包含，右不包含），
					也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，
					得到的切片长度high-low，容量等于底层数组的cap-low
					s(len) = high - low
					s(cap) = len(a) - low

					举个例子：
						a := [5]int{1, 2, 3, 4, 5}
						s := a[1:3]  // s := a[low:high]
						fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)
					输出：
						s:[2 3] len(s):2 cap(s):4
				
				特殊格式：
					a[2:]  // 等同于 a[2:len(a)]
					a[:3]  // 等同于 a[0:3]
					a[:]   // 等同于 a[0:len(a)]

			2️⃣完整切片表达式
				格式：
					s:= a[low : high : max]
				解释：
					它会将得到的结果切片的容量设置为max-low。
					在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。

					举个例子：
						a := [5]int{1, 2, 3, 4, 5}
						t := a[1:3:5]
						fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
					输出：
						t:[2 3] len(t):2 cap(t):4

		(7).遍历：
			切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。

			举个例子：
				s := []int{1, 3, 5}

				for i := 0; i < len(s); i++ {
					fmt.Println(i, s[i])
				}
			
				for index, value := range s {
					fmt.Println(index, value)
				}
	`)

	/**
	切片的操作
	*/
	fmt.Println(`
	2.【切片的操作】
		(1).append：
			1️⃣作用：
				添加：
					Go语言的内建函数append()可以为切片动态添加元素。
					可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。

					举个例子：
						var s []int
						s = append(s, 1)        // [1]
						s = append(s, 2, 3, 4)  // [1 2 3 4]
						s2 := []int{5, 6, 7}  
						s = append(s, s2...)    // [1 2 3 4 5 6 7]

				删除：
					Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。
					总结一下就是：要从切片a中删除索引为index的元素，
					操作方法是a = append(a[:index], a[index+1:]...)

					举个例子：
						// 从切片中删除元素
						a := []int{30, 31, 32, 33, 34, 35, 36, 37}
						// 要删除索引为2的元素
						a = append(a[:2], a[3:]...)
						fmt.Println(a) //[30 31 33 34 35 36 37]


			2️⃣特殊：
				通过var声明的零值切片可以在append()函数直接使用，无需初始化。
			
				举个例子：
					var s []int
					s = append(s, 1, 2, 3)

				没有必要这么写：
					s := []int{}  // 没有必要初始化
					s = append(s, 1, 2, 3)
					
					var s = make([]int)  // 没有必要初始化
					s = append(s, 1, 2, 3)

		(2).copy:
			1️⃣作用：
				由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
				Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中。

			2️⃣格式：
				copy(destSlice, srcSlice []T)
				其中：
					srcSlice: 数据来源切片
					destSlice: 目标切片

				举个例子：
					// copy()复制切片
					a := []int{1, 2, 3, 4, 5}
					c := make([]int, 5, 5)
					copy(c, a)     	//使用copy()函数将切片a中的元素复制到切片c
					fmt.Println(a) 	//[1 2 3 4 5]
					fmt.Println(c) 	//[1 2 3 4 5]
					c[0] = 1000
					fmt.Println(a) 	//[1 2 3 4 5]
					fmt.Println(c) 	//[1000 2 3 4 5]
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	切片定义，切片（Slice）是一个拥有相同类型元素的可变长度的序列。
	它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	（因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性）
	数组是数值类型，切片是引用类型（它的内部结构包含地址、长度和容量）

	切片的本质：就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

	定义格式：【var 变量名称 [] 数据类型】 ，如：var name []string
	*/

	//切片定义、初始化
	var a []string              //仅仅是声明
	var b = []int{}             //声明并初始化
	var d []int                 //仅仅是声明
	var c = []bool{false, true} //声明并初始化

	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(d == nil)

	arr := [5]int{1, 2, 3, 4}
	//也是声明并初始化的一种，基于数组声明并初始化切片
	s := arr[1:3] //s:=arr[low;high]  左包含，右不包含
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", arr, len(arr), cap(arr))

	s1 := arr[2:] // 等同于 a[2:len(a)]
	s2 := arr[:3] // 等同于 a[0:3]
	s3 := arr[:]  // 等同于 a[0:len(a)]

	fmt.Println(s1, s2, s3)

	/**
	make方式构造切片  格式：【make([]T, size, cap)】 注意：要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	*/
	arr1 := make([]int, 2, 10)
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	//切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。
	//一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	//但是我们不能说一个长度和容量都是0的切片一定是nil
	var s11 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s21 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s31 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(s11, s21, s31)

	/**
	切面赋值拷贝，拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	*/
	s41 := make([]int, 3) //[0,0,0]
	s51 := s41            //将s1直接赋值给s2，s1和s2共用一个底层数组
	fmt.Println(s51)
	fmt.Println(s41)
	s51[0] = 21
	fmt.Println(s41)
	fmt.Println(s51)

	/**
	切片的遍历
	*/
	s61 := []int{1, 3, 5}
	for i := 0; i < len(s61); i++ {
		fmt.Println(i, s61[i])
	}

	for index, value := range s61 {
		fmt.Println(index, value)
	}

	/**
	切片的操作:append
	*/
	s61 = append(s61, 1)
	fmt.Println(s61)
	s61 = append(s61, 2, 3, 4)
	fmt.Println(s61)
	s71 := []int{5, 6, 7}
	s71 = append(s71, s61...) //追加其他切片中的元素
	fmt.Println(s71)

	/**
	切片的扩容
	*/

	var numSlice []int //len(s1)=0;cap(s1)=0;s1==nil
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", numSlice, len(numSlice), cap(numSlice))

	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var citySlice []string
	citySlice = append(citySlice, "北京", "上海", "深圳")
	citySliceTmp := []string{"苏州", "南京"}
	citySlice = append(citySlice, citySliceTmp...)
	fmt.Println(citySlice)

	//append不止可以实现为切片追加元素，还可以为切片删除元素
	e1 := []int{1, 2, 3, 4, 5}
	//表示要删除索引为2的元素
	e1 = append(e1[:2], e1[3:]...)

	/**
	切片操作 copy
	*/
	//因为切片是引用类型，这种方式可以实现复制，但是f1、v1依然指向一个空间，改动f1，v1也会发生变化
	v1 := []int{1, 2, 3, 4, 5}
	f1 := v1
	fmt.Println(v1)
	fmt.Println(f1)

	f1[0] = 1000
	fmt.Println(f1)
	fmt.Println(v1)

	//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	v2 := []int{1, 2, 3}
	f2 := make([]int, 3, 3)
	copy(f2, v2)
	fmt.Println(f2)

	f2[0] = 100
	fmt.Println(v2)
	fmt.Println(f2)

}
