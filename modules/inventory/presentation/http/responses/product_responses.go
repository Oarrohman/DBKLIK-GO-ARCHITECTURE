// Package responses builds the inventory module's HTTP responses on top
// of the shared response envelope, keeping controllers free of formatting.
package responses

import (
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/presentation/http/resources"
	"github.com/dbklik/dbklik-go-architecture/pkg/response"
	"github.com/gin-gonic/gin"
)

func Product(c *gin.Context, message string, out dto.ProductOutput) {
	response.OK(c, message, resources.FromOutput(out))
}

func ProductCreated(c *gin.Context, out dto.ProductOutput) {
	response.Created(c, "product created", resources.FromOutput(out))
}

func ProductList(c *gin.Context, list dto.ListOutput) {
	response.Paginated(c, "products retrieved", resources.FromOutputList(list.Items), gin.H{
		"total": list.Total,
		"page":  list.Page,
	})
}
