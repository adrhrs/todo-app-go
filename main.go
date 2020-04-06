package main

import (
	"log"
	"net/http"
	"time"

	pg "github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *pg.DB
}

func (a *App) init() {

	t0 := time.Now()

	a.Router = a.initRouter()
	a.DB = a.initDB()

	log.Printf("running on %s, init on %s ", APP_PORT, time.Since(t0))
	log.Fatal(http.ListenAndServe(APP_PORT, a.Router))
}

func main() {

	app := App{}
	app.init()

}
