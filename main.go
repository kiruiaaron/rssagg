package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kiruiaaron/rssagg/internal/database"


	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries

}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the  environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL== "" {
		log.Fatal("DB_URL is not found in the  environment")
	}

	conn, err :=sql.Open("postgres",dbURL)
	if err != nil{
		log.Fatal(("can't connect to the database"))
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}


	router := chi.NewRouter()


	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handleReadiness)
	v1Router.Get("/err",handleErr)
	v1Router.Post("/users",apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v",portString)
	err = srv.ListenAndServe()									

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)
}
