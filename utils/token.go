package utils

import (
	"log"
	"os"
	"time"

	"harry/auth_system/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens generates both teh detailed token and refresh token
func GenerateAllTokens(email string, name string, uid string) (signedToken string, err error) {
	claims := &models.SignedDetails{
		Email: email,
		Name:  name,
		Id:    uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}
