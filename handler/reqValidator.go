package handler

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func customErrMsg(errs error) string {
	m := ""
	if errs != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := errs.(validator.ValidationErrors); ok {
			for _, err := range errs.(validator.ValidationErrors) {
				msg := prepareErrorMsg(err)
				if len(msg) > 0 {
					m += msg + ", "
				}
			}
		} else {
			m += errs.Error()
		}
	}
	return m
}

func prepareErrorMsg(err validator.FieldError) string {
	switch err.Namespace() {
	case "Loan.Principal":
		switch err.Tag() {
		case "required":
			return "initialAmount is required"
		case "min":
			return "initialAmount should be greater than 0"
		}
	case "Loan.Rate":
		switch err.Tag() {
		case "required":
			return "annualRate is required"
		case "min":
			return "annualRate should be greater than 0"
		}
	case "Loan.StartDate":
		switch err.Tag() {
		case "required":
			return "startDate is required"
		}
	case "Installment.Amount":
		switch err.Tag() {
		case "required":
			return "amount is required and should be greater than 0"
		case "min":
			return "amount is should be greater than 0"
		}
	case "Installment.Date":
		switch err.Tag() {
		case "required":
			return "date is required as yyyy-mm-dd (2006-01-02)"
		}
	}
	return ""
}
