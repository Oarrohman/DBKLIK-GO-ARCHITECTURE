// Package routes maps the inventory module's HTTP endpoints onto its
// controllers, kept separate from wiring (infrastructure/provider) so the
// route table itself is easy to review.
package routes

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/presentation/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAPI(api *gin.RouterGroup, ctl *controllers.ProductController) {
	products := api.Group("/inventory/products")
	{
		products.POST("", ctl.Create)
		products.GET("", ctl.List)
		products.GET("/:id", ctl.Get)
		products.PATCH("/:id/stock", ctl.AdjustStock)
		products.POST("/:id/archive", ctl.Archive)
	}
}
