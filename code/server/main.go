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
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/thneutral/sdst/code/server/internal/database"
	"github.com/thneutral/sdst/code/server/internal/editorhub"
	"github.com/thneutral/sdst/code/server/internal/handlers"
)

var editorHub *editorhub.Hub
var queries *database.Queries
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	godotenv.Load()
	db_url := os.Getenv("CONN_STRING")
	if db_url == "" {
		fmt.Println("CONN_STRING is not found in .env")
		os.Exit(1)
	}

	port := flag.String("port", "5432", "Set server port. By default port is 5432")
	flag.Parse()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, db_url)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(ctx)

	queries = database.New(conn)

	editorHub = editorhub.GetNewEditorHub()
	go editorHub.Run()

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
		r.Delete("/delete", handlers.Gateway(queries, handlers.HandleDeleteUser(queries)))
		r.Post("/login-email", handlers.HandleLoginByEmail(queries))
		r.Post("/login-username", handlers.HandleLoginByUsername(queries))
	})

	router.Route("/utils", func(r chi.Router) {
		r.Get("/ping-gateway", handlers.Gateway(queries, handlers.HandlePingGateway))
	})

	router.Route("/editor", func(r chi.Router) {
		r.Get("/open", handlers.WSGateway(upgrader, queries, handlers.HandleEditorHub(editorHub)))
	})

	router.Route("/messenger", func(r chi.Router) {
		r.Post("/create", handlers.HandleCreateMessage(queries))
		r.Get("/fetch{projectId}", handlers.HandleGetMessages(queries))
	})

	log.Printf("Listening on port :%v\n", *port)
	http.ListenAndServe(":"+*port, router)
}
