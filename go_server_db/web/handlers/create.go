package handlers

import (
	"go_server_db/db"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	err := db.InsertIntoDB(w, r)
	if err != nil {
		log.Fatal("Insertion Failed")
	}
}
