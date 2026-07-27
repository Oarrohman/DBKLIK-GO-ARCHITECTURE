// Package controllers receives HTTP requests, validates them, calls the
// application layer, and returns a response. Controllers stay thin —
// no business logic lives here.
package controllers

import (
	"strconv"

	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/dto"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/application/services"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/presentation/http/requests"
	"github.com/dbklik/dbklik-go-architecture/modules/inventory/presentation/http/responses"
	apperrors "github.com/dbklik/dbklik-go-architecture/pkg/errors"
	"github.com/dbklik/dbklik-go-architecture/pkg/response"
	"github.com/dbklik/dbklik-go-architecture/pkg/validator"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *services.InventoryService
}

func NewProductController(service *services.InventoryService) *ProductController {
	return &ProductController{service: service}
}

func (ctl *ProductController) Create(c *gin.Context) {
	var req requests.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperrors.InvalidInput(err.Error()))
		return
	}
	if err := validator.Struct(req); err != nil {
		response.Error(c, apperrors.InvalidInput(err.Error()))
		return
	}

	out, err := ctl.service.CreateProduct(dto.CreateProductInput{
		Name:       req.Name,
		SKU:        req.SKU,
		PriceCents: req.PriceCents,
		Quantity:   req.Quantity,
	})
	if err != nil {
		response.Error(c, err)
		return
	}

	responses.ProductCreated(c, out)
}

func (ctl *ProductController) Get(c *gin.Context) {
	out, err := ctl.service.GetProduct(c.Param("id"))
	if err != nil {
		response.Error(c, err)
		return
	}
	responses.Product(c, "product retrieved", out)
}

func (ctl *ProductController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	out, err := ctl.service.ListProducts(page, pageSize)
	if err != nil {
		response.Error(c, err)
		return
	}
	responses.ProductList(c, out)
}

func (ctl *ProductController) AdjustStock(c *gin.Context) {
	var req requests.AdjustStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperrors.InvalidInput(err.Error()))
		return
	}

	if err := ctl.service.AdjustStock(c.Param("id"), req.Delta); err != nil {
		response.Error(c, err)
		return
	}

	out, err := ctl.service.GetProduct(c.Param("id"))
	if err != nil {
		response.Error(c, err)
		return
	}
	responses.Product(c, "stock adjusted", out)
}

func (ctl *ProductController) Archive(c *gin.Context) {
	if err := ctl.service.ArchiveProduct(c.Param("id")); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "product archived", nil)
}
