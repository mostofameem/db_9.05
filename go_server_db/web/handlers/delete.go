package handlers

import (
	"go_server_db/db"
	"log"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteUser(w, r)
	if err != nil {
		log.Fatal("Insertion Failed")
	}
}
