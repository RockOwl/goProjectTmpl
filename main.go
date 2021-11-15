package main

import (
	"fmt"
	"goProjectTmpl/config"
)

const (
	Path = "/Users/xxxx/CodeSpace/GoSpace/goProjectTmpl"
)

func main() {
	fmt.Println("hello world")

	//// 加载本地配置文件: config.WithProvider("file")
	//config.Load(Path+"/app.yaml", config.WithCodec("yaml"), config.WithProvider("file"))

	// 默认的DataProvider是使用本地文件
	cfg, _ := config.Load(Path+"/app.yaml", config.WithCodec("yaml"))

	// 读取bool类型配置
	value1 := cfg.GetBool("server.debug", false)
	fmt.Printf("Get1 %v\n", value1)

	// 读取String类型配置
	value2 := cfg.GetString("server.app", "default")
	fmt.Printf("Get2 %v\n", value2)

	value3 := cfg.GetString("server.admin.ip", "default")
	fmt.Printf("Get3 %v\n", value3)

}
