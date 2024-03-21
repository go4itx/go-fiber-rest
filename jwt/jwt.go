package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GetExpTime ...
func GetExpTime(day int64) int64 {
	return time.Now().Add(time.Hour * time.Duration(24*day)).Unix()
}

// CreateToken return token
func CreateToken(aud any, exp int64, secret string) (token string, err error) {
	claims := jwt.MapClaims{
		// "iss": iss,
		"aud": aud,
		"iat": time.Now().Unix(),
		"exp": exp,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

// ParseToken ...
func ParseToken(tokenString string, secret string) (claims jwt.MapClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	return
}
