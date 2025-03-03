package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Bat", Price: 1200},
	{ID: 2, Name: "Ball", Price: 200},
	{ID: 3, Name: "Stump", Price: 500},
}

func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func createProducts(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func main() {

	r := gin.Default() //create new gin router

	// r.GET("/ping", func(c *gin.Context) { // this is handler fn , define what will happen if someone visits /ping
	// 	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	// })

	r.GET("/Products", getProducts)
	r.POST("/createProduct", createProducts)

	r.Run(":8000")

}
