package handler

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create item
// @Security ApiKeyAuth
// @Tags items
// @Description Creating a new item
// @ID Create-Item
// @Accept  json
// @Produce  json
// @Param input body todo.TodoItem true "Item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/items [post]
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

// @Summary Get all items
// @Security ApiKeyAuth
// @Tags items
// @Description Get All items
// @ID Show-Items
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id/items [get]
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

// @Summary Get item
// @Security ApiKeyAuth
// @Tags items
// @Description Get item by ID
// @ID Item-by-ID
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:item_id [get]
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

// @Summary Update item
// @Security ApiKeyAuth
// @Tags items
// @Description Update item by ID
// @ID Item-update
// @Param input body todo.UpdateItemInput true "item update info"
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:item_id [post]
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

// @Summary Delete item
// @Security ApiKeyAuth
// @Tags items
// @Description Delete item by ID
// @ID Delete-ID
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/:item_id [delete]
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
