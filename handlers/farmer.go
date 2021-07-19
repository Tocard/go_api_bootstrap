package handlers

import (
	"chia_api/data"
	"chia_api/redis"
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/nickname32/discordhook"
	"net/http"
)

// GetFarmer get farmer.
func GetFarmer(params martini.Params) (int, string) {
	u, err := data.GetFarmer(params["launcher_id"])
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// GetFarmers get farmer.
func GetFarmers() (int, string) {

	u, err := data.GetFarmers()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// GetFarmers top get farmer.
func GetTopFarmers() (int, string) {
	redisTop := redis.GetFromToRedis(0, "topFarmer")
	if redisTop != "" {
		return http.StatusOK, string(redisTop)
	}else{
		u, err := data.GetTopFarmers()
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		d, _ := json.Marshal(u)
		redis.WriteToRedis(0, "topFarmer",  string(d))
		return http.StatusOK, string(d)
	}
}

// GetFarmersCount get farmer.
func GetFarmersCount() (int, string) {

	u, err := data.GetFarmersCount()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// PostFarmerDiscord post to discord each new farmer from chia-pool
func PostFarmerDiscord(params martini.Params) (int, string) {
	u := params["launcher_id"]

	wa, err := discordhook.NewWebhookAPI(861291081143681074, "dKHd1iYI71H0rc1rPM1vBNPawdE_uhodXSqKLNDb53wYXP_Y-EcR3zihdjKo3ullMEWX", true, nil)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	_, err = wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Title:       "Ici la French Farmer Pool",
				Description: "Un nouveau Miner nous rejoint sur le testnet avec l'id " + u,
			},
		},
	}, nil, "")
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
