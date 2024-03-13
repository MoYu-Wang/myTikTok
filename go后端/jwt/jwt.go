package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySecret = []byte("先这样写")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"userID"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

var outtime = 3 * 24 * time.Hour

// GenToken 生成JWT
func GenToken(userID int64, userName string) (string, error) {
	// 创建一个我们自己的声明的数据
	g := MyClaims{
		userID,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(outtime).Unix(), // 过期时间
			Issuer:    "tiktok",                       // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, g)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken解析jwt
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析Token
	var mg = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mg, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if token.Valid {
		fmt.Println(err)
		return mg, nil
	}
	return nil, errors.New("invaild token")
}
