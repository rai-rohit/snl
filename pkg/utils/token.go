package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string
	jwt.StandardClaims
}

var key = []byte("my_pass")

func GenToken(email string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 10)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknString, err := token.SignedString(key)
	if err != nil {
		return tknString, err
	}

	return tknString, nil
}

func Parse(tkn string) (Claims, error) {
	claims := Claims{}

	token, err := jwt.ParseWithClaims(tkn, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, err
	}

	return claims, nil
}
