package cmd

import (
	"log"
	"wallet/helpers"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	err := r.Run(":" + helpers.GetEnv("PORT", "8080")) // listen and serve on
	if err != nil {
		log.Fatal(err)
	}
}
