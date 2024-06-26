package service

import (
	"basic-new-go/webook/internal/domain"
	"basic-new-go/webook/internal/repository"
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码错误")
var ErrUserIdNotFund = errors.New("用户信息不存在")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//加密问题
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Created(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	//查看账号是否存在
	u, err := svc.repo.FindByEmail(ctx, email)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//对比密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//打日志
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, err

}

func (svc *UserService) Edit(ctx context.Context, u domain.User) error {
	// 打印编辑的用户信息
	fmt.Printf("Editing User: %+v\n", u)

	// 查看id是否存在是否被删除了
	_, err := svc.repo.FindById(ctx, u.Id)
	if errors.Is(err, repository.ErrUserNotFound) {
		return ErrUserIdNotFund
	}
	return svc.repo.Update(ctx, u)
}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	// 查看id是否存在是否被删除了
	userData, err := svc.repo.FindById(ctx, id)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrUserIdNotFund
	}
	return userData, nil
}
