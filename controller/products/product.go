package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mid "github.com/aydinsercan/basketapi/middleware"
	ProductModel "github.com/aydinsercan/basketapi/model/product"
)

type (
	Router struct {
		mid.Request
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {}

func (r *Router) List(c echo.Context) error {
	mod := ProductModel.Product{}
	data, err := mod.List()
	if err != nil {
		return r.Send(c, http.StatusInternalServerError, "Cannot get data", err.Error(), nil)
	}
	return r.Send(c, http.StatusOK, "OK", "", data)
}
