package mw

import (
	"api-gateway/biz/model/api"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey []byte

func init() {
	SecretKey = []byte("123")
}

type MyClaims struct {
	jwt.StandardClaims
	UId int64
}

// 验证通过设置uid，不通过则不设置，通过req.GetInt64("uid")方法获取uid
func JWTCheck(ctx context.Context, req *app.RequestContext) {
	token := api.Token{}
	err := req.BindAndValidate(&token)
	if err != nil {
		req.JSON(consts.StatusOK, utils.H{
			"status_code": consts.StatusInternalServerError,
			"status_msg":  "服务内部解析Token错误",
		})
		req.Abort()
	} else if token.Token == nil {
		req.Next(ctx)
	} else {
		pt, err := ParseToken(token.GetToken())
		if err == nil && pt.Valid {
			claims := pt.Claims.(*MyClaims)
			req.Set("uid", claims.UId)
		}
	}

}

func SignedToken(uid int64) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 30*24*60*60,
			Issuer:    "api-gateway",
		},
		UId: uid,
	})
	signedString, err := claims.SignedString(SecretKey)
	return signedString, err
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	return token, err
}
