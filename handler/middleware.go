package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentefy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		}

		ctx.Set(userCtx, userId)
		ctx.Next()
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return 0, errors.New("user id is invalid type")
	}

	return idInt, nil
}
