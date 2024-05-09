package main

import (
	"go_server_db/db"
	"go_server_db/web"
	"log"
	"net/http"
)

func main() {

	if err := db.InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}
	//defer db.Db.Close()

	mux := web.StartServer()
	log.Printf("Server Running on port 30000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
