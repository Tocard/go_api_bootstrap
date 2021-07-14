package main

import (
	"chia_api/data"
	"chia_api/redis"
	"chia_api/server"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	RedisHost     string `yaml:"host"`
	RedisPort     int    `yaml:"port"`
	RedisPassword string `yaml:"password"`
	RedisLifetime int    `yaml:"lifetime"`
}

func getConf() {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile("redis_config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	redis.Host = c.RedisHost
	redis.Password = c.RedisPassword
	redis.Port = c.RedisPort
	redis.Lifetime = c.RedisLifetime
}

func main() {
	var bddType string
	var dbPath string
	var dbHost string
	var dbUser string
	var dbPassword string
	var dbDatabase string
	var debug bool
	var apiPort string
	var apiHost string

	flag.StringVar(&bddType, "dbmode", "sqlite3", "sqlite3 | mysql")
	flag.StringVar(&dbPath, "dbpath", "", "path to sqlite")
	flag.StringVar(&dbHost, "dbHost", "", "hosts:port of mysql instance")
	flag.StringVar(&dbUser, "dbUser", "", "username for mysql")
	flag.StringVar(&dbDatabase, "dbDatabase", "", "database name for mysql")
	flag.StringVar(&dbPassword, "dbPassword", "", "password for mysql user")
	flag.StringVar(&apiPort, "apiPort", "3000", "Api port")
	flag.StringVar(&apiHost, "apiHost", "localhost", "Api host")
	flag.BoolVar(&debug, "debug", false, "true | false")

	flag.Parse()
	data.DB_USER = dbUser
	data.DB_PASSWORD = dbPassword
	data.DB_PATH = dbPath
	data.DB_DRIVER = bddType
	data.DB_HOST = dbHost

	data.API_PORT = apiPort
	data.API_HOST = apiHost

	data.Migrate()
	server.GetServer().Run()
}
