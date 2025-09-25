package jwt

import (
	"dream-vue-admin/global"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GeneratorToken 生成token
func GeneratorToken(sub any) (string, error) {
	iat := jwt.NewNumericDate(time.Now())
	exp := iat.Add(time.Minute * time.Duration(global.Cfg.Jwt.Expire))
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprint(sub),
		"iat": iat,
		"exp": exp,
		"iss": "dream-vue-admin",
	}).SignedString([]byte(global.Cfg.Jwt.SecretKey))
}

// ParseToken 解析token
func ParseToken(token string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Cfg.Jwt.SecretKey), nil
	})
	return claims, err
}
