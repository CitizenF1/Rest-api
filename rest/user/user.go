package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/database/db"
	"strconv"

	"github.com/gorilla/mux"
)

// User represents a user instance
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	db := db.Initdb()

	sqlQuery := `INSERT INTO users(first_name, last_name) VALUES($1, $2) RETURNING id`

	var id int64

	err = db.QueryRow(sqlQuery, user.FirstName, user.LastName).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	res := Response{
		ID:      id,
		Message: "User created successfully",
	}
	defer db.Close()
	json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	user, err := getUser(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	if user.FirstName == "" {
		res := Response{
			ID:      int64(id),
			Message: "No rows returned!",
		}
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	user := User{}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	updatedRows := updateUser(int64(id), user)

	msg := fmt.Sprintf("User updated successfully %v", updatedRows)
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	deletRows := deleteUser(int64(id))

	msg := fmt.Sprint("User delete successfully ", deletRows)

	res := Response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func getUser(id int64) (User, error) {
	db := db.Initdb()
	defer db.Close()

	user := User{}

	sqlQuery := `SELECT * FROM users WHERE id=$1`

	row := db.QueryRow(sqlQuery, id)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	switch err {
	case sql.ErrNoRows:
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatal(err)
	}
	return user, err
}

func updateUser(id int64, user User) int64 {
	db := db.Initdb()
	defer db.Close()

	sqlStatement := `UPDATE users SET first_name=$2, last_name=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, user.FirstName, user.LastName)
	if err != nil {
		log.Fatal(err)
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rowsAff
}

func deleteUser(id int64) int64 {
	db := db.Initdb()
	defer db.Close()

	sqlStatement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAff
}
