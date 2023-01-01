package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("xtek_secret_key")

// Claims ...
type Claims struct {
	UserID string
	jwt.StandardClaims
}

// GenerateToken ...
func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// VerifyToken ...
func VerifyToken(tknStr string) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	if !tkn.Valid {
		return "", ErrInvalidToken
	}
	return claims.UserID, nil
}

// ErrInvalidToken ...
var ErrInvalidToken = errors.New("token is invalid")
