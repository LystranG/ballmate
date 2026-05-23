package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ballmate/config"
)

// GenerateToken 签发 JWT token
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(config.TokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

// ParseToken 解析 JWT token，返回 userID
func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrTokenInvalidClaims
	}

	userID := uint(claims["user_id"].(float64))
	return userID, nil
}
