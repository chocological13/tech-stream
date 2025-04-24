package handlers

import (
	"github.com/chocological13/tech-stream/api-gateway/internal/client"
	"github.com/chocological13/tech-stream/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

type AuthHandler struct {
	userClient *client.UserClient
}

func NewAuthHandler(userClient *client.UserClient) *AuthHandler {
	return &AuthHandler{
		userClient: userClient,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to protobuf request
	protoReq := &user.CreateUserRequest{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Bio:       req.Bio,
	}

	// Call user service
	createdUser, err := h.userClient.CreateUser(c, protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user err: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":        createdUser.Id,
			"username":  createdUser.Username,
			"email":     createdUser.Email,
			"firstName": createdUser.FirstName,
			"lastName":  createdUser.LastName,
			"bio":       createdUser.Bio,
		},
	})
}
