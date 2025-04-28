package repo

import (
	"base_frame/internal/repo/models"
	"context"
	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	DeleteByIds(ctx context.Context, ids []uint64) error
	FindById(ctx context.Context, id uint64) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	WithTx(tx *gorm.DB) User
}

func NewUser(client *gorm.DB) User {
	return &user{
		client: client,
	}
}

type user struct {
	client *gorm.DB
}

func (obj *user) Create(ctx context.Context, user *models.User) error {
	return obj.client.WithContext(ctx).Create(user).Error
}

func (obj *user) Update(ctx context.Context, user *models.User) error {
	// 使用Updates更新非零字段的写法
	// 这里当user包函主键的时候可以省略Where不写
	return obj.client.WithContext(ctx).Model(models.User{}).Updates(&user).Error
	// 使用Save的全量更新的简洁写法, 使用Save的时候Gorm会自动根据传入的model识别对应的数据库表名
	// return obj.client.WithContext(ctx).Save(&user).Error
}

func (obj *user) DeleteByIds(ctx context.Context, ids []uint64) error {
	//users := make([]*models.User, 0)
	return obj.client.WithContext(ctx).Delete(&models.User{}, ids).Error
}

func (obj *user) FindById(ctx context.Context, id uint64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (obj *user) FindAll(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := obj.client.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (obj *user) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	//TODO implement me
	panic("implement me")
}

// WithTx 事务数据库
func (obj *user) WithTx(tx *gorm.DB) User {
	return &user{client: tx}
}
