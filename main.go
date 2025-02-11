package main

import (
    "billing-microservice/config"
    "billing-microservice/database"
    "log"
)

func main() {
    // database connect
    database.Connect()

    //Initialize the server
    config.InitializeServer()

    // start the server
    if err := config.Router.Run(":8080"); err != nil {
        log.Fatal("Error starting the server: ", err)
    }
}