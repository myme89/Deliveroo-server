package main

import (
	"deliveroo/db"
	"deliveroo/routers"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello Trong Nhat")

	db.Init()
	serverAddr := "0.0.0.0:3000"
	router := routers.InitRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)
	srv := &http.Server{
		Handler: handler,
		Addr:    serverAddr,
		//enforce timeouts for servers
		WriteTimeout: 400 * time.Second,
		ReadTimeout:  400 * time.Second,
	}
	log.Info("Start server at : ", serverAddr)
	log.Fatal(srv.ListenAndServe())
}
