package handlers

import (
	"chia_api/data"
	"encoding/json"
	"net/http"
)

func MiningPoolStat() (int, string) {
	u, err := data.GetMiningStatPool()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}