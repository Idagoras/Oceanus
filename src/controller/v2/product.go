package v2

import "github.com/gin-gonic/gin"

func AddProduct(c *gin.Context) {
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")

	c.JSON(200, gin.H{
		"v2":    "AddProduct",
		"name":  name,
		"price": price,
	})
}
