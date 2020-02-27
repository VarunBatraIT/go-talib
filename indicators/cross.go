package indicators

// Crossover returns true if series1 is crossing over series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for buy signals.
//          It assumes first values are the most recent.
//          The crossover function does not use most recent value, since usually it's not a complete candle.
//          The second recent values and the previous are used, instead.
func Crossover(series1, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-2] <= series2[N-2] && series1[N-1] > series2[N-1]
}

// Crossunder returns true if series1 is crossing under series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for sell signals.
func Crossunder(series1, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-1] <= series2[N-1] && series1[N-2] > series2[N-2]
}
