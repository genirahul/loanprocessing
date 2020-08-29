package loan

import (
	"fmt"
	"loanprocessing/custom"
	"math"
	"sync"
	"time"
)

// Loan stores details of the loan that is processed currently.
type Loan struct {
	Principal   float64     `json:"initialAmount"`
	Rate        float32     `json:"annualRate"`
	StartDate   custom.Date `json:"startDate"`
	Payments    []Installment
	Adjustments []adjustment
}

// Installment store load payment details.
type Installment struct {
	Amount float64     `json:"amount"`
	Date   custom.Date `json:"date"`
}

type adjustment struct {
	CurrPrincipal   float64
	PendingInterest float64
	Date            custom.Date
}

var l *Loan
var once sync.Once

// GetLoanInstance ...
func GetLoanInstance(resetLoan bool) *Loan {
	once.Do(func() {
		l = &Loan{}
	})
	if resetLoan {
		fmt.Println("Just for fun")
	}
	return l
}

// NewInstallment returns new instance of Installment struct
func NewInstallment() Installment {
	return Installment{}
}

func (l *Loan) reset() {
	l.Principal = 0
	l.Rate = 0
	l.StartDate = custom.Date(time.Time{})
	l.Payments = []Installment{}
	l.Adjustments = []adjustment{}
}

// IsLoanStarted returns true is principal is greater than 0
func (l *Loan) IsLoanStarted() bool {
	if l.Principal > 0 {
		return true
	}

	return false
}

// Start loan by adding initial principal and start date to adjustments
func (l *Loan) Start() {
	adj := adjustment{
		CurrPrincipal:   l.Principal,
		Date:            l.StartDate,
		PendingInterest: 0,
	}
	l.Adjustments = append(l.Adjustments, adj)
	fmt.Println("Loan started")
}

// AddPayment ...
func (l *Loan) AddPayment(i Installment) error {
	// validate installment data
	err := ValidateInstallment(*l, i)
	if err != nil {
		return err
	}
	l.Payments = append(l.Payments, i)
	l.AdjustPrincipal(i)
	return nil
}

// GetBalance - Returns balance loan amount.
func (l *Loan) GetBalance(d custom.Date) (float64, error) {
	// validate d > l.StartDate
	err := ValidateGetBalanceDate(*l, d)
	if err != nil {
		return 0, err
	}
	// Get last adjustment data before date.
	adj := l.getLastAdjustmentAsOfDate(d)
	// calculate interest till date.
	interest := l.calculateInterest(adj.CurrPrincipal, adj.Date, d)
	fmt.Println("Interest : ", interest, "Current Principal : ", adj.CurrPrincipal)
	// Add principal and interest to get current balance.
	return adj.CurrPrincipal + adj.PendingInterest + interest, nil
}

// AdjustPrincipal - Decrease principal is amount paid is greater than total interest till date.
func (l *Loan) AdjustPrincipal(i Installment) {
	adj := l.getLastAdjustmentAsOfDate(i.Date)
	// calculate interest till date.
	interest := l.calculateInterest(adj.CurrPrincipal, adj.Date, i.Date)
	totalInterest := adj.PendingInterest + interest
	newAdj := adjustment{}
	if totalInterest < i.Amount {
		fmt.Println("Adjusting principal")
		newAdj.CurrPrincipal = adj.CurrPrincipal - (i.Amount - totalInterest)
		newAdj.PendingInterest = 0
	} else {
		newAdj.CurrPrincipal = adj.CurrPrincipal
		newAdj.PendingInterest = totalInterest - i.Amount
	}

	newAdj.Date = i.Date
	l.Adjustments = append(l.Adjustments, newAdj)
	fmt.Println("Adjustments :", l.Adjustments)
}

func (l *Loan) calculateInterest(currPrincipal float64, lastAdjDate, d custom.Date) float64 {
	dailyInterest := (float64(l.Rate) * currPrincipal) / (100 * 365)
	// Number of days since last principal adjustment.
	noOfDays := int(d.Sub(lastAdjDate).Hours() / 24)

	totalInterest := float64(noOfDays) * dailyInterest
	return math.Round(totalInterest*100) / 100
}

// Get last adjustment made to principal
func (l *Loan) getLastAdjustmentAsOfDate(d custom.Date) adjustment {
	adjustments := l.Adjustments
	for i := len(adjustments) - 1; i >= 0; i-- {
		adj := adjustments[i]
		if d.Sub(adj.Date) >= 0 {
			fmt.Println("Last Adjustment :", adj)
			return adj
		}
	}

	return adjustment{}
}

// getLastPaymentDate Returns last payment date.
func (l *Loan) getLastPaymentDate() custom.Date {
	payments := l.Payments
	if len(payments) > 0 {
		lastPayment := payments[len(payments)-1]
		return lastPayment.Date
	}
	return custom.Date(time.Time{})
}
