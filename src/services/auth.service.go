package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJwt(userID uuid.UUID) (tokenString string, err error) {

	// Create the Claims
	claims := JwtClaims{
		"someone@example.com", // This value specified on puepose, to be changed later
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "movie-api",
			Subject:   userID.String(),
			Audience:  []string{"movie-api"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(jwtSecret)

	return
}

func GetTokenClaims(signedToken string) (claims *JwtClaims, err error) {

	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = errors.New("token expired")
		return nil, err
	}

}

func ValidateToken(signedToken string) (err error) {

	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Email, claims.RegisteredClaims.Issuer)
		return nil
	} else {
		err = errors.New("token expired")
		return
	}

}

func ValidateAdminToken(signedToken string) (err error) {

	token, err := jwt.ParseWithClaims(signedToken, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		if isAdmin(claims) {
			fmt.Printf("%v %v", claims.Email, claims.RegisteredClaims.Issuer)
			return nil
		}
		err = errors.New("not an admin")
		return

	} else {
		err = errors.New("token expired")
		return
	}

}

func CheckUser(context *gin.Context, userID string) error {

	//get user from token
	bearerToken := context.GetHeader("Authorization")
	accessToken := strings.Split(bearerToken, "Bearer ")[1]
	claims, err := GetTokenClaims(accessToken)

	if err != nil {
		return err
	}

	if claims.Subject != userID {
		return errors.New("user from token is not the same as the user from the request")
	}

	return nil

}

func isAdmin(claims *JwtClaims) bool {

	var user, _ = GetUserByID(claims.Subject)
	print(user.Role)
	if user.Role == "admin" {
		return true
	}
	return false
}
