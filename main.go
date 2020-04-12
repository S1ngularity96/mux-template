package main

import (
	"mux-template/models"
	"mux-template/router"
)

func main(){
	models.OpenDatabase()
	models.MigrateData()
	defer models.DbCon.Close()
	router.StartRouter()
}
