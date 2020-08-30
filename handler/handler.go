package handler

import (
	"fmt"
	"loanprocessing/custom"
	"loanprocessing/loan"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Health Return ok if server is healthy
func Health(c *gin.Context) {
	fmt.Println("Service healthy")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Service is healthy!"})
}

// StartLoan Start new loan.
func StartLoan(c *gin.Context) {
	fmt.Println("Starting new loan")
	l := loan.GetLoanInstance(true)
	e := c.BindJSON(l)
	if e != nil {
		msgs := customErrMsg(e)
		sendErrorResponse(c, http.StatusBadRequest, "StartLoan :: "+msgs)
		return
	}

	l.Start()
	fmt.Println("Payment:", l.Principal, "Rate :", l.Rate, "Start Date:", time.Time(l.StartDate))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Loan Started Successfully"})
}

// AddPayment Get remaining amount of the loan
func AddPayment(c *gin.Context) {
	fmt.Println("Adding Payment")
	l := loan.GetLoanInstance(false)
	if l.Principal == 0 {
		sendErrorResponse(c, http.StatusNotFound, "AddPayment :: Loan is not started")
		return
	}
	installment := loan.NewInstallment()
	e := c.BindJSON(&installment)
	if e != nil {
		msgs := customErrMsg(e)
		sendErrorResponse(c, http.StatusBadRequest, "AddPayment :: "+msgs)
		return
	}

	e = l.AddPayment(installment)
	if e != nil {
		sendErrorResponse(c, http.StatusBadRequest, "AddPayment :: "+e.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Payment Accepted"})
}

// GetBalance Attach installment payment to the existing loan
func GetBalance(c *gin.Context) {
	fmt.Println("Getting balanace")
	l := loan.GetLoanInstance(false)
	if l.Principal == 0 {
		sendErrorResponse(c, http.StatusBadRequest, "GetBalance :: Loan is not started")
		return
	}
	params := c.Request.URL.Query()
	dateStr := params.Get("date")
	if len(dateStr) == 0 {
		sendErrorResponse(c, http.StatusBadRequest, "GetBalance :: date is required in yyyy-mm-dd format")
		return
	}
	date, e := custom.Parse(dateStr)
	if e != nil {
		sendErrorResponse(c, http.StatusBadRequest, "GetBalance :: provide date as yyyy-mm-dd (2006-01-02)")
		return
	}
	bal, err := l.GetBalance(date)
	fmt.Println("Balance : ", bal)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "GetBalance : "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Balance": bal})
}

func sendErrorResponse(c *gin.Context, status int, e string) {
	// fmt.Println("ERROR ::", e)
	c.JSON(status, gin.H{"status": status, "message": e})
	c.Abort()
}
