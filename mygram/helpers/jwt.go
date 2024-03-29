package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "RAHASIA"

func GenerateToken(id int, email string) string {
	// time := time.Now()
	claims := jwt.MapClaims{
		"id": id,
		// "email": email,
		// "time":  time,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Processed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errResponse
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
