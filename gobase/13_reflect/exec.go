package reflectZhang

import (
	"fmt"
	"reflect"
)

type myInt int64

func Exec() {

	fmt.Printf(`
	1.【reflect】
		(1).定义：
			在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，
			并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。

		(2).API：
			1️⃣TypeOf
				1.定义：
					在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

					举个例子：
						func reflectType(x interface{}) {
							v := reflect.TypeOf(x)
							fmt.Printf("type:%v\n", v)
						}
		
						func main() {
							var a float32 = 3.14
							reflectType(a) // type:float32
							var b int64 = 100
							reflectType(b) // type:int64
						}
				
				2.结构：
					在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
					因为在Go语言中我们可以使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
					但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。
				
					举个例子：
						type myInt int64

						func reflectType(x interface{}) {
							t := reflect.TypeOf(x)
							fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
						}
						
						func main() {
							var a *float32 // 指针
							var b myInt    // 自定义类型
							var c rune     // 类型别名
							reflectType(a) // type: kind:ptr
							reflectType(b) // type:myInt kind:int64
							reflectType(c) // type:int32 kind:int32
						
							type person struct {
								name string
								age  int
							}
							type book struct{ title string }
							var d = person{
								name: "沙河小王子",
								age:  18,
							}
							var e = book{title: "《跟小王子学Go语言》"}
							reflectType(d) // type:person kind:struct
							reflectType(e) // type:book kind:struct
						}

				3.Kind类型：
					type Kind uint
					const (
						Invalid Kind = iota  // 非法类型
						Bool                 // 布尔型
						Int                  // 有符号整型
						Int8                 // 有符号8位整型
						Int16                // 有符号16位整型
						Int32                // 有符号32位整型
						Int64                // 有符号64位整型
						Uint                 // 无符号整型
						Uint8                // 无符号8位整型
						Uint16               // 无符号16位整型
						Uint32               // 无符号32位整型
						Uint64               // 无符号64位整型
						Uintptr              // 指针
						Float32              // 单精度浮点数
						Float64              // 双精度浮点数
						Complex64            // 64位复数类型
						Complex128           // 128位复数类型
						Array                // 数组
						Chan                 // 通道
						Func                 // 函数
						Interface            // 接口
						Map                  // 映射
						Ptr                  // 指针
						Slice                // 切片
						String               // 字符串
						Struct               // 结构体
						UnsafePointer        // 底层指针
					)
				
				4.返回值reflect.Type类型
					(1).相关方法：
						Field(i int) StructField		根据索引，返回索引对应的结构体字段的信息。
						NumField() int					返回结构体成员字段数量。
						FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。
						FieldByIndex(index []int) StructField			多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
						FieldByNameFunc(match func(string) bool) (StructField,bool)		根据传入的匹配函数匹配需要的字段。
						NumMethod() int					返回该类型的方法集中方法的数目
						Method(int) Method				返回该类型方法集中的第i个方法
						MethodByName(string)(Method, bool)				根据方法名返回该类型方法集中的方法

				5.StructField类型
					StructField类型用来描述结构体中的一个字段的信息。

					结构如下：
						type StructField struct {
							// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
							Name    string
							PkgPath string
							Type      Type      // 字段的类型
							Tag       StructTag // 字段的标签
							Offset    uintptr   // 字段在结构体中的字节偏移量
							Index     []int     // 用于Type.FieldByIndex时的索引切片
							Anonymous bool      // 是否匿名字段
						}

			2️⃣ValueOf
				1.定义：
					reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
			
					举个例子：
						func reflectValue(x interface{}) {
							v := reflect.ValueOf(x)
							k := v.Kind()
							switch k {
							case reflect.Int64:
								// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
								fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
							case reflect.Float32:
								// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
								fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
							case reflect.Float64:
								// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
								fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
							}
						}
						func main() {
							var a float32 = 3.14
							var b int64 = 100
							reflectValue(a) // type is float32, value is 3.140000
							reflectValue(b) // type is int64, value is 100
							// 将int类型的原始值转换为reflect.Value类型
							c := reflect.ValueOf(10)
							fmt.Printf("type c :%T\n", c) // type c :reflect.Value
						}

				2.设置变量的值：
					想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
					而反射中使用专有的Elem()方法来获取指针对应的值。
			
					举个例子：
						func reflectSetValue1(x interface{}) {
							v := reflect.ValueOf(x)
							if v.Kind() == reflect.Int64 {
								v.SetInt(200) //修改的是副本
							}
						}
						func reflectSetValue2(x interface{}) {
							v := reflect.ValueOf(x)
							// 反射中使用 Elem()方法获取指针对应的值
							if v.Elem().Kind() == reflect.Int64 {
								v.Elem().SetInt(200)
							}
						}
						func main() {
							var a int64 = 100
							// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
							reflectSetValue2(&a)
							fmt.Println(a)
						}

			3️⃣isNil()和isValid()
				func (v Value) IsNil() bool
				IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

				func (v Value) IsValid() bool
				IsValid()返回v是否持有一个值。如果v是Value零值会返回假

				IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。

				举个例子：
					func main() {
						// *int类型空指针
						var a *int
						fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
						// nil值
						fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
						// 实例化一个匿名结构体
						b := struct{}{}
						// 尝试从结构体中查找"abc"字段
						fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
						// 尝试从结构体中查找"abc"方法
						fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
						// map
						c := map[string]int{}
						// 尝试从map中查找一个不存在的键
						fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
					}
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	var a *float32
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	var c myInt
	reflectType(c)
	var d rune
	reflectType(d)

	type person struct {
		name string
		age  int
	}

	type book struct {
		title string
	}

	var p = person{
		name: "张北海",
		age:  19,
	}

	var e = book{title: "go从入门到放弃"}

	reflectType(p)
	reflectType(e)

	/**
	ValueOf
	*/
	var a1 float32 = 3.14
	var b1 int64 = 100

	reflectValue(a1)
	reflectValue(b1)

	c1 := reflect.ValueOf(10)
	fmt.Printf("type c1 :%T\n", c1)

	/**
	通过反射设置变量的值
	*/
	var a2 int64 = 100
	reflectSetValue1(&a2)
	fmt.Println(a2)

	/**
	isNil()和isValid()
	IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	*/
	var a4 *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a4).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	//实例化一个匿名结构体
	b4 := struct {
	}{}

	fmt.Println("不存在的结构体成员：", reflect.ValueOf(b4).FieldByName("aaa"))
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(b4).MethodByName("aaa"))

	//实例化map
	c4 := map[string]int{}

	fmt.Println("map中不存在的键：", reflect.ValueOf(c4).MapIndex(reflect.ValueOf("南京")).IsValid())

	//reflect.TypeOf(c4).Field()
	//reflect.TypeOf(c4).FieldByName("")
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()

	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64 , value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32 , value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64 , value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
