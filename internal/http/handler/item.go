package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"myToDoApp/entities"
	"myToDoApp/internal/utils"
)

func (h *Handler) createItem(c *gin.Context) {
	logger := log.WithField("handler", "create item")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	var input entities.TodoItem
	if err := c.BindJSON(&input); err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	id, err := h.Service.TodoItem.Create(userId, listId, input)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItem(c *gin.Context) {
	logger := log.WithField("handler", "get all items")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	items, err := h.Service.TodoItem.GetAll(userId, listId)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {
	logger := log.WithField("handler", "get item by id")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	item, err := h.Service.TodoItem.GetAll(userId, itemId)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	logger := log.WithField("handler", "get item by id")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	err = h.Service.TodoItem.Delete(userId, itemId)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}
