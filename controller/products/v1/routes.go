package productV1

import (
	"github.com/labstack/echo/v4"
	"github.com/aydinsercan/basketapi/controller/products"
	mid "github.com/aydinsercan/basketapi/middleware"
)

type (
	Router struct {
		product.Router
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {
	g.Use(mid.Auth(false))
	g.GET("", r.List)
}
