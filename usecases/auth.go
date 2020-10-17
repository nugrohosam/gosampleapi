package usecases

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// AuthBasic ...
func AuthBasic(emailOrUsername, password string) (string, error) {
	user := userRepo.FindByEmailOrUsername(emailOrUsername)

	if len(user.Username) == 0 || len(user.Email) == 0 {
		return "", errors.New("Cannot find user, username or email")
	}

	if isPasswordValid := helpers.CompareHash([]byte(user.Password), password); isPasswordValid {
		return "", errors.New("Cannot find user, password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    user.Username,
		"email":       user.Email,
		"expiredTime": time.Now().AddDate(1, 0, 0),
	})

	bytedString := helpers.GetBytedSecret()
	tokenString, err := token.SignedString(bytedString) // always use byted string
	if err != nil {
		return "", errors.New("Cannot make token")
	}

	return tokenString, nil
}

// AuthorizationValidation ...
func AuthorizationValidation(tokenString string) error {
	token, err := jwt.Parse(tokenString, validateToken)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else if err != nil {
		return errors.New("Cannot validate auth token")
	} else {
		return errors.New("Cannot validate")
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	bytedString := helpers.GetBytedSecret()

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return bytedString, nil
}