package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrInvalidClaims = errors.New("invalid claims")
)

func GenerateJWT(claims jwt.MapClaims) (string, error) {
	now := time.Now().UTC()

	ttlMinutes := GetEnvAsInt("AUTH_SERVER_JWT_TTL_MINUTES", 60)
	ttl := time.Minute * time.Duration(ttlMinutes)
	exp := now.Add(ttl)

	claims["iss"] = GetEnv("AUTH_SERVER_JWT_ISSUER", "devs")
	claims["iat"] = now.Unix()
	claims["exp"] = exp.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := GetEnv("AUTH_SERVER_JWT_SECRET", "development")

	return token.SignedString([]byte(secret))
}

func Verify(token string) (jwt.MapClaims, error) {
	parser := &jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Alg()}}
	t, err := parser.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}

		secret := GetEnv("AUTH_SERVER_JWT_SECRET", "development")
		return []byte(secret), nil
	})

	if err != nil || !t.Valid {
		if err == nil {
			err = ErrInvalidToken
		}
		return nil, err
	}

	mc, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return mc, nil
}
