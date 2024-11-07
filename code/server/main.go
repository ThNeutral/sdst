package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thneutral/sdst/code/server/internals/dummydb"
	"github.com/thneutral/sdst/code/server/internals/handlers"
)

var db *dummydb.DummyDB

func main() {
	port := flag.String("port", "8080", "Set server port. By default port is 8080")
	flag.Parse()

	db = dummydb.GetDummyDB()
	go db.Run()

	router := chi.NewRouter()

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", handlers.HandleCreateUser(db))
		r.Post("/login-email", handlers.HandleLoginByEmail(db))
		r.Post("/login-username", handlers.HandleLoginByUsername(db))
	})

	router.Route("/utils", func(r chi.Router) {
		r.Get("/ping-gateway", handlers.Gateway(db, handlers.HandlePingGateway))
	})

	log.Printf("Listening on port :%v\n", *port)
	http.ListenAndServe(":"+*port, router)
}
