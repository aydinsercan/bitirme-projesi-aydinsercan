package cartV1

import (
	cart "github.com/aydinsercan/basketapi/controller/carts"
	mid "github.com/aydinsercan/basketapi/middleware"
	"github.com/labstack/echo/v4"
)

type (
	Router struct {
		cart.Router
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {
	g.Use(mid.Auth(false))
	g.GET("", r.List)
	g.POST("", r.Insert)
	g.PUT("/:id", r.Update)
	g.DELETE("/:id", r.Delete)
	g.DELETE("/orders", r.Order)
}
