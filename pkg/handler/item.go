package handler

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createitem(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid Id param")
		return
	}
	var input todo.TodoItem
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userid, listid, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid Id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userid, listid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)

}

func (h *Handler) getItemById(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	itemid, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid item_Id param")
		return
	}

	items, err := h.services.TodoItem.GetItemId(userid, itemid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)

}

func (h *Handler) updateItem(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	itemid, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid item_Id param")
		return
	}
	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.UpdateItem(userid, itemid, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}
	itemid, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid item_Id param")
		return
	}
	err = h.services.TodoItem.DeleteIdItem(userid, itemid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
