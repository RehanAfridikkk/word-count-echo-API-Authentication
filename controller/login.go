package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/RehanAfridikkk/API-Authentication/structure"
	
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "Rehan" || password != "Bahi!" {
		return echo.ErrUnauthorized
	}

	claims := &structure.JwtCustomClaims{
		"Rehan Afrdi",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
