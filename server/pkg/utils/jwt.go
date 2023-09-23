package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt() (string, error) {
	exp := time.Now().Add(60*60*24*7*time.Second)
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: exp},
		IssuedAt: &jwt.NumericDate{Time: time.Now()},
	}
	
	utoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := utoken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeJwt(token string) (bool, error) {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		} 

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return false, err
	}
	
	return true, nil
}