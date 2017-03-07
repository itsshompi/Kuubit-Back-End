package common

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
)

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	privKeyPath = "keys/privKey.key"
	pubKeyPath  = "keys/pubKey.pem"
)

var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys0]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys1]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys2]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys3]: %s\n", err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(id, role string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Get token from request
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"El Token a expirado, obtenga un nuevo Token",
					401,
				)
				return

			default:
				DisplayAppError(w,
					err,
					"Error leyendo el Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error leyendo el Token!",
				500)
			return
		}

	}
	if token.Valid {
		// Set user name to HTTP context
		claims := token.Claims.(jwt.MapClaims)
		context.Set(r, "id", claims["id"])
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Token invalido!",
			401,
		)
	}
}
