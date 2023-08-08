package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("PORT is not set")
	}
	router:=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1router:=chi.NewRouter();
	v1router.Get("/ready",handlerReadiness)
	v1router.Get("/*",handleerr)

	router.Mount("/v1",v1router)




	server:=&http.Server{
		Addr:":"+portstring,
		Handler:router,
	}

	log.Printf("Listening on port %s",portstring)
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}

}
