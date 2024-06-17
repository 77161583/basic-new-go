package repository

import (
	"basic-new-go/webook/internal/domain"
	"basic-new-go/webook/internal/repository/dao"
	"context"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Created(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})

	//如果有缓存，这里操作缓存
}

func (r *UserRepository) FindById(int642 int64) {
	//先从cache找
	//再从dao找
	//找到会写cache

}
