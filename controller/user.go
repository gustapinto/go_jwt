package controller

import (
	"crypto/sha256"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gustapinto/go_jwt/database"
	"github.com/gustapinto/go_jwt/model"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{database.DB()}
}

func (c *UserController) Create(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	hashedPassword := sha256.New().Sum([]byte(user.Password))
	user.Password = string(hashedPassword)

	if r := c.db.Create(&user); r.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, r.Error)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, user)
}

func (c *UserController) Auth(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	hashedPassword := sha256.New().Sum([]byte(user.Password))

	r := c.db.First(&model.User{}, "username = ? AND password = ?", user.Username, string(hashedPassword))
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusUnauthorized, r.Error)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := struct {
		Token string `json:"token"`
	}{ss}

	ctx.IndentedJSON(http.StatusOK, response)
}
