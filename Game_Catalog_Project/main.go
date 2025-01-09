package main

import (
	"fmt"
	"log"
	"net/http"
	"Game_Catalog_Project/configs"
	"Game_Catalog_Project/handlers"
	"Game_Catalog_Project/middlewares"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	PORT := 2650
	
	configs.ConnectDB()
	if configs.DB == nil {
		log.Fatal("Database connection failed")
	}
	defer func() {
		if err := configs.DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("catalog"))
	mux.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeStaticFile(w, r, "catalog", fileServer)
	})
	mux.Handlefunc("/api/games/", handlers.HandleGames)
	mux.Handlefunc("/api/games", handlers.HandleGames)

	loggedMux := middlewares.LogRequestHandler(mux)

	fmt.Printf("Server berjalan di http://localhost:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), loggedMux))
}