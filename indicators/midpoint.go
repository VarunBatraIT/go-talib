package indicators

// MidPoint - MidPoint over period
func MidPoint(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))
	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := inTimePeriod - 1
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded

	for today < len(inReal) {
		lowest := inReal[trailingIdx]
		trailingIdx++
		highest := lowest
		for i := trailingIdx; i <= today; i++ {
			tmp := inReal[i]
			if tmp < lowest {
				lowest = tmp
			} else if tmp > highest {
				highest = tmp
			}
		}
		outReal[outIdx] = (highest + lowest) / 2.0
		outIdx++
		today++
	}
	return outReal
}

// MidPrice - Midpoint Price over period
func MidPrice(inHigh, inLow []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inHigh))

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := inTimePeriod - 1
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	for today < len(inHigh) {
		lowest := inLow[trailingIdx]
		highest := inHigh[trailingIdx]
		trailingIdx++
		for i := trailingIdx; i <= today; i++ {
			tmp := inLow[i]
			if tmp < lowest {
				lowest = tmp
			}
			tmp = inHigh[i]
			if tmp > highest {
				highest = tmp
			}
		}
		outReal[outIdx] = (highest + lowest) / 2.0
		outIdx++
		today++
	}
	return outReal
}
