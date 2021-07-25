package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/nickname32/discordhook"
	"net/http"
)

// PostNewBlock post new block on discord
func PostNewBlock(r *http.Request) (int, string) {
	block := data.WinBlock{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&block)
	wa, err := discordhook.NewWebhookAPI(861291081143681074, "dKHd1iYI71H0rc1rPM1vBNPawdE_uhodXSqKLNDb53wYXP_Y-EcR3zihdjKo3ullMEWX", true, nil)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	_, err = wa.Execute(nil, &discordhook.WebhookExecuteParams{
		Embeds: []*discordhook.Embed{
			{
				Title: "Un nouveau block pour la pool",
			},
		},
	}, nil, "")
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	return http.StatusOK, ""
}
