package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/thneutral/sdst/code/server/internal/database"
	"github.com/thneutral/sdst/code/server/internals/dummydb"
	"github.com/thneutral/sdst/code/server/internals/handlers"
)

var db *dummydb.DummyDB
var queries *database.Queries

func main() {
	godotenv.Load()
	db_url := os.Getenv("CONN_STRING")
	if db_url == "" {
		fmt.Println("CONN_STRING is not found in .env")
		os.Exit(1)
	}

	port := flag.String("port", "8080", "Set server port. By default port is 8080")
	flag.Parse()

	db = dummydb.GetDummyDB()
	go db.Run()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, db_url)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(ctx)

	queries = database.New(conn)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", handlers.HandleCreateUser(db, queries))
		r.Post("/login-email", handlers.HandleLoginByEmail(db))
		r.Post("/login-username", handlers.HandleLoginByUsername(db))
	})

	router.Route("/utils", func(r chi.Router) {
		r.Get("/ping-gateway", handlers.Gateway(db, handlers.HandlePingGateway))
	})

	log.Printf("Listening on port :%v\n", *port)
	http.ListenAndServe(":"+*port, router)
}
