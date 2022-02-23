package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"myToDoApp/internal/service/user"
)

func (h *Handler) signUp(c *gin.Context) {
	logger := log.WithField("handler", "signUp")

	var body user.User

	if err := c.BindJSON(&body); err != nil {
		logger.WithError(err).Fatal("invalid authentication data")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}

	id, err := h.Service.Authorization.CreateUser(body)
	if err != nil {
		logger.WithError(err).Error("failed to create user")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	NickName string `json:"nickName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	logger := log.WithField("handler", "signIn")

	var signInInput signInInput

	if err := c.BindJSON(&signInInput); err != nil {
		logger.WithError(err).Fatal("incorrect authentication data")
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		})
		return
	}

	token, err := h.Service.Authorization.GenerateToken(signInInput.NickName, signInInput.Password)
	if err != nil {
		logger.WithError(err).Error("failed to generate token")
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
