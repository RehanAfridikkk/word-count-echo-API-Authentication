package main 


import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/RehanAfridikkk/API-Authentication/controller"
	"github.com/RehanAfridikkk/API-Authentication/structure"
	"github.com/golang-jwt/jwt/v5"
)



func main() {
	e := echo.New()


	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.POST("/login", controller.Login)

	
	e.GET("/", controller.Accessible)

	
	result := e.Group("/restricted")

	
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(structure.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	result.Use(echojwt.WithConfig(config))
	result.GET("", controller.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}