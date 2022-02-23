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
		err := fmt.Errorf("cannot parse token: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	c.Set(userCTX, userId)
}

func getUserId(c *gin.Context) (int, error) {
	logger := log.WithField("handler", "get user id")

	id, ok := c.Get(userCTX)
	if !ok {
		err := fmt.Errorf("user id: %w", entities.ErrNotFound)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	idInt, ok := id.(int)
	if !ok {
		err := fmt.Errorf("user id: %w, invalid type", entities.ErrNotFound)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	return idInt, nil
}
