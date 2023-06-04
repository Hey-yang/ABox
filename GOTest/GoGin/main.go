package main

import goginserver "GoGin/GoGinServer"

func main() {
	// 设置路由信息
	r := goginserver.GetGinServerSi()
	// 启动服务器并监听 8080 端口
	r.Run(":8082")
}
