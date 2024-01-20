package model

import (
	"WebVideoServer/dao"
	"WebVideoServer/web/mysql"
	"context"
)

func InsertUser(ctx context.Context, user *dao.User) error {
	err := mysql.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return err
}
