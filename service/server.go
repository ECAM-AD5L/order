package service

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartAPIServer() {
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.LoggingHandler(os.Stdout, handlers.CORS(originsOk, headersOk, methodsOk)(handlers.RecoveryHandler()(router))),
	}
	go listenServer(server)
	waitForShutdown(server)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Println(fmt.Sprintf("Error in starting server - %s", err))
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	fmt.Println("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	fmt.Println("API server shutdown complete")
}

