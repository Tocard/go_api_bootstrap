package server

import (
	"bytes"
	"chia_api/handlers"
	"chia_api/utils"
	"encoding/json"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
	"time"
)

// Log represent a log to save in mongo database.
type Log struct {
	UserID uint
	Date   time.Time
	URL    string
	Data   interface{}
}

// LogRequest is a logging middleware to activate with martini. It gets
// request body, date and url, and set it to mongodb.
func LogRequest(r *http.Request) {

	l := Log{
		UserID: 0,
		Date:   time.Now(),
		URL:    r.URL.String(),
	}

	// keep the body content in a []byte
	b, _ := ioutil.ReadAll(r.Body)
	// rewind the body, so that json.Decoder will be able to read
	// then entire content, and we will be able to reset
	// the body with previously saved content...
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	// read the content to decode
	json.NewDecoder(r.Body).Decode(&l.Data)

	// we can now reset body for later use
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

}

func server() *martini.ClassicMartini {
	app := martini.Classic()
	app.Use(LogRequest)
	app.Use(martini.Static("assets"))

	// Allow CORS
	app.Use(AcceptCORS)

	// Add nice json headers
	app.Use(AddJSONHeader)

	// just to check api is responding
	app.Get("/healthz", handlers.Healthz) // a "response checker"
	app.Get("/version", handlers.Version)
	app.Post("/generateAdmin/:launcher_id", handlers.GenerateAdmin)
	app.Get("/name", handlers.Name)
	app.Get("/miningpoolstat", handlers.MiningPoolStat)
	app.Post("/login", handlers.Login)
	app.Group("/farmer", func(r martini.Router) {
		r.Get("/all", handlers.GetFarmers)
		r.Get("/top", handlers.GetTopFarmers)
		r.Get("/totalpoints", handlers.GetTotalPoint)
		r.Get("/count", handlers.GetFarmersCount)
		r.Get("/:launcher_id", handlers.GetFarmer)
		r.Post("/:launcher_id", handlers.PostFarmerDiscord)
	})
	app.Group("/partial", func(r martini.Router) {
		r.Get("/all", handlers.GetPartials)
		r.Get("/:launcher_id", handlers.GetPartial)
		r.Get("/netspace/total", handlers.GetNetSpaceTotal)
		r.Get("/netspace/:launcher_id", handlers.GetNetSpaceByLauncherId)
		r.Post("/:launcher_id", handlers.PostPartialDiscord)
		r.Post("/update/solo_plot", utils.Authorized, handlers.PostPartialSoloplot)

	})

	return app
}
func GetServer() *martini.ClassicMartini {
	return server()
}
