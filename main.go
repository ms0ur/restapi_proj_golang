package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Categories []string `json:"categories"`
}

var products = []Product{
	{Id: "1", Name: "CPU", Price: 1000.0, Categories: []string{"Processors"}},
	{Id: "2", Name: "Motherboard", Price: 500.0, Categories: []string{"Motherboards"}},
	{Id: "3", Name: "RAM", Price: 200.0, Categories: []string{"Memory"}},
	{Id: "4", Name: "HDD", Price: 300.0, Categories: []string{"Storage"}},
	{Id: "5", Name: "GPU", Price: 1500.0, Categories: []string{"Graphics Cards"}},
}

func main() {
	r := gin.Default()

	r.GET("/products", getAllProducts)

	r.GET("/products/:id", getProductById)

	r.GET("/categories/:category", getProductsByCategory)

	r.POST("/products/new", createProduct)

	r.PUT("/products/:id", updateProductById)

	r.DELETE("/products/:id", deleteProductById)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProductById(c *gin.Context) {
	id := c.Param("id")
	for _, prd := range products {
		if prd.Id == id {
			c.JSON(http.StatusOK, prd)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func getProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	var prds []Product
	for _, prd := range products {
		for _, cat := range prd.Categories {
			if cat == category {
				prds = append(prds, prd)
				break
			}
		}
	}
	if len(prds) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found in this category"})
		return
	}
	c.JSON(http.StatusOK, prds)
}

func createProduct(c *gin.Context) {
	var prd Product
	if err := c.ShouldBindJSON(&prd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	products = append(products, prd)
	c.JSON(http.StatusCreated, prd)
}

func updateProductById(c *gin.Context) {
	id := c.Param("id")
	for i, prd := range products {
		if prd.Id == id {
			var updatedPrd Product
			if err := c.ShouldBindJSON(&updatedPrd); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			products[i] = updatedPrd
			c.JSON(http.StatusOK, updatedPrd)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func deleteProductById(c *gin.Context) {
	id := c.Param("id")
	for i, prd := range products {
		if prd.Id == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
