package handlers

import (
	"chia_api/data"
	"encoding/json"
	"net/http"
)

func Healthz() (int, string) {
	return http.StatusOK, "OK"
}

func Name() (int, string) {
	return http.StatusOK, data.GetName()
}

func Version() (int, string) {
	return http.StatusOK, "0.0.0"
}

func Fees() (int, string) {
	return http.StatusOK, data.GetFees()
}

func MiningPoolStat() (int, string) {
	u, err := data.GetMiningStatPool()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
