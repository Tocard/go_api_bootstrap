package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"github.com/nickname32/discordhook"
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
        u, err := data.GetFarmer(params["launcher_id"])

        wa, err := discordhook.NewWebhookAPI(857380040492056618, "did3j-_Sv_kKCkJQ8TbvndbzvAxTVVwL3jNXUOIlMoxz5VpHtpS5V7VngbzVk4uf6Utg", true, nil)
        if err != nil {
                return http.StatusInternalServerError, err.Error()
        }

        _, err = wa.Execute(nil, &discordhook.WebhookExecuteParams{
                Content: "L'equipe s'agrandit",
                Embeds: []*discordhook.Embed{
                        {
                                Title:       "Ici la French Farmer Pool",
                                Description: "Un nouveau Miner nous rejoint sur le testnet",
                        },
                },
        }, nil, "")
        if err != nil {
                return http.StatusInternalServerError, err.Error()
        }
        d, _ := json.Marshal(u)
        return http.StatusOK, string(d)
}