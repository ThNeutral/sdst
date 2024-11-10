package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", handlers.HandleCreateUser(queries))
		r.Post("/login-email", handlers.HandleLoginByEmail(queries))
		r.Post("/login-username", handlers.HandleLoginByUsername(queries))
	})

	router.Route("/utils", func(r chi.Router) {
		r.Get("/ping-gateway", handlers.Gateway(queries, handlers.HandlePingGateway))
	})

	log.Printf("Listening on port :%v\n", *port)
	http.ListenAndServe(":"+*port, router)
}
