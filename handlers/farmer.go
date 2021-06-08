package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/go-martini/martini"
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

// GetFarmersCount get farmer.
func GetFarmersCount() (int, string) {

	u, err := data.GetFarmersCount()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
