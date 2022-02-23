package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"myToDoApp/entities"
)

func ParseErrorToHTTPErrorCode(err error) int {
	switch {
	case errors.Is(err, entities.ErrForbidden):
		return http.StatusForbidden
	case errors.Is(err, entities.ErrBadRequest):
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func NewErrorResponse(logger *logrus.Entry, c *gin.Context, status int, message string) {
	switch status {
	case http.StatusInternalServerError:
		fallthrough
	default:
		logger.Error(message)
		status = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(status, map[string]interface{}{
		"status":  status,
		"message": message,
	})
}
