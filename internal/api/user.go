package api

import (
	"base_frame/internal/dto"
	"base_frame/internal/services"
	"base_frame/pkg/response"
	"github.com/gin-gonic/gin"
)

type User interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUser(svc services.User) User {
	return &user{svc: svc}
}

type user struct {
	svc services.User
}

func (obj *user) Create(c *gin.Context) {
	req := &dto.CreateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Create(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

func (obj *user) Update(c *gin.Context) {
	req := &dto.UpdateUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Update(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	response.Json(c, response.WithMsg("success"))
}

func (obj *user) Delete(c *gin.Context) {
	req := &dto.DeleteUserReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}

	if err := obj.svc.Delete(c, req); err != nil {
		response.Json(c, response.WithErr(err))
		return
	}
	response.Json(c, response.WithMsg("success"))
}
