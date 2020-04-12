package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"mux-template/middleware"
	"net/http"
	"time"
)

const (
	PORT = "8000"
	Addr = "127.0.0.1"
)

func StartRouter(){
	r := mux.NewRouter()
	// Api for Notes
	noteSubRouter := r.PathPrefix("/api/note").Subrouter()
	// RootPrefixes
	rootSubRouter := r.PathPrefix("/").Subrouter()
	rootSubRouter.Use(middleware.LoggingMiddleware)

	DescribeNoteSubRouter(noteSubRouter)
	DescribeRootSubRouter(rootSubRouter)
	fmt.Println("Listening on ", Addr, " on PORT ", PORT)
	srv := &http.Server{
		Handler : r,
		Addr: Addr+":"+PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15* time.Second,
	}

	srv.ListenAndServe()
}