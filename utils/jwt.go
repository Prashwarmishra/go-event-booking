package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "SECRET_KEY"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}	

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func (jwtToken *jwt.Token) (any, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid signing method!")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parse auth token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
