package handlers

import (
	"chia_api/data"
	"chia_api/redis"
	"chia_api/utils"
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
)

// GetNetSpaceByLauncherId get netspace estimation from launcher_id.
func GetNetSpaceByLauncherId(params martini.Params) (int, string) {
	launcherId, _ := params["launcher_id"]
	redisValue := redis.GetFromToRedis(0, launcherId)
	redisNetspace := utils.StringToFloat(redisValue)
	if redisNetspace != 0.0 {
		d, _ := json.Marshal(redisNetspace)
		return http.StatusOK, string(d)
	}
	u, err := data.GetNetSpaceByLauncherId(params["launcher_id"])
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	redis.WriteToRedis(0, launcherId, utils.FloatToString(u))
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// GetNetSpaceTotal get netspace estimation .
func GetNetSpaceTotal() (int, string) {
	u, err := data.GetNetSpaceTotal()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
