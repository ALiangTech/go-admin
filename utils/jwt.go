package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//

var signingKey = []byte("#@-s889vske2!)}")

type CustomClaim struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

const fixedString = "Bearer "

// Bearer
func getJwtFromAuthorization(Authorization string) string {
	start := len(fixedString)
	end := len(Authorization)
	return Authorization[start:end]
}

// 生成jwt

func GenerateJwt(uuid string) (string, error) {
	clamis := CustomClaim{
		uuid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "hj",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	jwtToken, err := token.SignedString(signingKey)
	return fmt.Sprintf("%s%s", fixedString, jwtToken), err
}

// 解析/验证jwt

func ParseJwt(jwtToken string) (*CustomClaim, error) {
	payload := CustomClaim{}
	var ok bool
	token, err := jwt.ParseWithClaims(getJwtFromAuthorization(jwtToken), &CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return &payload, errors.New("jwt 认证失败")
	}
	claims, ok := token.Claims.(*CustomClaim)
	if ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("jwt 认证失败")
	}
}
