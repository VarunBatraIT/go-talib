package indicators

// Rocp - Rate of change Percentage: (price-prevPrice)/prevPrice
func Rocp(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 1 {
		return outReal
	}

	startIdx := inTimePeriod
	outIdx := startIdx
	inIdx := startIdx
	trailingIdx := startIdx - inTimePeriod
	for inIdx < len(outReal) {
		tempReal := inReal[trailingIdx]
		if tempReal != 0.0 {
			outReal[outIdx] = (inReal[inIdx] - tempReal) / tempReal
		} else {
			outReal[outIdx] = 0.0
		}
		trailingIdx++
		outIdx++
		inIdx++
	}

	return outReal
}

// Roc - Rate of change : ((price/prevPrice)-1)*100
func Roc(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	startIdx := inTimePeriod
	outIdx := inTimePeriod
	inIdx := startIdx
	trailingIdx := startIdx - inTimePeriod

	for inIdx < len(inReal) {
		tempReal := inReal[trailingIdx]
		if tempReal != 0.0 {
			outReal[outIdx] = ((inReal[inIdx] / tempReal) - 1.0) * 100.0
		} else {
			outReal[outIdx] = 0.0
		}
		trailingIdx++
		outIdx++
		inIdx++
	}
	return outReal
}

// Rocr - Rate of change ratio: (price/prevPrice)
func Rocr(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	startIdx := inTimePeriod
	outIdx := inTimePeriod
	inIdx := startIdx
	trailingIdx := startIdx - inTimePeriod

	for inIdx < len(inReal) {
		tempReal := inReal[trailingIdx]
		if tempReal != 0.0 {
			outReal[outIdx] = (inReal[inIdx] / tempReal)
		} else {
			outReal[outIdx] = 0.0
		}
		trailingIdx++
		outIdx++
		inIdx++
	}
	return outReal
}

// Rocr100 - Rate of change ratio 100 scale: (price/prevPrice)*100
func Rocr100(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	startIdx := inTimePeriod
	outIdx := inTimePeriod
	inIdx := startIdx
	trailingIdx := startIdx - inTimePeriod

	for inIdx < len(inReal) {
		tempReal := inReal[trailingIdx]
		if tempReal != 0.0 {
			outReal[outIdx] = (inReal[inIdx] / tempReal) * 100.0
		} else {
			outReal[outIdx] = 0.0
		}
		trailingIdx++
		outIdx++
		inIdx++
	}
	return outReal
}
