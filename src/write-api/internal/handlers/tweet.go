package handlers

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) CreateTweet(c *gin.Context) {
	var data *dto.CreateDto

	user, ok := middlewares.UserFromCtx(c)
	if !ok {
		common.SendResponse(c.Writer, http.StatusInternalServerError, "Could not get user from context", nil)
		return
	}

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tweet, err := h.service.CreateTweet(c.Request.Context(), user.Username, data.Content)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusCreated, "ok", tweet)
}

func (h *Handlers) UpdateTweet(c *gin.Context) {
	var data *dto.UpdateDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tweet, err := h.service.UpdateTweet(c.Request.Context(), data.Id, data.Content)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", tweet)
}

func (h *Handlers) DeleteTweet(c *gin.Context) {
	var data *dto.DeleteDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := h.service.DeleteTweet(c.Request.Context(), data.Id)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", id)
}
