package handlers

import (
	"chia_api/data"
	"chia_api/redis"
	"encoding/json"
	"net/http"
)

//Get basic pool info
func PoolStat() (int, string) {

	redisPoolInfos := redis.GetFromToRedis(0, "poolinfos")
	if redisPoolInfos != "" {
		return http.StatusOK, string(redisPoolInfos)
	}else{
		u, err := data.GetPoolInfo()
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		d, _ := json.Marshal(u)

		redis.WriteToRedis(0, "poolinfos", string(d))

		return http.StatusOK, string(d)
	}
}