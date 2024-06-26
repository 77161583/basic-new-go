package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error //两个写法
	return u, err
}

func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error //两个写法
	return u, err
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueconflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueconflictsErrNo {
			//偶像冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) Update(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	// 执行更新操作
	err := dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", u.Id).Updates(&u).Error
	if err != nil {
		// 打印错误信息
		fmt.Printf("MySQL Error: %v\n", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return fmt.Errorf("mysql error: %v", mysqlErr.Message)
		}
	}
	return err
}

// User 对标数据库，定义模型
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//唯一索引
	Email           string `gorm:"unique"`
	Password        string
	NickName        string
	Birthday        string
	PersonalProfile string `gorm:"type:text"`
	IsDel           int64  `gorm:"default:0;not null"`

	//创建时间 毫秒数
	Ctime int64
	//更新时间
	Utime int64
}
