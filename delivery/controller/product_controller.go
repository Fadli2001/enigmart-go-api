package controller

import (
	"enigmart-api/model"
	"enigmart-api/usecase"
	"net/http"

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
		product := pc.prodUseCase.CreateNewProduct(newProduct)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    product,
		})
	}
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	products := pc.prodUseCase.GetAllProduct()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    products,
	})
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var newProduct *model.Product
	err := ctx.BindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		productRes := pc.prodUseCase.UpdateProduct(*newProduct)
		if (productRes == model.Product{}) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Id tidak ditemukan",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"data":    productRes,
			})
		}
	}
}

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	responseUc := p.prodUseCase.DeleteProduct(idProduct)
	if responseUc {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak ditemukan",
		})
	}
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
	idProduct := ctx.Param("id")
	responseUc := p.prodUseCase.GetProductById(idProduct)
	if (responseUc == model.Product{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id tidak ditemukan",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    responseUc,
		})
	}
}

func NewProductController(router *gin.Engine, prodUseCase usecase.ProductUseCase) *ProductController {
	newProdController := ProductController{
		router:      router,
		prodUseCase: prodUseCase,
	}

	router.POST("/product", newProdController.CreateNewProduct)
	router.GET("/product", newProdController.GetAllProduct)
	router.PUT("/product", newProdController.UpdateProduct)
	router.DELETE("product/:id",newProdController.DeleteProduct)
	router.GET("product/:id",newProdController.GetProductById)
	return &newProdController
}
