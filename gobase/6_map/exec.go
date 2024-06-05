package mapZhang

import (
	"fmt"
)

func Exec() {

	fmt.Println(`
	1.【map】
		(1).格式：
			map[KeyType]ValueType

			其中，
				KeyType:表示键的类型。
				ValueType:表示键对应的值的类型。

		(2).声明：
			var 变量名  map[KeyType]ValueType

		(3).初始化：
			1️⃣map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
				变量名 := make(map[KeyType]ValueType, [cap])
		
					举个例子：
						scoreMap := make(map[string]int, 8)
						scoreMap["张三"] = 90
						scoreMap["小明"] = 100
						fmt.Println(scoreMap)
						fmt.Println(scoreMap["小明"])
						fmt.Printf("type of a:%T\n", scoreMap)
		
					输出：
						map[小明:100 张三:90]
						100
						type of a:map[string]int
			
			2️⃣map也支持在声明的时候填充元素。
				变量名:= map[KeyType]ValueType{"key1":"value1" , .....}

					举个例子：
						userInfo := map[string]string{
							"username": "沙河小王子",
							"password": "123456",
						}

		(4).遍历：
			Go语言中使用for range遍历map。

				举个例子：
					scoreMap := make(map[string]int)
					scoreMap["张三"] = 90
					scoreMap["小明"] = 100
					scoreMap["娜扎"] = 60
					for k, v := range scoreMap {
						fmt.Println(k, v)
					}

				但我们只想遍历key的时候，可以按下面的写法：
					scoreMap := make(map[string]int)
					scoreMap["张三"] = 90
					scoreMap["小明"] = 100
					scoreMap["娜扎"] = 60
					for k := range scoreMap {
						fmt.Println(k)
					}

			注意： 遍历map时的元素顺序与添加键值对的顺序无关。	
		
		(5).常规操作：
			1️⃣判断某个键是否存在。
				value, ok := map[key]

				举个例子：
					scoreMap := make(map[string]int)
					scoreMap["张三"] = 90
					scoreMap["小明"] = 100
					// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
					v, ok := scoreMap["张三"]
					if ok {
						fmt.Println(v)
					} else {
						fmt.Println("查无此人")
					}

			2️⃣删除键值对
				delete(map, key)
					其中，
						map:表示要删除键值对的map
						key:表示要删除的键值对的键

					举个例子：
						scoreMap := make(map[string]int)
						scoreMap["张三"] = 90
						scoreMap["小明"] = 100
						scoreMap["娜扎"] = 60
						delete(scoreMap, "小明")//将小明:100从map中删除
						for k,v := range scoreMap{
							fmt.Println(k, v)
						}
	
		(6).两个特殊类型
			1️⃣元素为map类型的切片
				举个例子：
					var mapSlice = make([]map[string]string, 3)
			2️⃣值为切片类型的map
				举个例子：
					var sliceMap = make(map[string][]string, 3)
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	/**
	map定义语法：map[KeyType]ValueType
	创建map：make(map[KeyType]ValueType, [cap]) 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	*/
	scoreMap := make(map[string]int, 8)
	scoreMap["苏州"] = 8
	scoreMap["南京"] = 10
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["苏州"])
	fmt.Println("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "zhangp",
		"password": "1234",
	}
	fmt.Println(userInfo)

	/**
	判断某个键是否存在  格式：【value, ok := map[key]】
	*/

	value, ok := userInfo["zhangp"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}

	/**
	map的遍历
	*/
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	for k := range scoreMap {
		fmt.Println(k)
	}
	for k, _ := range scoreMap {
		fmt.Println(k)
	}
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	/**
	delete删除键值对
	*/
	delete(scoreMap, "苏州")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	/**
	元素为map类型的切片
	*/
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("after init")

	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "zhangp"
	mapSlice[0]["pass"] = "123456"
	mapSlice[0]["addr"] = "成都"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	/**
	值为切片类型的map
	*/
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")

	key := "南京"
	val, ok := sliceMap[key]
	if !ok {
		val = make([]string, 0, 2)
	}
	val = append(val, "苏州", "杭州")
	sliceMap[key] = val
	fmt.Println(sliceMap)
}
