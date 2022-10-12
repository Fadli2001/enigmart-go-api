package controller

import (
	"enigmart-api/model"
	"enigmart-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router      *gin.Engine
	prodUseCase usecase.ProductUseCase
}

func (pc *ProductController) CreateNewProduct(ctx *gin.Context) {
	var newProduct *model.Product
	err := ctx.BindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		pc.prodUseCase.CreateNewProduct(newProduct)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    newProduct,
		})
	}
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	responseUc, err := p.prodUseCase.GetProductById(idProduct)
	if (responseUc == model.Product{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    responseUc,
		})
	}
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	page,_ := strconv.Atoi(ctx.Query("page"))
	totalRows,_ := strconv.Atoi(ctx.Query("totalRows"))	
	if page == 0 || totalRows == 0 {
		page = 1
		totalRows = 5
	}
	products, err := pc.prodUseCase.GetAllProduct(page, totalRows)	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    products,
			"page" : page,
			"totalRows" : totalRows,						
		})
	}
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var newProduct *model.Product
	err := ctx.BindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := pc.prodUseCase.UpdateProduct(*newProduct)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    newProduct,
			})
		}
	}
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	err := p.prodUseCase.DeleteProduct(idProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}

func NewProductController(router *gin.Engine, prodUseCase usecase.ProductUseCase) *ProductController {
	newProdController := ProductController{
		router:      router,
		prodUseCase: prodUseCase,
	}	

	product := router.Group("/product")
	product.POST("", newProdController.CreateNewProduct)
	product.GET("/:id", newProdController.GetProductById)
	product.GET("", newProdController.GetAllProduct)
	product.PUT("", newProdController.UpdateProduct)
	product.DELETE("/:id", newProdController.DeleteProduct)
	return &newProdController
}
