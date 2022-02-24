package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"myToDoApp/entities"
	"myToDoApp/internal/utils"
)

func (h *Handler) createList(c *gin.Context) {
	logger := log.WithField("handler", "create list")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var todolist entities.TodoList
	if err := c.BindJSON(&todolist); err != nil {
		err := fmt.Errorf("error: %w", entities.ErrBadRequest)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	id, err := h.Service.TodoList.Create(userId, todolist)
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

type getAllListsResponse struct {
	Data []entities.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	logger := log.WithField("handler", "get all lists")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.Service.TodoList.GetAll(userId)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {
	logger := log.WithField("handler", "get list by id")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	list, err := h.Service.TodoList.GetById(userId, id)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(c *gin.Context) {
	logger := log.WithField("handler", "delete list")

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}

	err = h.Service.TodoList.Delete(userId, id)
	if err != nil {
		err := fmt.Errorf("error: %w", err)
		logger.WithError(err)

		utils.NewErrorResponse(logger, c, utils.ParseErrorToHTTPErrorCode(err), err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
