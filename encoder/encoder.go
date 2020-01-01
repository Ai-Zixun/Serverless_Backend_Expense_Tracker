package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey []byte = []byte("SigningKey")

func EncodeJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	claims["usr"] = username
	claims["exp"] = time.Now().Add(time.Hour * 0).Unix()

	encoded, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Errorf("Encoder - Encoding Error: %s", err.Error())
		return "", err
	}

	return encoded, nil
}

func DecodeJWT(encoded string) (string, bool, error) {
	token, _ := jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		err := fmt.Errorf("Encoder - Decoding Error")
		return "", false, err
	}

	if exp, ok := claims["exp"].(float64); ok {
		if exp < float64(time.Now().Unix()) {
			err := fmt.Errorf("Encoder - Decode JWT Expired Error")
			return "", false, err
		}
	} else {
		err := fmt.Errorf("Encoder - Decode JWT Type Error")
		return "", false, err
	}

	if usr, ok := claims["usr"].(string); ok {
		return usr, true, nil
	}
	err := fmt.Errorf("Encoder - Decode JWT Type Error")
	return "", false, err
}

func main() {
	fmt.Println("Encoder")
	token, _ := EncodeJWT("mypassword")
	fmt.Println(token)

	username, _, _ := DecodeJWT(token)
	fmt.Println(username)

}
