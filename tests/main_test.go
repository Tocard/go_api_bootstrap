package main_test

/*
Testing helpers.
*/

import (
	"api/data"
	"api/server"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
)



// Generate a testing server.
func getTestingServer() *httptest.Server {
	return httptest.NewServer(server.GetServer())
}
