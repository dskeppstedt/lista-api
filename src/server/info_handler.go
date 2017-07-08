package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func appInfo(response http.ResponseWriter, request *http.Request) {
	log.Println(DbStore)
	info := DbStore.GetAppInfo()
	json.NewEncoder(response).Encode(&info)
}
