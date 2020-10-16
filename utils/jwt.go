package units

import (
	"errors"
	"fmt"
	"gin-web/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-ini/ini"
	"time"
)

type MyClaims struct {
	Id       uint   `josn:"id"`
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

func GenerateToken(username, password, platform string) (string, error) {

	user := models.GetUserByName(username)
	// 加密password字符串
	if password != user.Password {
		return "", errors.New("Incorrect username or password")
	}
	c := MyClaims{
		user.ID,
		user.Name,
		platform,
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
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return claims, nil
		} else {
			return nil, errors.New("Token expired")
		}
	}
	return nil, errors.New("Invalid token")
}
