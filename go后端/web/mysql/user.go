package mysql

import (
	"WebVideoServer/dao"
	"context"
)

func InsertUser(ctx context.Context, user *dao.User) error {
	db := GetDB(ctx)

	return db.Table("users").Create(user).Error
}
