package handlers

import (
	"chia_api/data"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/nickname32/discordhook"
	"net/http"
	"time"
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

// GetNetSpaceByLauncherId get netspace estimation from launcher_id.
func GetNetSpaceByLauncherId(params martini.Params) (int, string) {
	u, err := data.GetNetSpaceByLauncherId(params["launcher_id"])
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
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

// GetPartials get Partial.
func GetPartials() (int, string) {

	u, err := data.GetPartials()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}

// PostPartialDiscord post new partial on discord
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

// PostPartialSoloplot post solo plot on bdd share
func PostPartialSoloplot(r *http.Request) (int, string) {

	res := []data.SolotPlot{}
	db := data.GetConn()
	defer db.Close()
	dec := json.NewDecoder(r.Body)
	dec.Decode(&res)
	t := time.Now().Unix() - 3600
	for _, soloFarmer := range res {
		interval := 3600 / soloFarmer.Point
		farmer, _ := data.GetFarmer(soloFarmer.LauncherId)
		if farmer != nil {
			fmt.Println(farmer.Points, soloFarmer.Point)
			farmer.Points += soloFarmer.Point
			err := data.UpdateFarmerPoint(farmer)
			if err != nil {
				return http.StatusServiceUnavailable, err.Error()
			}
			for i := 0; i < soloFarmer.Point; i++ {
				p := data.NewPArtial(soloFarmer.LauncherId, t+int64(interval), 1)
				err := p.AddSoloPartial()
				if err != nil {
					return http.StatusServiceUnavailable, err.Error()
				}
				t += int64(interval)
			}
		}
	}
	return http.StatusNoContent, "ok"
}
