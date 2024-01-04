package helpers

import (
	"fmt"
	"skeleton-svc/helpers/models"
	"strings"

	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	authorizationSecretKey = ""
	// secretKey              = []byte(authorizationSecretKey)
)

// Authorize ..
// func Authorize(header models.RequestHeader) (pModels.Customer, bool) {

// 	customer := pModels.Customer{}

// 	// decode auth token
// 	name, LocationCode, decodeAuthTokenSuccessful := decodeAuthToken(header)
// 	if !decodeAuthTokenSuccessful {
// 		return customer, false
// 	}
// 	if name == "" {
// 		logs.Error("username not found")
// 		return customer, false
// 	} else {
// 		customer.Name = name
// 		customer.LocationCode = LocationCode
// 	}

// 	return customer, true
// }

// DecodeAuthToken ..
func DecodeAuthToken(header models.RequestHeader) (string, string, bool) {
	authToken, getAuthTokenSuccessful := GetAuthToken(header)
	if !getAuthTokenSuccessful {
		return "", "", false
	}

	// decode auth token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(authorizationSecretKey), nil
	})
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to parse token: %v", err))
		return "", "", false
	}
	if !token.Valid {
		logs.Error("Token is not valid")
		return "", "", false
	}
	IDClaims, IDExist := claims["username"]

	if !IDExist {
		logs.Error("Token does not contain username information")
		return "", "", false
	}
	LocationCode, LocationCodeExist := claims["location_code"]
	if !LocationCodeExist {
		logs.Error("Token does not contain locationCode information")
		return "", "", false
	}

	fmt.Println("====== claims ======")
	fmt.Println(claims)
	fmt.Println("====== claims ======")

	return fmt.Sprintf("%v", IDClaims), fmt.Sprintf("%v", LocationCode), true
}

// GetAuthToken ..
func GetAuthToken(header models.RequestHeader) (string, bool) {
	if strings.TrimSpace(header.Authorization) == "" {
		logs.Error("Missing authorization token")
		return "", false
	}

	if !strings.HasPrefix(header.Authorization, "Bearer ") {
		logs.Error("Invalid bearer authorization token")
		return "", false
	}

	token := strings.TrimPrefix(header.Authorization, "Bearer ")

	return token, true
}

// GenerateToken ..
// func GenerateToken(customer pModels.Customer) (string, int64, error) {
// 	PrintHeader()
// 	nowTime := time.Now()
// 	// expireTime := nowTime.AddDate(0, 0, 1)
// 	expireTime := nowTime.Add(time.Hour * 6)
// 	expiredAt := expireTime.Unix()
// 	claims := Claims{
// 		customer.Name,
// 		customer.LocationCode,
// 		jwt.StandardClaims{
// 			ExpiresAt: expiredAt,
// 			Issuer:    "Indomarco",
// 		},
// 	}

// 	fmt.Printf("%+v", claims)

// 	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	token, err := tokenClaims.SignedString(secretKey)
// 	if err != nil {
// 		fmt.Println("Error  : ", err)
// 		return token, expiredAt, err
// 	}

// 	return token, expiredAt, nil
// }

// Authenticate ...
func Authenticate(encryptedPassword, reqPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(reqPassword)); err != nil {
		fmt.Println("Password hashes are not same. error : ", err)
		return false
	}
	return true
}

// HasingPassword ...
func HasingPassword(reqPassword string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
