package main

import (
	"fmt"
	"gin_frame/config"
	"gin_frame/routers"
	"gin_frame/services"
	"log"
	"syscall"

	"github.com/fvbock/endless"
)

// func main_bak() {
// 	s := &http.Server{
// 		Addr:           fmt.Sprintf(":%d", config.LoadConfig().Server.HttpPort),
// 		Handler:        routers.InitRouter(),
// 		ReadTimeout:    config.LoadConfig().Server.ReadTimeout,
// 		WriteTimeout:   config.LoadConfig().Server.WriteTimeout,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	ln, err := net.Listen("tcp4", s.Addr)
// 	if err != nil {
// 		log.Printf("Server err: %v", err)
// 	}
// 	s.Serve(ln)
// }
//kill -1 pid 热重启 kill -1 $(lsof -i:8080 |awk '{print $2}' | tail -n 1)
func main() {
	var url = "https://github.com/features/project-management/"
	var serviceShortUrl services.ShortUrl
	var shortUrl = serviceShortUrl.GetShortUrl(url)
	fmt.Println(shortUrl)
}
func main_bak() {
	endless.DefaultReadTimeOut = config.LoadConfig().Server.ReadTimeout
	endless.DefaultWriteTimeOut = config.LoadConfig().Server.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("0.0.0.0:%d", config.LoadConfig().Server.HttpPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
