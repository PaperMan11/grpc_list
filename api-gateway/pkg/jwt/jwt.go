package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("helloworld")
var expire = 24 * time.Hour // 令牌存在时长（h）

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(UserId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)
	myClaims := MyClaims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tan",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
