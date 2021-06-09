package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"github.com/nickname32/discordhook"
)

// GetPartial get Partial.
func GetPartial(params martini.Params) (int, string) {
	u, err := data.GetPartial(params["launcher_id"])
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// GetPartials get Partial.
func GetPartials() (int, string) {

	u, err := data.GetPartials()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// PostPartials get Partial.
func PostPartial(params martini.Params) (int, string) {
	u, err := data.GetPartial(params["launcher_id"])

	wa, err := discordhook.NewWebhookAPI(857380040492056618, "did3j-_Sv_kKCkJQ8TbvndbzvAxTVVwL3jNXUOIlMoxz5VpHtpS5V7VngbzVk4uf6Utg", true, nil)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	_, err = wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Content: "Example text",
		Embeds: []*discordhook.Embed{
			{
				Title:       "Hi there",
				Description: "Un nouveau Miner nous rejoint sur le testnet,
			},
		},
	}, nil, "")
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}