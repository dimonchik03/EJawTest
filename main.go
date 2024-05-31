package main

import (
	"EJawTest/db"
	"EJawTest/handlers"
	"EJawTest/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()

	router.HandleFunc("/auth", handlers.AuthHandler).Methods("POST")

	apiRouter := router.PathPrefix("/orders").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware)

	apiRouter.HandleFunc("", handlers.GetAllOrders).Methods("GET")
	apiRouter.HandleFunc("/{id:[0-9]+}", handlers.GetOrder).Methods("GET")
	apiRouter.HandleFunc("", handlers.CreateOrder).Methods("POST")
	apiRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateOrder).Methods("PUT")
	apiRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteOrder).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
