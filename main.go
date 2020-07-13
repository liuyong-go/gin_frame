package main

import (
	"fmt"
	"gin_frame/config"
)

func main() {
	fmt.Print(config.LoadConfig().Database.Host)
}
