package handlers

import (
	"chia_api/data"
	"encoding/json"
	"net/http"
)

//Get basic pool info
func PoolStat() (int, string) {
	u, err := data.GetPoolInfo()
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)

	return http.StatusOK, string(d)
}