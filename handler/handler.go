package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Health Return ok if server is healthy
func Health(c *gin.Context) {
	fmt.Println("Service healthy")
}

// StartLoan Start new loan.
func StartLoan(c *gin.Context) {
	fmt.Println("Starting new loan")
}

// GetBalance Get remaining amount of the loan
func GetBalance(c *gin.Context) {
	fmt.Println("Getting balance")
}

// AddBalance Attach installment payment to the existing loan
func AddBalance(c *gin.Context) {
	fmt.Println("Adding balanace")
}
