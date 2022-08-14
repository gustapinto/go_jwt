package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gustapinto/go_jwt/utils"
)

type FooController struct{}

func (c *FooController) Foo(ctx *gin.Context) {
	_, err := utils.ValidateTokenFromHeaders(ctx.Request.Header)
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, "Foo")
}
