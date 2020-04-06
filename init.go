package main

import (
	"log"
	"net/http"
	"time"

	pg "github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

func (a *App) initRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/api/activity/list", a.GetActivityList).Methods("GET")
	r.HandleFunc("/api/activity/create", a.InsertActivity).Methods("POST")
	r.HandleFunc("/api/activity/update", a.UpdateActivity).Methods("PUT")
	r.HandleFunc("/api/activity/delete", a.DeleteActivity).Methods("DELETE")

	r.Use(loggingMiddleware)

	return r

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Println(r.Method, t2.Sub(t1))
	})
}

func (a *App) initDB() *pg.DB {

	db := pg.Connect(&pg.Options{
		User:     "localuser",
		Password: "localpass",
		Database: "localdb",
	})

	return db
}
