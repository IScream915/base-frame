package services

import (
	"base_frame/internal/dto"
	"base_frame/internal/repo"
	"base_frame/internal/repo/models"
	"base_frame/pkg/pcontext"
	"context"
)

type User interface {
	Create(ctx context.Context, req *dto.CreateUserReq) error
	Update(ctx context.Context, req *dto.UpdateUserReq) error
	Delete(ctx context.Context, req *dto.DeleteUserReq) error
}

func NewUser(repo repo.User, tokenRepo repo.UserToken) User {
	return &user{
		tokenRepo: tokenRepo,
		repo:      repo,
	}
}

type user struct {
	tokenRepo repo.UserToken
	repo      repo.User
}

func (obj *user) Create(ctx context.Context, req *dto.CreateUserReq) error {
	user := models.User{
		Account:  req.Account,
		NickName: req.NickName,
		Password: req.Password,
		Age:      req.Age,
		Sex:      req.Sex,
	}
	return obj.repo.Create(ctx, &user)
}

func (obj *user) Update(ctx context.Context, req *dto.UpdateUserReq) error {
	userToken, err := pcontext.GetUserTokenFromCtx(ctx)
	if err != nil {
		return err
	}
	newUser := &models.User{}
	newUser.ID = userToken.UserID
	newUser.NickName = req.NickName
	newUser.Password = req.Password
	newUser.Age = req.Age
	newUser.Sex = req.Sex

	return obj.repo.Update(ctx, newUser)
}

func (obj *user) Delete(ctx context.Context, req *dto.DeleteUserReq) error {
	return obj.repo.DeleteByIds(ctx, req.IDs)
}
