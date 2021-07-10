package handlers

import (
	"chia_api/data"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

// GenerateAdmin generated admin user on given launcher id with auto random password
func GenerateAdmin(params martini.Params) (int, string) {

	u := data.NewAdmin(params["launcher_id"])
	fmt.Println(u)
	err := u.Save()
	if err != nil {
		return http.StatusServiceUnavailable, err.Error()
	}
	return http.StatusNoContent, "ok"
}
