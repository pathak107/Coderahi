package auth

import (
	"errors"
	"fmt"

	"github.com/MicahParks/keyfunc"
	jwt "github.com/golang-jwt/jwt/v4"
)

type AuthInfo struct {
	UserID      string
	UserEmail   string
	AccessToekn string
	UserRole    string
}

type JWTService struct {
	JwksURL string
	Jwks    *keyfunc.JWKS
}

func NewJWTService(jwksURL string) (*JWTService, error) {
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return &JWTService{}, err
	}
	return &JWTService{
		JwksURL: jwksURL,
		Jwks:    jwks,
	}, nil
}

func (j *JWTService) VerifyJWTToken(jwtB64 string) (*AuthInfo, error) {
	// Create the JWKS from the resource at the given URL.
	// Parse the JWT.
	token, err := jwt.Parse(jwtB64, j.Jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return &AuthInfo{}, nil
	} else {
		return &AuthInfo{}, errors.New("error in verifying token")
	}

}
