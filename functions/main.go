package main

import (
	"log"
	"os"
	"pheuanpet-user-management/router"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
    }

    r := router.SetupRouter()
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Could not listen to port %s: %v", port, err)
    }
}