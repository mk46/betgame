package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("GoLinuxCloudKey")

func generateJWT(phone string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		err = fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func validateToken(authtoken string) (phone string, err error) {

	token, err := jwt.Parse(authtoken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("failed to parse jwt token")
		}
		return sampleSecretKey, nil

	})

	if token == nil {
		err = errors.New("empty token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims != nil {
			phone = fmt.Sprint(claims["phone"])
		}

	}

	return
}
