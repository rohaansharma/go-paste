package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pastebin-backend/db"
	"pastebin-backend/handlers"
	"pastebin-backend/middleware"
	"pastebin-backend/utils"
)

func main() {
	err := utils.LoadEnvFile(".env")
	if err != nil {
		log.Fatal("No .env file found or error reading .env file:", err)
	}

	db.InitDB()
	defer db.CloseDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/paste", handlers.CreatePasteHandler)
	mux.HandleFunc("/api/paste/", handlers.GetPasteHandler)

	fs := http.FileServer(http.Dir("./build"))
	mux.Handle("/", fs)

	handler := middleware.EnableCORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
