package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{
		ID:   1,
		Name: "John Doe",
	},
	{
		ID:   2,
		Name: "John",
	},
	{
		ID:   3,
		Name: "Doe",
	},
}

func findUserByID(id int) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid ID format", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID format"))
		return
	}
	user := findUserByID(id)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// func main() {
// 	http.HandleFunc("/user", GetUser)
// 	http.ListenAndServe(":8080", nil)
// }
