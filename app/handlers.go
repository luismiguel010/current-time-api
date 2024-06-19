package app

import (
	"encoding/json"
	"fmt"
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
			loc, err := time.LoadLocation(tz)
			if err != nil {
				writer.WriteHeader(http.StatusNotFound)
				writer.Write([]byte(fmt.Sprintf("invalid time zone %s", tz)))
			} else {
				currentTime := time.Now().In(loc)
				response[tz] = currentTime
			}
		}
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
