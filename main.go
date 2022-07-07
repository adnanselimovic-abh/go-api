package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId string `json:"clientId,omitempty"`
	name   string `json:"ip,omitempty"`
}

var users = []User{}

func getUserById(c *gin.Context) {
	ID := c.Param("id")

	for _, i := range users {
		if i.UserId == ID {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found!"})
}

func addUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}

	users = append(users, newUser)

	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	router.GET("/user/:id", getUserById)
	router.POST("/addUser", addUser)

	router.Run(":8080")
}
