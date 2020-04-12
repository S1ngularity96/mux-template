package controller

import (
	"mux-template/auxiliary"
	"net/http"
)

func GetRootPage(w http.ResponseWriter, r *http.Request){
	auxiliary.OkResponse(w,r,"Willkommen auf der Hauptseite")
}