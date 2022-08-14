package controller

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gustapinto/go_jwt/model"
)

type FooController struct{}

func (c *FooController) Foo(ctx *gin.Context) {
	authorization, exists := ctx.Request.Header["Authorization"]
	if !exists {
		ctx.IndentedJSON(http.StatusBadRequest, "Missing authorization header")
		return
	}

	token := strings.Split(authorization[0], " ")[1]
	claims := &model.User{}

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Status(http.StatusInternalServerError)
		return
	}

	if !jwtToken.Valid {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{"Foo"}

	ctx.IndentedJSON(http.StatusOK, response)
}
