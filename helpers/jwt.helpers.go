package helpers

import (
	"fmt"
	"net/http"
	"reflect"
	"skeleton-svc/constants"
	"skeleton-svc/helpers/models"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtSecret ..
var JwtSecret = []byte("")

// Claims ..
type Claims struct {
	Username     string `json:"username"`
	LocationCode string `json:"location_code"`
	jwt.StandardClaims
}

// JWT ..
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		res := models.Response{
			Meta: GetMetaResponse(constants.RcSuccess),
		}

		res.Meta = GetMetaResponse(constants.RcSuccess)
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			res.Meta = GetMetaResponse(constants.RcBadrequest)
		} else {
			_, err := ParseToken(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					res.Meta = GetMetaResponse(constants.RcUnauthorized)
				default:
					res.Meta = GetMetaResponse(constants.RcGeneralError)
				}
			}
		}

		if res.Meta != GetMetaResponse(constants.RcSuccess) {
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}
		c.Next()
	}
}

// ParseToken ..
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetIDFromClaims ..
func GetIDFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
