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

	numOfSteps, _ := utils.CountActionsToMakePasswordStrong(requestData.InitPassword)

	db.LogRequest(requestData.InitPassword, "Recommendation")

	c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
}
