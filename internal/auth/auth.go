package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/evlbit/notesmd/internal/env"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userID int) (string, error) {
	exp := time.Hour * time.Duration(env.Vars.JWTExpHours)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": strconv.Itoa(userID),
			"exp":     time.Now().Add(exp).Unix(),
		},
	)

	tokenStr, err := token.SignedString([]byte(env.Vars.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(token string) (int, error) {
	tokenStrc, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(env.Vars.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !tokenStrc.Valid {
		return 0, fmt.Errorf("token invalid")
	}

	claims := tokenStrc.Claims.(jwt.MapClaims)
	tokenStr := claims["user_id"].(string)

	userId, err := strconv.Atoi(tokenStr)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
