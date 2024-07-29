package utils

import (
	// "errors"
	// "remind/todo/models"
	// "remind/todo/repository"
	"unicode"
	// "github.com/golang-jwt/jwt/v4"
)

// func ValidateToken(authorization, tokenType string) (models.Claims, error) {
// 	var claim models.Claims
// 	tokenSplit := strings.Split(authorization, " ")
// 	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
// 		return claim, errors.New("missing token")
// 	}

//		token, err := jwt.ParseWithClaims(tokenSplit[1], &claim, func(token *jwt.Token) (interface{}, error) {
//			return []byte(configs.Yaml.JWT.Secret), nil
//		})
//		if err != nil || !token.Valid {
//			return claim, errors.New("token is invalid")
//		}
//		if claim.Type != tokenType {
//			return claim, errors.New("invalid token")
//		}
//		err = repository.RequiredAuth(claim)
//		if err != nil {
//			return claim, err
//		}
//		return claim, nil
//	}
func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var number, upper, lower, special bool
	for _, char := range password {
		switch {
		case ' ' < char && char > '~':
			return false
		case unicode.IsDigit(char):
			number = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		}
	}
	return number && upper && lower && special
}
