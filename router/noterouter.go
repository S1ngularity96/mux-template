package router

import (
	"github.com/gorilla/mux"
	"mux-template/controller"
)

func DescribeNoteSubRouter(router *mux.Router){
	router.HandleFunc("/",controller.GetNotes).Methods("GET")
	router.HandleFunc("/",controller.PostNote).Methods("POST")
	router.HandleFunc("/{id}",controller.GetNoteById).Methods("GET")
	router.HandleFunc("/{id}",controller.DeleteNoteById).Methods("DELETE")
}