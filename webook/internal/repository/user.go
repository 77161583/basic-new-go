package repository

import (
	"basic-new-go/webook/internal/domain"
	"context"
)

type UserRepository struct {
}

func (r *UserRepository) Created(ctx context.Context, u domain.User) error {

}

func (r *UserRepository) FindById(int642 int64) {
	//先从cache找
	//再从dao找
	//找到会写cache

}
