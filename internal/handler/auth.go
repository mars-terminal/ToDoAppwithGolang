package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"myToDoApp/internal/service/user"
)

func (h *Handler) singUp(c *gin.Context) {
	logger := log.WithField("handler", "signUp")

	var body user.User

	if err := c.BindJSON(&body); err != nil {
		logger.WithError(err).Fatal("invalid authentication data")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	id, err := h.Service.Authorization.CreateUser(body)
	if err != nil {
		logger.WithError(err).Error("failed to create user")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
