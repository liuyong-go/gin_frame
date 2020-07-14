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
	ln, err := net.Listen("tcp4", s.Addr)
	if err != nil {
		log.Printf("Server err: %v", err)
	}
	s.Serve(ln)
}
