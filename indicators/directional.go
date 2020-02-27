package indicators

import "math"

// MinusDI - Minus Directional Indicator
func MinusDI(inHigh, inLow, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	lookbackTotal := 1
	if inTimePeriod > 1 {
		lookbackTotal = inTimePeriod
	}
	startIdx := lookbackTotal
	outIdx := startIdx

	prevHigh := 0.0
	prevLow := 0.0
	prevClose := 0.0
	if inTimePeriod <= 1 {
		today := startIdx - 1
		prevHigh = inHigh[today]
		prevLow = inLow[today]
		prevClose = inClose[today]
		for today < len(inClose)-1 {
			today++
			tempReal := inHigh[today]
			diffP := tempReal - prevHigh
			prevHigh = tempReal
			tempReal = inLow[today]
			diffM := prevLow - tempReal
			prevLow = tempReal
			if (diffM > 0) && (diffP < diffM) {
				tempReal = prevHigh - prevLow
				tempReal2 := math.Abs(prevHigh - prevClose)
				if tempReal2 > tempReal {
					tempReal = tempReal2
				}
				tempReal2 = math.Abs(prevLow - prevClose)
				if tempReal2 > tempReal {
					tempReal = tempReal2
				}

				if ((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001)) {
					outReal[outIdx] = 0.0
				} else {
					outReal[outIdx] = diffM / tempReal
				}
				outIdx++
			} else {
				outReal[outIdx] = 0.0
				outIdx++
			}
			prevClose = inClose[today]
		}
		return outReal
	}
	prevMinusDM := 0.0
	prevTR := 0.0
	today := startIdx - lookbackTotal
	prevHigh = inHigh[today]
	prevLow = inLow[today]
	prevClose = inClose[today]
	i := inTimePeriod - 1

	for i > 0 {
		i--
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM += diffM
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR += tempReal
		prevClose = inClose[today]
	}
	i = 1
	for i != 0 {
		i--
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM = prevMinusDM - (prevMinusDM / float64(inTimePeriod)) + diffM
		} else {
			prevMinusDM = prevMinusDM - (prevMinusDM / float64(inTimePeriod))
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR = prevTR - (prevTR / float64(inTimePeriod)) + tempReal
		prevClose = inClose[today]
	}
	if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
		outReal[startIdx] = (100.0 * (prevMinusDM / prevTR))
	} else {
		outReal[startIdx] = 0.0
	}
	outIdx = startIdx + 1
	for today < len(inClose)-1 {
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM = prevMinusDM - (prevMinusDM / float64(inTimePeriod)) + diffM
		} else {
			prevMinusDM = prevMinusDM - (prevMinusDM / float64(inTimePeriod))
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR = prevTR - (prevTR / float64(inTimePeriod)) + tempReal
		prevClose = inClose[today]
		if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
			outReal[outIdx] = (100.0 * (prevMinusDM / prevTR))
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}

	return outReal
}

// PlusDI - Plus Directional Indicator
func PlusDI(inHigh, inLow, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	lookbackTotal := 1
	if inTimePeriod > 1 {
		lookbackTotal = inTimePeriod
	}
	startIdx := lookbackTotal
	outIdx := startIdx

	prevHigh := 0.0
	prevLow := 0.0
	prevClose := 0.0
	if inTimePeriod <= 1 {
		today := startIdx - 1
		prevHigh = inHigh[today]
		prevLow = inLow[today]
		prevClose = inClose[today]
		for today < len(inClose)-1 {
			today++
			tempReal := inHigh[today]
			diffP := tempReal - prevHigh
			prevHigh = tempReal
			tempReal = inLow[today]
			diffM := prevLow - tempReal
			prevLow = tempReal
			if (diffP > 0) && (diffP > diffM) {
				tempReal = prevHigh - prevLow
				tempReal2 := math.Abs(prevHigh - prevClose)
				if tempReal2 > tempReal {
					tempReal = tempReal2
				}
				tempReal2 = math.Abs(prevLow - prevClose)
				if tempReal2 > tempReal {
					tempReal = tempReal2
				}

				if ((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001)) {
					outReal[outIdx] = 0.0
				} else {
					outReal[outIdx] = diffP / tempReal
				}
				outIdx++
			} else {
				outReal[outIdx] = 0.0
				outIdx++
			}
			prevClose = inClose[today]
		}
		return outReal
	}
	prevPlusDM := 0.0
	prevTR := 0.0
	today := startIdx - lookbackTotal
	prevHigh = inHigh[today]
	prevLow = inLow[today]
	prevClose = inClose[today]
	i := inTimePeriod - 1

	for i > 0 {
		i--
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffP > 0) && (diffP > diffM) {
			prevPlusDM += diffP
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR += tempReal
		prevClose = inClose[today]
	}
	i = 1
	for i != 0 {
		i--
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffP > 0) && (diffP > diffM) {
			prevPlusDM = prevPlusDM - (prevPlusDM / float64(inTimePeriod)) + diffP
		} else {
			prevPlusDM = prevPlusDM - (prevPlusDM / float64(inTimePeriod))
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR = prevTR - (prevTR / float64(inTimePeriod)) + tempReal
		prevClose = inClose[today]
	}
	if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
		outReal[startIdx] = (100.0 * (prevPlusDM / prevTR))
	} else {
		outReal[startIdx] = 0.0
	}
	outIdx = startIdx + 1
	for today < len(inClose)-1 {
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffP > 0) && (diffP > diffM) {
			prevPlusDM = prevPlusDM - (prevPlusDM / float64(inTimePeriod)) + diffP
		} else {
			prevPlusDM = prevPlusDM - (prevPlusDM / float64(inTimePeriod))
		}
		tempReal = prevHigh - prevLow
		tempReal2 := math.Abs(prevHigh - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}
		tempReal2 = math.Abs(prevLow - prevClose)
		if tempReal2 > tempReal {
			tempReal = tempReal2
		}

		prevTR = prevTR - (prevTR / float64(inTimePeriod)) + tempReal
		prevClose = inClose[today]
		if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
			outReal[outIdx] = (100.0 * (prevPlusDM / prevTR))
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}

	return outReal
}
