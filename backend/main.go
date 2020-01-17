package main

import (
	"errors"
	"firebase-authentication-with-react-and-go/backend/domain/model"
	rc "firebase-authentication-with-react-and-go/backend/router/context"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"firebase-authentication-with-react-and-go/backend/datastore"
	m "firebase-authentication-with-react-and-go/backend/router/middleware"
)

func main() {
	db, err := datastore.NewDB()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	//--- restricted api ---//
	g := e.Group("api", m.Auth())
	{
		g.POST("/users",  wrapCustomContext(func(c *rc.Context) error {
			var params model.User
			user := model.User{}

			if err := c.Bind(&params); !errors.Is(err, nil) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}

			// Check to see if the user already registered
			err = db.Where(&model.User{UUID: params.UUID}).First(&user).Error
			if !errors.Is(err, nil) && !gorm.IsRecordNotFoundError(err) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			if user.UUID != "" {
				return c.JSON(http.StatusBadRequest, "The user already registered")
			}

			// Create a new user when not registered
			user = params
			err = db.Create(&user).Error
			if !errors.Is(err, nil) {
				return err
			}
			return c.JSON(http.StatusCreated, user)
		}))
	}

	err = e.Start(":8080")
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func wrapCustomContext(fn func(c *rc.Context) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return fn(ctx.(*rc.Context))
	}
}
