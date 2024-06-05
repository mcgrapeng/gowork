package arrayZhang

import "fmt"

func Exec() {

	/**
	数组
	*/
	fmt.Printf(`
	1.【数组】
		(1).定义：
			数组是同一种数据类型元素的集合。 
			在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
			数组的长度必须是常量，并且长度是数组类型的一部分。
			例如：
				var a [3]int
				var b [4]int
				a = b //不可以这样做，因为此时a和b是不同的类型

		(2).声明格式：
			var 数组变量名 [元素数量]T

		(3).初始化：
			1️⃣初始化数组时可以使用初始化列表来设置数组元素的值。
			举个例子：
				var testArray [3]int                        		//数组会初始化为int类型的零值
				var numArray = [3]int{1, 2}                 		//使用指定的初始值完成初始化
				var cityArray = [3]string{"北京", "上海", "深圳"} 	//使用指定的初始值完成初始化
			输出：
				[0 0 0]
				[1 2 0]
				[北京 上海 深圳]

			2️⃣按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度。
			举个例子：
				var numArray = [...]int{1, 2}
				var cityArray = [...]string{"北京", "上海", "深圳"}
			输出：
				[1 2]
				[北京 上海 深圳]

			3️⃣还可以使用指定索引值的方式来初始化数组
			举个例子：
				a := [...]int{1: 1, 3: 5}
			输出：
				[0 1 0 5]

		(4).遍历：
			1️⃣for 循环遍历
			for i := 0; i < len(a); i++ {
				fmt.Println(a[i])
			}
			
			2️⃣for range遍历
			for index, value := range a {
				fmt.Println(index, value)
			}
	`)

	/**
	多维数组
	*/
	fmt.Printf(`
	2.【多维数组】
		(1).定义：
			a := [3][2]string{
				{"北京", "上海"},
				{"广州", "深圳"},
				{"成都", "重庆"},
			}

			输出：
				[[北京 上海] [广州 深圳] [成都 重庆]]

			遍历：
				for _, v1 := range a {
					for _, v2 := range v1 {
						fmt.Printf("%s\t", v2)
					}
					fmt.Println()
				}

			输出：
				北京	上海	
				广州	深圳	
				成都	重庆

		(2).特殊写法：
			多维数组只有第一层可以使用...来让编译器推导数组长度。

			举个例子：
				//支持的写法
				a := [...][2]string{
					{"北京", "上海"},
					{"广州", "深圳"},
					{"成都", "重庆"},
				}
				//不支持多维数组的内层使用...
				b := [3][...]string{
					{"北京", "上海"},
					{"广州", "深圳"},
					{"成都", "重庆"},
				}
	`)

	/**
	值类型
	*/
	fmt.Println(`
	3.【值类型】
		数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。


		**注意**
			1.数组支持 “=="、”!=" 操作符，因为内存总是被初始化过的。
			2.[n]*T表示指针数组，*[n]T表示数组指针 。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	数组初始化
	*/
	//数组会初始化为int类型的零值
	var testArray [3]int

	//使用指定的初始值完成初始化
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"苏州", "南京", "杭州"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
	numArray = [...]int{3, 4, 5}
	cityArray = [...]string{"成都", "上海", "合肥"}
	fmt.Println(numArray)
	fmt.Println(cityArray)

	//可以指定索引值初始化
	ageArray := [...]int{1: 1, 3: 2}
	fmt.Println(ageArray)

	/**
	数组遍历
	*/
	var a = [...]string{"北京", "深圳", "上海"}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, value := range a {
		fmt.Println(index, value)
	}

	/**
	多维数组, 注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。
	*/
	b := [3][2]string{
		{"苏州", "南京"},
		{"杭州", "宁波"},
		{"合肥", "成都"},
	}
	fmt.Println(b)

	b = [...][2]string{
		{"西安", "天津"},
		{"重庆", "无锡"},
		{"常州", "上海"},
	}
	fmt.Println(b)

	/**
	二维数组遍历
	*/
	for _, val := range b {
		for _, val2 := range val {
			fmt.Println(val2)
		}
	}

	/**
	数组复制，数组是值类型,这个时候，arr2和arr3指向不同的内存空间
	*/
	var arr2 = [3]int{1, 2, 3}
	arr3 := arr2
	fmt.Println(arr3)
	arr3[0] = 100
	fmt.Println(arr2)
	fmt.Println(arr3)

}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
