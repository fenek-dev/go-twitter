package handlers

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/middlewares"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/dto"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Handlers) FindTweetById(c *gin.Context) {
	id := c.Param("id")
	ctx, span := h.tracer.Start(c.Request.Context(), "tweet.handler.FindTweetById")
	defer span.End()

	span.SetAttributes(attribute.String("id", id))

	if id == "" {
		common.SendResponse(c.Writer, http.StatusBadRequest, "incorrect_id", nil)
		return
	}

	tweet, err := h.service.FindTweetById(ctx, id)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("result_content", tweet.Content), attribute.String("author", tweet.Username))

	common.SendResponse(c.Writer, http.StatusOK, "ok", tweet)
}

func (h *Handlers) CreateTweet(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "tweet.handler.CreateTweet")
	defer span.End()
	var data *dto.CreateDto

	user, ok := middlewares.UserFromCtx(c)
	if !ok || user == nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, ErrCanNotGetUser.Error(), nil)
		return
	}

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}
	span.SetAttributes(attribute.String("username", data.Username), attribute.String("context", data.Content))

	tweet, err := h.service.CreateTweet(ctx, user.Username, data.Content)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("id", tweet.ID))

	common.SendResponse(c.Writer, http.StatusCreated, "ok", tweet)
}

func (h *Handlers) UpdateTweet(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "tweet.handler.UpdateTweet")
	defer span.End()
	var data *dto.UpdateDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("id", data.Id), attribute.String("context", data.Content))

	tweet, err := h.service.UpdateTweet(ctx, data.Id, data.Content)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", tweet)
}

func (h *Handlers) DeleteTweet(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "tweet.handler.DeleteTweet")
	defer span.End()
	var data *dto.DeleteDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("id", data.Id))

	id, err := h.service.DeleteTweet(ctx, data.Id)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", id)
}
