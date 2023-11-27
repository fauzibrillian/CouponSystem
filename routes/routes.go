package routes

import (
	"Day25/features/coupon"
	"Day25/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, uc user.Handler, cc coupon.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	RouteUser(e, uc)
	RouteCoupon(e, cc)
}

func RouteUser(e *echo.Echo, uc user.Handler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}

func RouteCoupon(e *echo.Echo, cc coupon.Handler) {
	e.POST("/coupon", cc.Add(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
	e.GET("/coupon/all", cc.GetAllTanpa())
	e.GET("/coupon/user", cc.GetAll(), echojwt.JWT([]byte("$!1gnK3yyy!!!")))
}
