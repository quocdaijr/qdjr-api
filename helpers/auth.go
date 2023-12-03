package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthHelper struct{}

func (_ AuthHelper) GenerateToken(sub string, signature string, ttl int) (string, error) {
	if ttl == 0 {
		ttl = 60
	}
	claims := jwt.MapClaims{}
	claims["iss"] = BaseHelper{}.GetEnv("APP_NAME", "qdjr-api")
	claims["jti"] = uuid.NewString()
	claims["sub"] = sub
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(ttl)).Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(signature))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (_ AuthHelper) VerifyToken(token string, signature string) (*jwt.Token, error) {
	verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signature), nil
	})
	if err != nil {
		return nil, err
	}
	return verifiedToken, nil
}

func (_ AuthHelper) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func (_ AuthHelper) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
