package product

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
	"net/http"
)

// @Summary Get all products
// @Description Retrieves all products from the database.
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200  {object}  []models.Product  "Products retrieved successfully"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /product/all [get]
func HandleGetAllProducts(c *gin.Context) {
	products, err := database.GetAllProducts(c.Request.Context())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve products")
		return
	} else if len(products) == 0 || products == nil {
		utils.RespondWithError(c, http.StatusNotFound, "No products found")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "Products retrieved successfully", products)
}

// @Summary Create a product
// @Description Creates a new product in the database.
// @Tags products
// @Accept  json
// @Produce  json
// @Param   product  body  models.Product  true  "Product object"
// @Success 201  {object}  models.Product  "Product created successfully"
// @Failure 400  {object}  nil  "Invalid request payload"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /product/create [post]
func HandleCreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdProduct, err := database.CreateProduct(c.Request.Context(), product)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, "Product created successfully", createdProduct)
}

// @Summary Create multiple products
// @Description Creates multiple products in the database.
// @Tags products
// @Accept  json
// @Produce  json
// @Param   products  body  []models.Product  true  "Array of product objects"
// @Success 201  {object}  []models.Product  "Products created successfully"
// @Failure 400  {object}  nil  "Invalid request payload"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /product/createMany [post]
func HandleCreateManyProducts(c *gin.Context) {
	var products []models.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdProducts, err := database.CreateManyProducts(c.Request.Context(), products)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create products")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, "Products created successfully", createdProducts)
}
