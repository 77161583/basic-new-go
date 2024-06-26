package repository

import (
	"basic-new-go/webook/internal/domain"
	"basic-new-go/webook/internal/repository/dao"
	"context"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
var ErrUserNotFound = dao.ErrUserNotFound

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:              u.Id,
		Email:           u.Email,
		NickName:        u.NickName,
		Birthday:        u.Birthday,
		PersonalProfile: u.PersonalProfile,
		IsDel:           u.IsDel,
		//Password:        u.Password,
	}, nil
}

func (r *UserRepository) Created(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})

	//如果有缓存，这里操作缓存
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	return r.dao.Update(ctx, dao.User{
		Id:              u.Id, // 添加ID
		NickName:        u.NickName,
		Birthday:        u.Birthday,
		PersonalProfile: u.PersonalProfile,
		IsDel:           u.IsDel,
	})

	//如果有缓存，这里操作缓存
}
