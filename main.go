package main

import (
	"fmt"
	"gin_frame/config"
	"gin_frame/routers"
	"log"
	"net"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.LoadConfig().Server.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    config.LoadConfig().Server.ReadTimeout,
		WriteTimeout:   config.LoadConfig().Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Print(config.LoadConfig().Server.ReadTimeout)
	ln, err := net.Listen("tcp4", s.Addr)
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	s.Serve(ln)
}
func main_bak() {
	//监听协议
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/user/login", UserLoginHandler)
	//监听服务
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.LoadConfig().Server.HttpPort), nil)

	if err != nil {
		fmt.Println("服务器错误")
	}
}
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w, "HelloWorld!")
}

func UserLoginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Fprintf(response, "Login Success")
}
