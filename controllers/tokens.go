package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
func GetTokenAttributes(token string) {
	tokenString := token
	claims := jwt.MapClaims{}
	type MyCustomClaims struct {
		email string `json:"foo"`
		jwt.StandardClaims
	}

	token1, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(tokenString *jwt.Token) (interface{}, error) {
		return []byte("email"), nil
	})
	fmt.Println(token1)
	fmt.Print(err)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}

type SignedDetails struct {
	Email string

	jwt.StandardClaims
}

func ValidateToken1(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
