package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment variable")
	}

	// router
	router := chi.NewRouter()

	// cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value is 31536000 seconds (1 year)
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // hooding the v1Router to the readiness handler
	v1Router.Get("/err", handlerErr)

	// v1Router to handle path value and nest the v1Router to the v1
	// such way allow us to make newer versioning later like v2,
	// without changing the readiness or other core components
	router.Mount("/v1", v1Router)

	// server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is running on port %v", portString)

	err := srv.ListenAndServe()

	// err handling, exit if err
	if err != nil {
		log.Fatal(err)
	}
}
