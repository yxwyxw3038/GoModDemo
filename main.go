package main

import (
	"fmt"
	. "GoModDemo/router"
)



func main() {
	fmt.Printf("启动服务")
	router := InitRouter()
	router.Run(":8080")
}
