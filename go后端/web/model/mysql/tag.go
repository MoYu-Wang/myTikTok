package mysql

import (
	"WebVideoServer/dao"
	"context"
)

func DeleteUserLookALLTag(ctx context.Context, userID int64) error {
	db := GetDB(ctx)
	tag := dao.UserLookTag{}
	return db.Table("UserLookTag").Where("UserID=?", userID).Delete(&tag).Error
}
