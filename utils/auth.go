package utils

import (
	"crypto/rand"
	"fmt"
	random "math/rand"
	"net/http"
	"strconv"
	"time"

	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-martini/martini"
)

var SECRET = []byte("Be aware with FrenchFarmer")

func GenerateToken(id int) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		Id:        strconv.Itoa(id),
		Issuer:    "FrenchFarmer",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}

func HttpSetToken(params martini.Params) (int, string) {
	if _, ok := params["id"]; !ok {
		return http.StatusBadRequest, "You must provide id and password"
	}
	var err error
	if id, err := strconv.Atoi(params["id"]); err == nil {
		token, err := GenerateToken(id)
		if err != nil {
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusOK, token
	}
	return http.StatusInternalServerError, err.Error()
}

// Authorized checks if JWT token is in header.
// If not, it will stop request and disallow to access data.
func Authorized(res http.ResponseWriter, r *http.Request) {
	token, _, err := GetToken(r)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad token, is invalid"))
		res.Write([]byte(err.Error()))
	} else if !token.Valid {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad token, is invalid"))
	}
}

// GetToken returns token found in request + claims. At this time we only
// need StandardClaims that provides Id and ExpirationDate.
func GetToken(r *http.Request) (*jwt.Token, *jwt.StandardClaims, error) {
	m := request.MultiExtractor{
		request.AuthorizationHeaderExtractor,
	}
	claims := &jwt.StandardClaims{}

	token, err := request.ParseFromRequestWithClaims(r, m, claims, func(t *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	return token, claims, err
}

func GenerateRandowStringByLenght(n int) (string, error) {

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func NewUuid() string {
	random.Seed(int64(time.Now().Nanosecond()))
	return regexp.
		MustCompile("[xy]").
		ReplaceAllStringFunc("xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx", func(s string) string {
			v := random.Intn(16)
			if s == "x" {
				return fmt.Sprintf("%x", v)
			}
			return fmt.Sprintf("%x", v&0x3|0x8)
		})
}
