package model

// // 注册用户
// func InsertUser(ctx context.Context, user *dao.User) error {
// 	//mysql 操作
// 	err := mysql.InsertUser(ctx, user)
// 	if err != nil {
// 		return err
// 	}
// 	//redis 操作

// 	return err
// }

// // 用户是否存在
// func QueryUserIsExist(ctx context.Context, userID int64) (bool, error) {
// 	ret, err := mysql.QueryUserIsExist(ctx, userID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if ret {
// 		return true, err
// 	}
// 	return false, err

// }

// // 修改密码
// func UpdateUserPassword(ctx context.Context, p *io.ParamUpdate) error {
// 	err := mysql.UpdateUserPassword(ctx, p)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
