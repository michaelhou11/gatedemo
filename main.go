package main

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/api/client"
)

func main() {
	// 这里只是演示导入。如果 gatechainsdk 有公共函数和类型，你可以这样调用：
	fmt.Println("GateChain SDK Demo")

	APIToken := "your_api_token"
	EndPoint := "http://your_server_ip:port"
	client := client.NewClient(EndPoint, APIToken)
	client.Version.Version()
}
