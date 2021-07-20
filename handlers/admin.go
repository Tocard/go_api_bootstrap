package handlers

import (
	"chia_api/data"
	"github.com/go-martini/martini"
	"net/http"
)

// GenerateAdmin generated admin user on given launcher id with auto random password
func GenerateAdmin(params martini.Params) (int, string) {

	u := data.NewAdmin(params["launcher_id"])
	err := u.Save()
	if err != nil {
		return http.StatusServiceUnavailable, err.Error()
	}
	return http.StatusNoContent, "ok"
}
