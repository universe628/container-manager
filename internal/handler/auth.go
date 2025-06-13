package handler

import (
	"container-manager/internal/handler/dto"
	"container-manager/internal/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	body := new(dto.UserLoginReq)

	if err := c.ShouldBindJSON(body); err != nil {
		c.Error(err)
		return
	}

	token, err := h.authService.Login(c.Request.Context(), &schema.User{
		UserName: body.UserName,
		Password: body.Password,
	})
	if err != nil {
		c.Error(err)
		return
	}

	dto.NewUserLoginRes(token)

	c.JSON(http.StatusOK, dto.NewUserLoginRes(token))

}

func (h *AuthHandler) Register(c *gin.Context) {
	body := new(dto.UserRegistReq)

	if err := c.ShouldBindJSON(body); err != nil {
		c.Error(err)
		return
	}

	err := h.authService.NewUser(c.Request.Context(), &schema.User{
		UserName: body.UserName,
		Password: body.Password,
	})

	if err != nil {
		c.Error(err)
		return
	}
}
