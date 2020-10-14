package units

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-ini/ini"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Paltform string `json:"paltform"`
	jwt.StandardClaims
}

const ExpireAt = time.Hour * 2

var JwtSecret string

func init() {
	var err error
	cfg, err := ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}
	JwtSecret = cfg.Section("app").Key("JWT_SECRET").String()
}

func GenerateToken() (string, error) {
	c := MyClaims{
		"username",
		"paltform",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireAt).Unix(),
			Issuer:    "event-mysql-web",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(JwtSecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid Token")
}
