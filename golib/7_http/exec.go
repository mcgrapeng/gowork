package httpZhang

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Exec() {

	fmt.Printf(`
	2.【http】
		(1).定义：
			Go语言内置的net/http包十分的优秀，提供了HTTP客户端和服务端的实现。
		
		(2).API：
			1️⃣客户端
				net/http包提供了Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
				程序在使用完response后必须关闭回复的主体。

				1.Get请求
					resp, err := http.Get("https://www.liwenzhou.com/")
					if err != nil {
						fmt.Printf("get failed, err:%v\n", err)
						return
					}
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Printf("read from resp.Body failed, err:%v\n", err)
						return
					}
					fmt.Print(string(body))

				2.Get请求(带参数)
					详情请见example目录

				3.Post请求
					详情请见example目录

			2️⃣服务端
				ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。
				Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

					举个例子：
						http.Handle("/foo", fooHandler)
						http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
							fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
						})
						log.Fatal(http.ListenAndServe(":8080", nil))

				1.默认的Server：
					http.HandleFunc("/", sayHello)
					err := http.ListenAndServe(":9090", nil)
					if err != nil {
						fmt.Printf("http server failed, err:%v\n", err)
						return
					}
					
				2.自定义Server：
					s := &http.Server{
						Addr:           ":8080",
						Handler:        myHandler,
						ReadTimeout:    10 * time.Second,
						WriteTimeout:   10 * time.Second,
						MaxHeaderBytes: 1 << 20,
					}
					log.Fatal(s.ListenAndServe())	
`)

	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))

	//http.HandleFunc("/", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Printf("http server failed, err:%v\n", err)
	//	return
	//}
}
