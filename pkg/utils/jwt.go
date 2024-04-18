package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	ID uint32
	jwt.RegisteredClaims
}

var _key = []byte("123456")

func GenToken(id uint32) (string, error) {
	gen := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "palace",
			Subject:   "auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	token, err := gen.SignedString(_key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenStr string) (*Claims, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return _key, nil
	}, jwt.WithIssuedAt(), jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("no implement *Claims")
	}
	return claims, nil
}
