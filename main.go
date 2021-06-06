package main

import (
	"chia_api/server"
)

func main() {

	//var (
	//	DB_USER        = "talend'"
	//	DB_PASSWORD    = "talend"
	//	DB_DRIVER      = "mysql"
	//	DB_HOST        = "127.0.0.1:3306"
	//	DB_DEBUG	   = false
	//)

	server.GetServer().Run()
}
