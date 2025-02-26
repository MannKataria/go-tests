// package main
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{1, 30, "John Doe"},
	{2, 30, "John"},
	{3, 30, "Doe"},
}

const url = "/ping/:id"
const userNotFoundErr = "User not found"

func findUserByID(id int) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func findIndexByID(id int) int {
	for ind, user := range users {
		if user.Id == id {
			return ind
		}
	}
	return -1
}

func validateId(c *gin.Context) int {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return -1
	}
	return id
}

func validateJSON(json *data, c *gin.Context) (data, error) {
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return *json, err
	}
	return *json, nil
}

func handleGet(r *gin.Engine) {
	r.GET(url, func(c *gin.Context) {
		id := validateId(c)
		if id == -1 {
			return
		}
		user := findUserByID(id)
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": userNotFoundErr})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"id":   user.Id,
				"name": user.Name,
				"Age":  user.Age,
			})
		}
	})
}

func handlePost(r *gin.Engine) {
	r.POST("/ping", func(c *gin.Context) {
		var json data
		json, err := validateJSON(&json, c)
		if err != nil {
			return
		}
		id := len(users) + 1
		user := User{
			Id: id, Age: json.Age, Name: json.Name,
		}
		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
			"id":      id,
			"name":    json.Name,
			"age":     json.Age,
		})
	})
}

func handlePut(r *gin.Engine) {
	r.PUT(url, func(c *gin.Context) {
		id := validateId(c)
		if id == -1 {
			return
		}
		index := findIndexByID(id)
		if index == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": userNotFoundErr})
			return
		} else {
			var json data
			json, err := validateJSON(&json, c)
			if err != nil {
				return
			}
			users[index] = User{
				Id: id, Name: json.Name, Age: json.Age,
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated",
				"id":      id,
				"name":    json.Name,
				"age":     json.Age,
			})
		}
	})
}

func handlePatch(r *gin.Engine) {
	r.PATCH(url, func(c *gin.Context) {
		id := validateId(c)
		if id == -1 {
			return
		}
		index := findIndexByID(id)
		if index == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": userNotFoundErr})
			return
		} else {
			var json data
			json, err := validateJSON(&json, c)
			if err != nil {
				return
			}
			if json.Name != "" {
				users[index].Name = json.Name
			} else if json.Age != -1 {
				users[index].Age = json.Age
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "User partially updated",
				"id":      id,
				"name":    users[index].Name,
				"age":     users[index].Age,
			})
		}
	})
}

func handleDelete(r *gin.Engine) {
	r.DELETE(url, func(c *gin.Context) {
		id := validateId(c)
		if id == -1 {
			return
		}
		index := findIndexByID(id)
		if index == -1 {
			c.JSON(http.StatusNotFound, gin.H{"error": userNotFoundErr})
			return
		} else {
			name := users[index].Name
			age := users[index].Age
			fmt.Println(users)
			users = append(users[:index], users[index+1:]...)
			fmt.Println(users)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted",
				"id":      id,
				"name":    name,
				"age":     age,
			})
		}
	})
}

func CreateRouter() *gin.Engine {
	r := gin.Default()

	handleGet(r)
	handlePost(r)
	handlePut(r)
	handlePatch(r)
	handleDelete(r)

	return r
}

// func main() {
// 	r := CreateRouter()
// 	r.Run(":8080") // Run the server
// }
