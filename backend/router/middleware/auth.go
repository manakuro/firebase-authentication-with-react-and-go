package middleware

import (
	"context"
	"net/http"

	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"

	rc "firebase-authentication-with-react-and-go/backend/router/context"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rcc := &rc.Context{
				Context: c,
				Token:   nil,
			}

			opt := option.WithCredentialsFile("/Users/kurodamanato/.keys/firebase-service-key-2.json")
			app, err := firebase.NewApp(context.Background(), nil, opt)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			auth, err := app.Auth(context.Background())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			header := c.Request().Header.Get(echo.HeaderAuthorization)
			idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
			token, err := auth.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			// Set context to token
			rcc.Token = token

			return next(rcc)
		}
	}
}
