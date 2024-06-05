package timeZhang

import "fmt"

func Exec() {
	fmt.Printf(`
	1.【time】
		(1).时间类型
			
			举个例子：
				// timeDemo 时间对象的年月日时分秒
				func timeDemo() {
					now := time.Now() // 获取当前时间
					fmt.Printf("current time:%v\n", now)
				
					year := now.Year()     // 年
					month := now.Month()   // 月
					day := now.Day()       // 日
					hour := now.Hour()     // 小时
					minute := now.Minute() // 分钟
					second := now.Second() // 秒
					fmt.Println(year, month, day, hour, minute, second)
				}

		(2).Location和time zone
		(3).定时器
		(4).时间间隔
		(5).时间操作
		(6).时间格式化
		(7).时间字符串解析
	`)

	fmt.Println("============================================================以下是例子=========================================================")

}
