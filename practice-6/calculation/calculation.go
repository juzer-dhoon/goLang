package calculation

func CalculateInterest(principal, rate, time float64) float64 {
	return (principal * rate * time) / 100
}
