package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"myToDoApp/entities"
	"myToDoApp/internal/utils"
)

const (
	authorizationHeader = "Authorization"
	userCTX             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	logger := log.WithField("middleware", "userIdentify")

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		err := fmt.Errorf("header is empty")
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := fmt.Errorf("header is invalid: %w", entities.ErrBadRequest)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	userId, err := h.Service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	c.Set(userCTX, userId)
}
