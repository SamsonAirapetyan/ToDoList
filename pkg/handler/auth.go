package handler

import (
	"github.com/SamsonAirapetyan/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Problem with request")
		return
	}
	var er error
	//er = h.services.Authorization.CreateUser(input)
	id, er := h.services.Authorization.CreateUser(input)
	if er != nil {
		NewErrorResponse(c, http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
