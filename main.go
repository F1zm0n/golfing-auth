package main

import (
	"database/sql"
	"github.com/F1zm0n/auth.git/handlers"
	Db "github.com/F1zm0n/auth.git/myDb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connString := "postgres://postgres:F1zm0_007@localhost:5432/goselfmake?sslmode=disable"
	DataB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("couldnt connect to database:", err)
	}
	Db.CreateUserTable(DataB)

	defer DataB.Close()
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.HandlerHealth)
	router.Mount("/bim", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + "8080",
	}
	log.Printf("Server is starting on port %v", "8080")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
