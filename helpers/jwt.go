package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var APPLICATION_NAME = "MY-GARM"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")


type MyClaims struct {
    jwt.StandardClaims
    Id 		 uint `json:"id"`
    Email    string `json:"Email"`
	Role 	 string `json:"role"`
}

func GenerateToken(id uint, email string, role string) (string, error) {

	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
			Issuer:    APPLICATION_NAME,
		},
		Id: id,
		Email: email,
		Role: role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}


func VerifyToken(c *gin.Context) (interface{}, error){
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	tokenString := strings.Split(headerToken, " ")[1]

	token,_ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}