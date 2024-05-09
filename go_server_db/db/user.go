package db

import (
	"encoding/json"
	"fmt"
	"go_server_db/web/utils"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InsertIntoDB(w http.ResponseWriter, r *http.Request) error {

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		message := "Error converting user to JSON"
		utils.SendError(w, http.StatusNotAcceptable, message, newUser)
		return err
	}
	// Insert the new user into the database
	_, errr := Db.Exec(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`, newUser.ID, newUser.Name, newUser.Email, newUser.Password)

	if errr != nil {
		log.Fatal("Error inserting user:", err)
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return errr
	}

	fmt.Println("User inserted successfully!")
	w.WriteHeader(http.StatusCreated)
	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Path[len("/users/"):]

	_, er := Db.Exec(`delete from users where id=$1;`, id)

	if er != nil {
		fmt.Fprintf(w, "Can not delete user")
		return er
	}
	fmt.Fprintf(w, "Deleted Sucessfully")
	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Path[len("/users/"):]
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		message := "Error converting user to JSON"
		utils.SendError(w, http.StatusNotAcceptable, message, user)
		return err
	}
	_, er := Db.Exec(`UPDATE users SET name = $1 WHERE id = $2`, user.Name, id)

	if er != nil {
		message := "Can not update the value"
		utils.SendError(w, http.StatusNotAcceptable, message, user)
	}
	fmt.Fprintf(w, "Update Sucessfully")
	return nil
}

func ReadUsers(w http.ResponseWriter) {

	var users []User

	err := Db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatal("Err not equal nil")
		return
	}
	utils.SendJson(w, http.StatusAccepted, users)

}
func ReadUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/users/"):]

	var user []User

	err := Db.Select(&user, `SELECT * from  users  WHERE id = $1`, id)

	if err != nil {
		message := "Error loading user"
		utils.SendError(w, http.StatusNotAcceptable, message, user)
		return
	}
	utils.SendJson(w, http.StatusAccepted, user)
}
