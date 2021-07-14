package handlers

import (
	"chia_api/data"
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/nickname32/discordhook"
	"math"
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

// GetpointTotal from pool
func GetTotalPoint() (int, string) {
	u, err := data.GetTotalPoint()
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
		interval := int64(3600 / soloFarmer.Point)
		farmer, _ := data.GetFarmer(soloFarmer.LauncherId)
		nbrShare := float64(farmer.Difficulty) * 0.001 * soloFarmer.Point
		partPartial, _ := data.GetFloatingPartial(farmer.LauncherId)
		nbrShare += partPartial.PartialPart
		partial, rest := math.Modf(nbrShare)
		f := data.NewFloatingPartial(farmer.LauncherId, rest)
		f.Save()
		err := data.UpdateFloatingPartial(f.LauncherId, f.PartialPart)
		fmt.Printf("launcher_id: %s nbrshare: %f partial: %f rest: %f with diff: %d", f.LauncherId, nbrShare, partial, rest, farmer.Difficulty)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		if farmer != nil {
			farmer.Points += soloFarmer.Point
			err := data.UpdateFarmerPoint(farmer)
			if err != nil {
				return http.StatusServiceUnavailable, err.Error()
			}
			for i := 0; i < int(partial); i++ {
				p := data.NewPartial(soloFarmer.LauncherId, t+interval, farmer.Difficulty)
				err := p.AddSoloPartial()
				if err != nil {
					return http.StatusServiceUnavailable, err.Error()
				}
				t += interval
			}
		}
	}
	return http.StatusNoContent, "ok"
}
