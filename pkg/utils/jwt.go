package utils

import (
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/xerr"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	JwtModule jwt.SigningMethod = jwt.SigningMethodHS256
)

type CustomClaims struct {
	UserId   int      `json:"user_id"`  //前台app的用户id
	Roles    []string `json:"roles"`    //角色id集合
	Nickname string   `json:"nickname"` //前台app的昵称
	Avatar   string   `json:"avatar"`   //头像
	jwt.StandardClaims
}

func CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(JwtModule, claims)
	return token.SignedString([]byte(driver.Conf.Jwt.Secret))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(driver.Conf.Jwt.Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, xerr.NewErrCode(xerr.TokenMalformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, xerr.NewErrCode(xerr.TokenExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, xerr.NewErrCode(xerr.TokenNotValidYet)
			} else {
				return nil, xerr.NewErrCode(xerr.TokenInvalid)
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, xerr.NewErrCode(xerr.TokenInvalid)
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(driver.Conf.Jwt.Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", xerr.NewErrCode(xerr.TokenMalformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "", xerr.NewErrCode(xerr.TokenExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", xerr.NewErrCode(xerr.TokenNotValidYet)
			} else {
				return "", xerr.NewErrCode(xerr.TokenInvalid)
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		t, _ := strconv.Atoi(driver.Conf.Jwt.Expire)
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(t) * time.Second).Unix()
		token, err := CreateToken(*claims)
		if err != nil {
			return "", errors.New("刷新生成token失败")
		}
		return token, nil
	}

	return "", xerr.NewErrCode(xerr.TokenInvalid)
}
