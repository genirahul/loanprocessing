package loan

import (
	"fmt"
	"time"
)

// Loan stores details of the loan that is processed currently.
type Loan struct {
	Principal float64
	Rate      float32
	StartDate time.Time
	Payment   []Installment
}

// Installment store load payment details.
type Installment struct {
	Amount float64
	Date   time.Time
}

var loan *Loan

// GetLoanInstance ...
func GetLoanInstance() *Loan {
	if loan == nil {
		return &Loan{}
	} else {
		loan.reset()
		return loan
	}
}

// Start starts a new loan.
func (l *Loan) Start(a float64, r float32, d string) error {
	t, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return err
	}
	l.Principal = a
	l.Rate = r
	l.StartDate = t
	return nil
}

func (l *Loan) reset() {
	l.Principal = 0
	l.Rate = 0
	l.StartDate = time.Time{}
	l.Payment = []Installment{}
}

// AddPayment ...
func (l *Loan) AddPayment(i Installment) error {
	duration := i.Date.Sub(l.StartDate)
	if duration < 0 {
		return fmt.Errorf("payment should be after loan start date")
	}
	l.Payment = append(l.Payment, i)
	return nil
}

// GetBalance - Returns balance loan amount.
func (l *Loan) GetBalance(d string) (float64, error) {
	t, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return 0, err
	}
	// calculate interest till date.
	interest := l.calculateInterest(t)
	// Add principal and interest to get current balance.
	return l.Principal + interest, nil
}

// AdjustPrincipal - Decrease principal is amount paid is greater than total interest till date.
func (l *Loan) AdjustPrincipal(i Installment) {
	// calculate interest till date.
	interest := l.calculateInterest(i.Date)
	// If i.Amount > interest; l.Principal = l.Principal - (i.Amount - interest)
	if interest < i.Amount {
		l.Principal = l.Principal - (i.Amount - interest)
	}
}

func (l *Loan) calculateInterest(d time.Time) float64 {
	dailyInterest := (float64(l.Rate) * l.Principal) / (100 * 365)
	daysSince := d.Sub(l.StartDate).Hours() / 24
	return float64(daysSince) * dailyInterest
}
