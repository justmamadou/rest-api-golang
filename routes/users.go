package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justmamadou/rest-api-golang/models"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	//user.ID = 1
	err = user.Signup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
