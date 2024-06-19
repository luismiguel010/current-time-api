package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(writer http.ResponseWriter, request *http.Request) {
	param := request.URL.Query().Get("tz")
	response := make(map[string]time.Time)
	if param == "" {
		currentTime := time.Now().UTC()
		response["current_time"] = currentTime
	} else {
		params := strings.Split(param, ",")
		for _, tz := range params {
			loc, _ := time.LoadLocation(tz)
			currentTime := time.Now().In(loc)
			response[tz] = currentTime
		}
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
