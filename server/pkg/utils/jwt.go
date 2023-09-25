package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

type JwtUser struct {
	UserID uuid.UUID
}

func GenerateJwt(userId uuid.UUID) (string, error) {
	exp := time.Now().Add(60*60*24*7*time.Second)
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: exp},
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
		},
	}
	
	utoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := utoken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeJwt(token string) (uuid.UUID, error) {
	decoded, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		} 

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	log.Printf("decoding uuid -> decoded: %v", decoded)
	log.Printf("decoding uuid -> error: %v", err)

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := decoded.Claims.(*Claims);
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid claims type")
	}
	
	return claims.UserID, nil
}