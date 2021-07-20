package handlers

import (
	"github.com/go-martini/martini"
	"github.com/nickname32/discordhook"
	"net/http"
)

// PostNewBlock post new block on discord
func PostNewBlock(params martini.Params) (int, string) {

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
