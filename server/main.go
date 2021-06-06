package server

import (
	"bytes"
	"encoding/json"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
	"time"
)

func Healthz() string {
	return "ok"
}

func Version() string {
	return "0.0.0"
}

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

	// Allow CORS
	app.Use(AcceptCORS)

	// Add nice json headers
	app.Use(AddJSONHeader)

	// just to check api is responding
	app.Get("/healthz", Healthz) // a "response checker"
	app.Get("/version", Version)


	return app
}
func GetServer() *martini.ClassicMartini {
	return server()
}
