package main

import (
	"loanprocessing/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/loan/health", handler.Health)

	v1 := r.Group("/loan/v1.0/")
	{
		v1.GET("/start", handler.StartLoan)
		v1.GET("/add-payment", handler.AddBalance)
		v1.GET("/get-balance", handler.GetBalance)
	}

	r.Run("127.0.0.1:8080")
}
