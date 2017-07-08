package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func appInfo(response http.ResponseWriter, request *http.Request) {
	log.Println(DbStore)
	info := DbStore.GetAppInfo()
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	json.NewEncoder(response).Encode(&info)
}
