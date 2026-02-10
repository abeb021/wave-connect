package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	//UserID string  `json:"user_id"`
	//Email  string  `json:"email"`
	jwt.RegisteredClaims
}

type AuthService struct{
	secret []byte
}

func NewAuthService(key string) *AuthService {
	return &AuthService{secret: []byte(key)}
}

func (a *AuthService) GenerateToken(userID, email string) (string, error){
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "auth-service",
			Subject: userID,
			Audience: []string{"wave-connect"},
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			ID: "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.secret)
}

func (a *AuthService) ValidateToken(tokenString string) (bool, error){
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return a.secret, nil
	})

	if err != nil{
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}
