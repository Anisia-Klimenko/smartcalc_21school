package credit

import (
	"errors"
	"log"
	"math"
	"strconv"
)

// Calculate calculates monthly payment, loan overpayment and total payment
// for given total loan amount, term, interest rate and type of payment
// (true for annuity, false for differentiated)
func Calculate(sum, term, rate string, isAnnuity bool) (monthly, overpay, total string, err error) {
	log.Println("credit: parameters: sum =", sum, ", term =", term, ", rate =", rate, ", is type annuity =", isAnnuity)

	// Check if passed parameters are valid
	sumFloat, termFloat, rateFloat, err := validParameters(sum, term, rate)
	if err != nil {
		return
	}

	var monthlyFloat, overpayFloat, totalFloat float64
	var monthlyFloatSlice []float64
	if isAnnuity {
		// Calculate result for annuity type of payment
		monthlyFloat, overpayFloat, totalFloat = annuityPayment(sumFloat, termFloat, rateFloat)
		// Convert monthly payment value to string
		monthly = strconv.FormatFloat(monthlyFloat, 'f', -1, 64)
	} else {
		// Calculate result for differentiated type of payment
		monthlyFloatSlice, overpayFloat, totalFloat = differentiatedPayment(sumFloat, termFloat, rateFloat)
		// Convert monthly payment slice to string in format "first_elem..second_elem"
		monthly = strconv.FormatFloat(monthlyFloatSlice[0], 'f', -1, 64) + ".." +
			strconv.FormatFloat(monthlyFloatSlice[len(monthlyFloatSlice)-1], 'f', -1, 64)
	}

	// Convert loan overpayment and total payment to string
	overpay = strconv.FormatFloat(overpayFloat, 'f', -1, 64)
	total = strconv.FormatFloat(totalFloat, 'f', -1, 64)

	log.Println("credit: calculated: monthly =", monthly, ", overpay =", overpay, ", total =", total)

	return
}

// annuityPayment calculates monthly payment, loan overpayment and total payment
// for given total loan amount, term and interest rate in case of annuity type of payment
func annuityPayment(sum, term, rateAnnual float64) (monthly, overpay, total float64) {
	rateMonthly := rateAnnual / 12 / 100

	// Round each value to 2 decimal places
	monthly = roundFloat(sum*rateMonthly*math.Pow(1+rateMonthly, term)/(math.Pow(1+rateMonthly, term)-1), 2)
	total = roundFloat(monthly*term, 2)
	overpay = roundFloat(total-sum, 2)

	return
}

// differentiatedPayment calculates monthly payment, loan overpayment and total payment
// for given total loan amount, term and interest rate in case of differentiated type of payment
func differentiatedPayment(sum, term, rateAnnual float64) (monthly []float64, overpay, total float64) {
	rateMonthly := rateAnnual / 12 / 100
	var left float64
	for i := 0; i < int(term); i++ {
		left = roundFloat(sum-sum*float64(i)/term, 2)
		if left <= 0 {
			break
		}
		monthly = append(monthly, roundFloat(sum/term+left*rateMonthly, 2))
		total += monthly[i]
	}

	// Round each value to 2 decimal places
	total = roundFloat(total, 2)
	overpay = roundFloat(total-sum, 2)

	return
}

// validParameters converts total loan amount, term and interest rate to float64 and returns them
// and error in case of error
func validParameters(sum, term, rate string) (sumFloat, termFloat, rateFloat float64, err error) {
	sumFloat, err = strconv.ParseFloat(sum, 64)
	if err == nil {
		termFloat, err = strconv.ParseFloat(term, 64)
		if err == nil {
			rateFloat, err = strconv.ParseFloat(rate, 64)
		}
	}
	if err != nil {
		return
	}

	// Check borders and term value
	if sumFloat <= 0 || termFloat != float64(int(termFloat)) || (rateFloat <= 0 || rateFloat > 100) {
		err = errors.New("non-valid value")
		return
	}

	return
}

// roundFloat rounds val to precision decimal places
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
