package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	// Name string    `json:"name"`
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func JwtMakeToken(id uuid.UUID) string {
	claims := &jwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("d8789236abac4eb24b4a5fc301be5053347ac4c9f91d3f2ac525336e7967f4c3"))
	if err != nil {
		return err.Error()
	}

	return t
}

func JwtUserID(token *jwt.Token) uuid.UUID {
	user := token
	claims := user.Claims.(*jwtCustomClaims)

	return claims.ID
}

func JwtConfig() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("d8789236abac4eb24b4a5fc301be5053347ac4c9f91d3f2ac525336e7967f4c3"),
	}

	return config
}

func JwtMiddlewareSet() echo.MiddlewareFunc {
	config_set := echojwt.WithConfig(JwtConfig())

	return config_set
}
