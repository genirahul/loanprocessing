package loan

import (
	"errors"
	"fmt"
	"loanprocessing/custom"
)

// ValidateInstallment : validate installment amount and date
func ValidateInstallment(l Loan, i Installment) error {
	bal, err := l.GetBalance(i.Date)
	if err != nil {
		return err
	}
	if i.Amount > bal {
		return errors.New("Invalid Payment : Amount cannout be greater than remaining balance")
	}
	if i.Date.Sub(l.StartDate).Hours() < 0 {
		return errors.New("Invalid Payment : Date cannot be before loan start date (" + l.StartDate.String() + ")")
	}
	d := l.getLastPaymentDate()
	if i.Date.Sub(d).Hours() < 0 {
		return errors.New("Invalid Payment : Date cannot be before last payment date (" + d.String() + ")")
	}
	return nil
}

// ValidateGetBalanceDate : validate get balance date
func ValidateGetBalanceDate(l Loan, d custom.Date) error {
	fmt.Println("start date :", l.StartDate.String(), "Date :", d.String())
	if d.Sub(l.StartDate).Hours() < 0 {
		fmt.Println("ERROR")
		return errors.New("Invalid Date : Date cannot be before loan start date " + l.StartDate.String())
	}
	return nil
}
