package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCTX)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	return
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
