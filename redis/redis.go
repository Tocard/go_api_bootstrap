package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	Host     = "localhost"
	Port     = 6379
	Password = ""
	Clis     = make(map[int]*redis.Client, 2)
	Lifetime = 60
)

func init() {

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
	if err == redis.Nil || err != nil {
		return ""
	}
	return val
}
