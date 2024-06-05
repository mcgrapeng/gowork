package structZhang

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type NewInt int
type MyInt = int

func Exec() {
	fmt.Println(`
	1.【类型别名和自定义类型】
		(1).自定义类型：
			在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。
			自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。

			举个例子：
				//基于基本类型自定义新类型
				type MyInt int
			
				//基于struct自定义新类型，其实就是结构体
				type person struct {
					name string
					city string
					age  int8
				}

		(2).类型别名：
			类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
			定义格式：type TypeAlias = Type
			
			举个例子：
				type byte = uint8
				type rune = int32

		(3).区别：
			//类型定义
			type NewInt int
			//类型别名
			type MyInt = int

			func main() {
				var a NewInt
				var b MyInt
				
				fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
				fmt.Printf("type of b:%T\n", b) //type of b:int
			}

		(4).总结：
			结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。
			MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	`)

	fmt.Printf(`
	2.【结构体】
		(1).定义：
			1️⃣标准格式：
				type 类型名 struct {
					字段名 字段类型
					字段名 字段类型
					…
				}
			
			2️⃣简化写法（同样类型的字段也可以写在一行）：
				type 类型名 struct {
					字段名1,字段名2 字段类型
					字段名2 字段类型
					…
				}
			
			其中：
				类型名：标识自定义结构体的名称，在同一个包内不能重复。
				字段名：表示结构体字段名。结构体中的字段名必须唯一。
				字段类型：表示结构体字段的具体类型。

			举个例子：
				type person struct {
					name string
					city string
					age  int8
				}

				type person1 struct {
					name, city string
					age        int8
				}

		(2).种类：
			1️⃣标准结构体
			2️⃣匿名结构体

		(3).声明：
				结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

				标准结构体声明格式：
					var 结构体名称 结构体类型

					举个例子：
						type person struct {
							name string
							city string
							age  int8
						}
		
						var p1 person

				匿名结构体声明格式：
					var 类型名称 struct{字段名 字段类型; 字段名 字段类型....}

					举个例子：
						var user struct{Name string; Age int}

		(4).实例化：
			只有当结构体实例化时，才会真正地分配内存。

			1️⃣实例化值类型结构体：
				var 结构体名称 结构体类型

				举个例子：
					var p1 person
					fmt.Printf("p1=%v\n", p1)  //p1={  0}
					fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"", city:"", age:0}

			2️⃣实例化指针类型结构体：
				方法一：var 结构体名称 = new(结构体类型)

					举个例子：
						var p2 = new(person)
						fmt.Printf("%T\n", p2)     //*main.person
						fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

				方法二：结构体名称 := &结构体类型{}
					
					举个例子：
						p3 := &person{}
						fmt.Printf("%T\n", p3)     //*main.person
						fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}

		(5).初始化：
			没有初始化的结构体，其成员变量都是对应其类型的零值。
			
				举个例子：
					var p4 person
					fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}

			1️⃣使用键值对初始化：
				使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

				举个例子：
					p5 := person{
						name: "zhangp",
						city: "北京",
						age:  18,
					}
					fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"zhangp", city:"北京", age:18}
	
					p6 := &person{
						name: "zhangp",
						city: "北京",
						age:  18,
					}
					fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"zhangp", city:"北京", age:18}
	
					p7 := &person{
						city: "北京",
					}
					fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}


			2️⃣使用值的列表初始化：
				初始化结构体的时候可以简写，也就是初始化的时候不写键

				举个例子：
					p8 := &person{
						"zhangp",
						"北京",
						28,
					}
					fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"zhangp", city:"北京", age:28}
					
			⚠️使用这种格式的注意事项：
				1.必须初始化结构体的所有字段。
				2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
				3.该方式不能和键值初始化方式混用。

		(6).内存布局：
			结构体占用一块连续的内存。

		(7).空结构体：
			空结构体是不占用空间的。

				举个例子：
					var v struct{}
					fmt.Println(unsafe.Sizeof(v))  // 0

		(8).嵌套：
			1️⃣定义：
			一个结构体中可以嵌套包含另一个结构体或结构体指针。

				举个例子：
					type Address struct {
						Province string
						City     string
					}
				
					type User struct {
						Name    string
						Gender  string
						Address Address
					}
					
					user1 := User{
						Name:   "小王子",
						Gender: "男",
						Address: Address{
							Province: "山东",
							City:     "威海",
						},
					}
					fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

			2️⃣字段冲突：
			嵌套结构体内部可能存在相同的字段名。
			在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。

				举个例子：
					type Address struct {
						Province   string
						City       string
						CreateTime string
					}
				
					type Email struct {
						Account    string
						CreateTime string
					}
	
					type User struct {
						Name   string
						Gender string
						Address
						Email
					}
					
					var user3 User
					user3.Name = "zhangp"
					user3.Gender = "男"
					user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
					user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime


		(9).继承：
			Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
			通过组合的方式实现继承。

				举个例子：
					type Animal struct {
						name string
					}
					
					func (a *Animal) move() {
						fmt.Printf("%s会动！\n", a.name)
					}

					type Dog struct {
						Feet    int8
						*Animal //通过嵌套匿名结构体实现继承
					}

					func (d *Dog) wang() {
						fmt.Printf("%s会汪汪汪~\n", d.name)
					}

		(10).可见性：
			结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

		(11).序列化：
			
			举个例子：
				c := &Class{
					Title:    "101",
					Students: make([]Student, 0, 200),
				}
				//JSON序列化：结构体-->JSON格式的字符串
				data, err := json.Marshal(c)
	
				//JSON反序列化：JSON格式的字符串-->结构体
				str = “json串”
				c1 := &Class{}
				err = json.Unmarshal([]byte(str), c1)

		(12).标签：
			1️⃣定义：
				Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来。
	`)

	fmt.Printf(`
	3.【构造函数】
		(1).定义：
			Go语言的结构体没有构造函数，我们可以自己实现。因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
			这种情况下建议构造函数返回的是结构体指针类型；除此之外可以返回结构体值类型。

			举个例子：
				func newPerson(name, city string, age int8) *person {
					return &person{
						name: name,
						city: city,
						age:  age,
					}
				}

		(2).调用：
			变量名:= 构造函数名称{变量值1,变量值2....}

			举个例子：
				p9 := newPerson("zhangp", "成都", 90)
				fmt.Printf("%#v\n", p9) //&main.person{name:"zhangp", city:"成都", age:90}
	`)

	fmt.Printf(`
	4.【方法】
		(1).定义：
			Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
			接收者的概念就类似于其他语言中的this或者 self。

			func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
				函数体
			}

			其中：
				1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
							例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
				2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
				3.方法名、参数列表、返回参数：具体格式与函数定义相同。

			举个例子：
				func (p Person) Dream() {
					fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
				}

		(2).接收者：
			1️⃣值类型接收者
				当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
				在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

				举个例子：
					// SetAge2 设置p的年龄
					// 使用值接收者
					func (p Person) SetAge2(newAge int8) {
						p.age = newAge
					}
	
					p1 := NewPerson("小王子", 25)
					p1.Dream()
					fmt.Println(p1.age) // 25
					p1.SetAge2(30) // (*p1).SetAge2(30)
					fmt.Println(p1.age) // 25

			2️⃣指针类型接收者
				指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
				这种方式就十分接近于其他语言中面向对象中的this或者self。

				举个例子：
					// SetAge 设置p的年龄
					// 使用指针接收者
					func (p *Person) SetAge(newAge int8) {
						p.age = newAge
					}
	
				调用：
					p1 := NewPerson("小王子", 25)
					fmt.Println(p1.age) // 25
					p1.SetAge(30)
					fmt.Println(p1.age) // 30

			3️⃣什么时候应该使用指针类型接收者
				1.需要修改接收者中的值。
				2.接收者是拷贝代价比较大的大对象。
				3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

		(3).区别：
			函数不属于任何类型，方法属于特定的类型。

	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	自定义类型和类型别名
	*/
	fmt.Println(`
		将MyInt定义为int类型
		type MyInt int
		
		TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型
		type TypeAlias = Type
		type byte = uint8
		type rune = int32
	`)

	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	fmt.Println(`
		结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
		b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	`)

	/**
	结构体定义
	格式：
		type 类型名 struct {
			字段名 字段类型
			字段名 字段类型
			…
		}

	其中：
		类型名：标识自定义结构体的名称，在同一个包内不能重复。
		字段名：表示结构体字段名。结构体中的字段名必须唯一。
		字段类型：表示结构体字段的具体类型。
	*/
	fmt.Println(`
		type person struct {
			name string
			city string
			age  int8
		}

		type person1 struct {
			name, city string //同样类型的字段也可以写在一行
			age        int8
		}
	`)

	/**
	结构体实例化
	*/
	var p1 person
	//p1.name = "张鹏"
	//p1.city = "成都"
	//p1.age = 19

	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	/**
	匿名结构体
	*/
	var user struct {
		Name string
		Age  int
	}
	user.Name = "张三丰"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	/**
	指针类型结构体实例化 - new
	*/
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	p2.name = "张北海"
	p2.age = 19
	p2.city = "北海"
	fmt.Printf("p2=%#v\n", p2)

	/**
	取结构体的地址实例化
	使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	*/
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)

	p3.name = "项羽"
	p3.age = 30
	p3.city = "苏州"
	fmt.Printf("p3=%#v\n", p3)

	fmt.Println(`
		p3.name = "张鹏"  其实在底层是(*p3).name = "张鹏"，这是Go语言帮我们实现的语法糖。
	`)

	/**
	结构体初始化
	没有初始化的结构体，其成员变量都是对应其类型的零值。
	*/
	var p4 person
	fmt.Printf("p4=%#v\n", p4)

	/**
	使用键值对初始化
	*/
	//使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p5 := person{
		name: "zhang",
		city: "南京",
		age:  19,
	}
	fmt.Printf("p5=%#v\n", p5)

	//也可以对结构体指针进行键值对初始化
	p6 := &person{
		name: "zhangp",
		age:  19,
		city: "苏州",
	}
	fmt.Printf("p6=%#v\n", p6)

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		city: "宁波",
	}
	fmt.Printf("p7=%#v\n", p7)

	//使用值的列表初始化
	p8 := &person{
		"zhang",
		"苏州",
		28,
	}
	fmt.Printf("p8=%#v\n", p8)
	fmt.Println(`
		使用这种格式初始化时，需要注意：
			1.必须初始化结构体的所有字段。
			2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
			3.该方式不能和键值初始化方式混用。
	`)

	/**
	结构体内存布局

	结构体占用一块连续的内存。
	*/
	n := test{1, 2, 3, 4}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	/**
	空结构体
	空结构体是不占用空间的。
	*/
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) //0

	/**
	构造函数

	Go语言的结构体没有构造函数，我们可以自己实现。
	例如，下方的代码就实现了一个person的构造函数。
	因为struct是值类型，如果结构体比较复杂的话，
	值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
	*/
	p9 := newPersion("zhang", "南京", 18)
	fmt.Printf("%#v\n", p9)

	/**
	方法和接收者
	格式：
	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	    函数体
	}

	其中，
	接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
	接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
	方法名、参数列表、返回参数：具体格式与函数定义相同。

	举个例子：
		func (p person) sayHi(a int) int8  {
			return 0
		}

		func (p *person) sayHello(a *int)  int {
			return 0
		}
	*/
	p10 := newPersion("zhang", "金陵", 10)
	p10.dream()

	/**
	指针类型接收者
	*/
	fmt.Println(`
		func (p *person) setAge(newAge int8)  {
			p.age = newAge
		}
	`)

	p10.setAge(100)
	fmt.Println(p10.age)

	/**
	值类型接收者
	*/
	fmt.Println(`
		func (p person) setAge2(newAge int8) {
			p.age = newAge
		}
	`)
	p10.setAge2(1)
	fmt.Println(p10.age)

	fmt.Println(`
		什么时候应该使用指针类型接收者
			1.需要修改接收者中的值
			2.接收者是拷贝代价比较大的大对象
			3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	`)

	/**
	为任意类型添加方法
	在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
	举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
	*/
	fmt.Println(`
		type NewInt int
		
		func (m NewInt) sayNo()  {
	
		}
	`)

	var m NewInt
	fmt.Println(m)
	m.sayNo()
	m = 1000
	fmt.Printf("%#v %T\n", m, m)

	fmt.Println(`
		注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
	`)

	/**
	结构体匿名字段

	type Person struct {
		string
		int
	}

	**注意：**这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，
	结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	*/
	p11 := Person{
		"zhang",
		18,
	}
	fmt.Printf("%#v\n", p11)
	fmt.Println(p11.int, p11.string)

	/**
	嵌套结构体
	*/
	user1 := User{
		Name: "张",
		Addr: Address{
			Province: "四川",
			City:     "成都",
		},
	}
	fmt.Printf("user=%#v\n", user1)

	fmt.Println(`
		嵌套结构体匿名字段
		type User struct {
			Name    string
			Gender  string
			Address //匿名字段
		}
	`)

	var user2 User1
	user2.Name = "张北海"
	user2.Address.Province = "江苏"
	user2.Address.City = "南京"
	user2.City = "苏州" //匿名字段可以省略
	fmt.Printf("user=%#v\n", user2)

	fmt.Println(`
		注意：当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	`)

	/**
	嵌套结构体的字段名冲突

	嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。
	*/
	var std Student
	std.Address.CreateTime = "2023"
	std.Email.CreateTime = "2024"
	fmt.Printf("std=%#v\n", std)

	/**
	结构体的“继承”

	type Order struct {
		name string
	}

	type OrderView struct {
		desc   string
		*Order //通过嵌套匿名结构体实现继承
	}

	func (a *Order) desc() {
		fmt.Println("被夹了")
	}

	func (d *OrderView) remark() {
		fmt.Println("夹脑袋")
	}
	*/
	ord := &OrderView{
		desc: "司马夹",
		Order: &Order{
			name: "夹夹夹",
		},
	}
	ord.remark()
	ord.Order.desc()

	//不可以这么写：
	//var ord1 Order
	//ord1 = OrderView{
	//	desc: "司马夹",
	//	Order: &Order{
	//		name: "夹夹夹",
	//	},
	//}
	//fmt.Println(ord1)

	/**
	结构体字段的可见性
	*/
	fmt.Println(`
		结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
	`)

	/**
	结构体与JSON序列化
	*/
	c := &Class{
		Title:    "3班",
		Students: make([]Student, 0, 200),
	}

	for i := 0; i < 5; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json Marshal failed")
		return
	}

	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101"
,"Students":[{"Gender":"男","Name":"stu00"},{"Gender":"男","Name":"stu01"},{"Gender":"男","Name":"stu02"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json Unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	/**
	结构体标签（Tag）

	Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`



	type Student struct {
		Name   string `json:"name"` //通过指定tag实现json序列化该字段时的key
		Gender string //json序列化是默认使用字段名作为key
		age    int //私有不能被json包访问
		Address
		Email
	}
	*/

	s1 := Student{
		Name:   "zhang",
		Gender: "男",
	}
	data1, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json Marshal failed")
		return
	}

	fmt.Printf("json str:%s\n", data1)

	/**
	因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。
	*/

	c2 := Class{}
	dreams := []string{"吃饭", "睡觉"}
	//c2.SetDreams(dreams)
	//fmt.Println(c2.dreams)
	//
	//dreams[1] = "不睡觉"
	//fmt.Println(c2.dreams)

	c2.SetDream1(dreams)
	fmt.Println(c2.dreams)

	dreams[1] = "不睡觉"
	fmt.Println(c2.dreams)

}

type person1 struct {
	name, city string //同样类型的字段也可以写在一行
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func newPersion(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

type person struct {
	name string
	city string
	age  int8
}

func (p person) dream() {
	fmt.Println("月入百万")
}

func (p person) sayHi(a int) int8 {
	return 0
}

func (p *person) sayHello(a *int) int {
	return 0
}

func (p *person) setAge(newAge int8) {
	p.age = newAge
}

func (p person) setAge2(newAge int8) {
	p.age = newAge
}

func (m NewInt) sayNo() {

}

type Person struct {
	string
	int
}

type Address struct {
	Province   string
	City       string
	CreateTime string
}

type User struct {
	Name string
	Addr Address
}

type User1 struct {
	Name string
	Address
}
type Email struct {
	Account    string
	CreateTime string
}

type Student struct {
	Name   string `json:"name1"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	age    int    //私有不能被json包访问
	Address
	Email
}

type Order struct {
	name string
}

type OrderView struct {
	desc   string
	*Order //通过嵌套匿名结构体实现继承
}

func (o *Order) desc() {
	fmt.Println("被夹了")
}

func (o *OrderView) remark() {
	fmt.Println("夹脑袋")
}

type Class struct {
	Title    string
	Students []Student
	dreams   []string
}

func (c *Class) SetDreams(dreams []string) {
	c.dreams = dreams
}

func (c *Class) SetDream1(dreams []string) {
	c.dreams = make([]string, len(dreams))
	copy(c.dreams, dreams)
}
