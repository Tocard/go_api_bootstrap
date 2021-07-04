package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/nickname32/discordhook"
	"net/http"
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
func PostPartialDiscord(params martini.Params) (int, string) {
	u, err := data.GetPartial(params["launcher_id"])

	wa, err := discordhook.NewWebhookAPI(861291081143681074, "dKHd1iYI71H0rc1rPM1vBNPawdE_uhodXSqKLNDb53wYXP_Y-EcR3zihdjKo3ullMEWX", true, nil)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	_, err = wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Title: "Un nouveau Partial soumis",
			},
		},
	}, nil, "")
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
