package handlers

import (
	"chia_api/data"
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
)

// GetLastBlock return last block found from the pool
func GetLastBlock(params martini.Params) (int, string) {
	u, err := data.GetFarmer(params["launcher_id"])
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}
	d, _ := json.Marshal(u)
	return http.StatusOK, string(d)
}
