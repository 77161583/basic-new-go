package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now

	return dao.db.WithContext(ctx).Create(&u).Error
}

// User 对标数据库，定义模型
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//唯一索引
	Email    string `gorm:"unique"`
	Password string

	//创建时间 毫秒数
	Ctime int64
	//更新时间
	Utime int64
}
