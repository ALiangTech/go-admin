package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//

var signingKey = []byte("#@-s889vske2!)}")

type CustomClaim struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

// 生成jwt

func GenerateJwt() (string, error) {
	clamis := CustomClaim{
		"1234",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
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
	return jwtToken, err
}

// 解析/验证jwt

func ParseJwt(jwtToken string) {
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		fmt.Printf("%v %vss", claims.Uid, claims.RegisteredClaims.IssuedAt)
	} else {
		fmt.Println(err, "err")
	}
}
