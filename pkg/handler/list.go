package handler

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create list
// @Security ApiKeyAuth
// @Tags lists
// @Description Creating a new list
// @ID CreateList
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "List info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userid, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAllListRequest struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags lists
// @Description Get all lists
// @ID ShowList
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAllLists(userid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllListRequest{
		Data: lists,
	})
}

// @Summary Get list
// @Security ApiKeyAuth
// @Tags lists
// @Description Get list by ID
// @ID List-ID
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}

	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid Id param")
		return
	}

	lists, err := h.services.TodoList.GetIdList(userid, listid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lists)
}

// @Summary Update list
// @Security ApiKeyAuth
// @Tags lists
// @Description Update list by ID
// @ID List-Update
// @Param input body todo.UpdateListInput true "list update info"
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id [post]
func (h *Handler) updateList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid Id param")
		return
	}

	err = h.services.TodoList.Update(userid, listid, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})

}

// @Summary Delete list
// @Security ApiKeyAuth
// @Tags lists
// @Description Delete list by ID
// @ID List-delete
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		return
	}
	listid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid Id param")
		return
	}

	err = h.services.TodoList.DeleteIdList(userid, listid)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
