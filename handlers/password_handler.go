package handlers

import (
	"net/http"

	"strong_password_backend/services"

	"github.com/gin-gonic/gin"
)

type PasswordHandler struct {
	passwordService *services.PasswordService
}

func NewPasswordHandler(passwordService *services.PasswordService) *PasswordHandler {
	return &PasswordHandler{passwordService: passwordService}
}

func (h *PasswordHandler) GetStrongPasswordSteps(c *gin.Context) {
	var req struct {
		InitPassword string `json:"init_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	steps := h.passwordService.GetMinimumStepsToStrongPassword(req.InitPassword)
	c.JSON(http.StatusOK, gin.H{"num_of_steps": steps})
}
