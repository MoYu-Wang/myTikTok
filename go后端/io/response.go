package io

import "errors"

//用户未登录错误
var ErrorUserNotLogin = errors.New("user not login")

//用户昵称已存在
var ErrorUserNameIsExist = errors.New("username is exist")
