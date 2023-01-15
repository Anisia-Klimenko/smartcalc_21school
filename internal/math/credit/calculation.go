package credit

import (
	"errors"
	"log"
	"math"
	"strconv"
)

func Calculate(sum, term, rate string, isAnnuity bool) (monthly, overpay, total string, err error) {
	sumFloat, termFloat, rateFloat, err := validParameters(sum, term, rate)
	if err != nil {
		return
	}

	var monthlyFloat, overpayFloat, totalFloat float64
	var monthlyFloatSlice []float64
	if isAnnuity {
		monthlyFloat, overpayFloat, totalFloat = annuityPayment(sumFloat, termFloat, rateFloat)
		monthly = strconv.FormatFloat(monthlyFloat, 'f', -1, 64)
	} else {
		monthlyFloatSlice, overpayFloat, totalFloat = differentiatedPayment(sumFloat, termFloat, rateFloat)
		monthly = strconv.FormatFloat(monthlyFloatSlice[0], 'f', -1, 64) + ".." +
			strconv.FormatFloat(monthlyFloatSlice[len(monthlyFloatSlice)-1], 'f', -1, 64)
	}
	overpay = strconv.FormatFloat(overpayFloat, 'f', -1, 64)
	total = strconv.FormatFloat(totalFloat, 'f', -1, 64)
	log.Println(monthly, overpay, total)
	return
}

func annuityPayment(sum, term, rateAnnual float64) (monthly, overpay, total float64) {
	rateMonthly := rateAnnual / 12 / 100
	monthly = roundFloat(sum*rateMonthly*math.Pow(1+rateMonthly, term)/(math.Pow(1+rateMonthly, term)-1), 2)
	total = roundFloat(monthly*term, 2)
	overpay = roundFloat(total-sum, 2)
	return
}

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
		log.Println("left", roundFloat(sum-total, 2), "monthly", monthly[i])
	}
	total = roundFloat(total, 2)
	overpay = roundFloat(total-sum, 2)
	return
}

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

	if sumFloat <= 0 || termFloat != float64(int(termFloat)) || (rateFloat <= 0 || rateFloat > 100) {
		err = errors.New("non-valid value")
		return
	}

	return
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
