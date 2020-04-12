package router

import (
	"github.com/gorilla/mux"
	"mux-template/controller"
)

func DescribeRootSubRouter(router *mux.Router){
	router.HandleFunc("/", controller.GetRootPage).Methods("GET")
}