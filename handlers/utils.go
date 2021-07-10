package handlers

import (
	"chia_api/data"
	"chia_api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Login user and return JWT token.
func Login(r *http.Request) (int, string) {

	where := data.PoolAdmin{}
	res := []data.PoolAdmin{}
	db := data.GetConn()
	defer db.Close()
	dec := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	dec.Decode(&where)

	if len(strings.TrimSpace(where.LauncherId)) == 0 {
		return http.StatusBadRequest, "You should give launcher_id"
	}
	if len(strings.TrimSpace(where.Password)) == 0 {
		return http.StatusBadRequest, "You should give password"
	}

	db.Model(&data.PoolAdmin{}).Where("launcher_id = ?", where.LauncherId).Find(&res)
	if len(res) > 0 {
		// Keep password
		pwd := where.Password
		//err := bcrypt.CompareHashAndPassword([]byte(res[0].Password), pwd) Idem, une fois le password stocké hashé
		if res[0].Password != pwd {
			return http.StatusUnauthorized, "Bad password"
		}

		token, _ := utils.GenerateToken(int(res[0].ID))
		return http.StatusOK, token

	}

	return http.StatusUnauthorized, "Bad login"
}

func Healthz() (int, string) {
	return http.StatusOK, "OK"
}

func Name() (int, string) {
	return http.StatusOK, data.GetName()
}

func Version() (int, string) {
	return http.StatusOK, "0.0.0"
}


