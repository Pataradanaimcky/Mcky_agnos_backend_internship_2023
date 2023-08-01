package api

import (
	"Mcky_agnos_backend_internship_2023/db"
	"Mcky_agnos_backend_internship_2023/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StrongPasswordStepsHandler(c *gin.Context) {
	var requestData struct {
		InitPassword string `json:"init_password"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Calculate the number of steps required to make the password strong
	numOfSteps, _ := utils.CountActionsToMakePasswordStrong(requestData.InitPassword)

	// Log the request in the database
	db.LogRequest(requestData.InitPassword, "Recommendation")

	// Return the number of steps as a response
	c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
}
