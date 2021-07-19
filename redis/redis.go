package redis

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	Host     = "should_never_be_seen"
	Port     = 6379
	Password = ""
	Clis     = make(map[int]*redis.Client, 2)
	Lifetime = 60
)

type Config struct {
	RedisHost     string `yaml:"host"`
	RedisPort     int    `yaml:"port"`
	RedisPassword string `yaml:"password"`
	RedisLifetime int    `yaml:"lifetime"`
}

func getRedisConf() {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile("redis_config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	Host = c.RedisHost
	Password = c.RedisPassword
	Port = c.RedisPort
	Lifetime = c.RedisLifetime
	fmt.Println(Host, c)
}

func init() {

	getRedisConf()
	connect(0)
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				if err := Clis[0].Ping(); err.Err() != nil {
					fmt.Println(err.Err())
					connect(0)
				}
			}
		}
	}()
}
func connect(DbNum int) {
	fmt.Println(Host, Port)
	Clis[DbNum] = redis.NewClient(&redis.Options{
		Addr:        Host + ":" + strconv.Itoa(Port),
		Password:    Password, // no password set
		DB:          DbNum,
		IdleTimeout: 10 * time.Second,
		// use default DB
	})

}

func WriteToRedis(DbNum int, key, value string) {
	Clis[DbNum].Set(key, value, time.Duration(Lifetime)*time.Second).Err() //TODO: Ajouter du logging et gestion d'erreur
}

func GetFromToRedis(DbNum int, key string) string {
	val, err := Clis[DbNum].Get(key).Result()
	fmt.Printf("result from redis %s\n", val)
	if err == redis.Nil || err != nil {
		return ""
	}
	return val
}
