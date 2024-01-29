package main

import (
    "log"
    "net/http"
    "go_application/controllers"
    "go_application/util"
    "github.com/gorilla/mux"
)

func main() {

    config, err := util.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    util.InitDB(config.DatabaseURL)

    r := mux.NewRouter()
    r.HandleFunc("/items", controllers.GetItems).Methods("GET")

    log.Println("Starting server on :8081")
    if err := http.ListenAndServe(":8081", r); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}
