package context

import (
	fa "firebase.google.com/go/auth"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	Token *fa.Token
}
