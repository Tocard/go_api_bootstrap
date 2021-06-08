package main

import (
	"chia_api/data"
	"chia_api/server"
	"flag"
)

func main() {
	var bddType string
	var dbPath string
	var dbHost string
	var dbUser string
	var dbPassword string
	var dbDatabase string
	var debug bool

	flag.StringVar(&bddType, "dbmode", "sqlite3", "sqlite3 | mysql")
	flag.StringVar(&dbPath, "dbpath", "", "path to sqlite")
	flag.StringVar(&dbHost, "dbHost", "", "hosts:port of mysql instance")
	flag.StringVar(&dbUser, "dbUser", "", "username for mysql")
	flag.StringVar(&dbDatabase, "dbDatabase", "", "database name for mysql")
	flag.StringVar(&dbPassword, "dbPassword", "", "password for mysql user")
	flag.BoolVar(&debug, "debug", false, "true | false")

	flag.Parse()
	data.DB_USER = dbUser
	data.DB_PASSWORD = dbPassword
	data.DB_PATH = dbPath
	data.DB_DRIVER = bddType
	data.DB_HOST = dbHost

	data.Migrate()
	server.GetServer().Run()
}
