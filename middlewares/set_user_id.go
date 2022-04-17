package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func SetUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString.Value == "" {
			return next(c)
		}
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// TODO 修正する
			return []byte("secret"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		var id int

		if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id = int(claim["id"].(float64))
		}

		c.Set("user_id", id)

		return next(c)
	}
}
