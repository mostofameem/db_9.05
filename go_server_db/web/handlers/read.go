package handlers

import (
	"go_server_db/db"
	"log"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
	log.Printf("this is Read")
	db.ReadUsers(w)
	/*
		if err != nil {
			http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
			return
		}
		utils.SendJson(w, http.StatusAccepted, users)
	*/
}
