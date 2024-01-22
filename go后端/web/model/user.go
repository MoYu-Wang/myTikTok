package model

import (
	"WebVideoServer/dao"
	"WebVideoServer/web/mysql"
	"context"
)

func InsertUser(ctx context.Context, user *dao.User) error {
	//mysql 操作
	err := mysql.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	//redis 操作

	return err
}
