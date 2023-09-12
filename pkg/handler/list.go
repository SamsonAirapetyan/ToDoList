package handler

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
