package indicators

import "math"

// Adx - Average Directional Movement Index
func Adx(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := (2 * inTimePeriod) - 1
	startIdx := lookbackTotal
	outIdx := inTimePeriod
	prevMinusDM := 0.0
	prevPlusDM := 0.0
	prevTR := 0.0
	today := startIdx - lookbackTotal
	prevHigh := inHigh[today]
	prevLow := inLow[today]
	prevClose := inClose[today]
	for i := inTimePeriod - 1; i > 0; i-- {
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM += diffM
		} else if (diffP > 0) && (diffP > diffM) {
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
	sumDX := 0.0
	for i := inTimePeriod; i > 0; i-- {
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		prevMinusDM -= prevMinusDM / inTimePeriodF
		prevPlusDM -= prevPlusDM / inTimePeriodF
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM += diffM
		} else if (diffP > 0) && (diffP > diffM) {
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

		prevTR = prevTR - (prevTR / inTimePeriodF) + tempReal
		prevClose = inClose[today]
		if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
			minusDI := 100.0 * (prevMinusDM / prevTR)
			plusDI := 100.0 * (prevPlusDM / prevTR)
			tempReal = minusDI + plusDI
			if !(((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001))) {
				sumDX += 100.0 * (math.Abs(minusDI-plusDI) / tempReal)
			}
		}
	}
	prevADX := sumDX / inTimePeriodF

	outReal[startIdx] = prevADX
	outIdx = startIdx + 1
	today++
	for today < len(inClose) {
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		prevMinusDM -= prevMinusDM / inTimePeriodF
		prevPlusDM -= prevPlusDM / inTimePeriodF
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM += diffM
		} else if (diffP > 0) && (diffP > diffM) {
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

		prevTR = prevTR - (prevTR / inTimePeriodF) + tempReal
		prevClose = inClose[today]
		if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
			minusDI := 100.0 * (prevMinusDM / prevTR)
			plusDI := 100.0 * (prevPlusDM / prevTR)
			tempReal = minusDI + plusDI
			if !(((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001))) {
				tempReal = 100.0 * (math.Abs(minusDI-plusDI) / tempReal)
				prevADX = ((prevADX * (inTimePeriodF - 1)) + tempReal) / inTimePeriodF
			}
		}
		outReal[outIdx] = prevADX
		outIdx++
		today++
	}
	return outReal
}

// AdxR - Average Directional Movement Index Rating
func AdxR(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))
	startIdx := (2 * inTimePeriod) - 1
	tmpadx := Adx(inHigh, inLow, inClose, inTimePeriod)
	i := startIdx
	j := startIdx + inTimePeriod - 1
	for outIdx := startIdx + inTimePeriod - 1; outIdx < len(inClose); outIdx, i, j = outIdx+1, i+1, j+1 {
		outReal[outIdx] = ((tmpadx[i] + tmpadx[j]) / 2.0)
	}
	return outReal
}

// Dx - Directional Movement Index
func Dx(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	lookbackTotal := 2
	if inTimePeriod > 1 {
		lookbackTotal = inTimePeriod
	}
	startIdx := lookbackTotal
	outIdx := startIdx
	prevMinusDM := 0.0
	prevPlusDM := 0.0
	prevTR := 0.0
	today := startIdx - lookbackTotal
	prevHigh := inHigh[today]
	prevLow := inLow[today]
	prevClose := inClose[today]
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
		} else if (diffP > 0) && (diffP > diffM) {
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

	if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
		minusDI := (100.0 * (prevMinusDM / prevTR))
		plusDI := (100.0 * (prevPlusDM / prevTR))
		tempReal := minusDI + plusDI
		if !(((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001))) {
			outReal[outIdx] = (100.0 * (math.Abs(minusDI-plusDI) / tempReal))
		} else {
			outReal[outIdx] = 0.0
		}
	} else {
		outReal[outIdx] = 0.0
	}

	outIdx = startIdx
	for today < len(inClose)-1 {
		today++
		tempReal := inHigh[today]
		diffP := tempReal - prevHigh
		prevHigh = tempReal
		tempReal = inLow[today]
		diffM := prevLow - tempReal
		prevLow = tempReal
		prevMinusDM -= prevMinusDM / float64(inTimePeriod)
		prevPlusDM -= prevPlusDM / float64(inTimePeriod)
		if (diffM > 0) && (diffP < diffM) {
			prevMinusDM += diffM
		} else if (diffP > 0) && (diffP > diffM) {
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

		prevTR = prevTR - (prevTR / float64(inTimePeriod)) + tempReal
		prevClose = inClose[today]
		if !(((-(0.00000000000001)) < prevTR) && (prevTR < (0.00000000000001))) {
			minusDI := (100.0 * (prevMinusDM / prevTR))
			plusDI := (100.0 * (prevPlusDM / prevTR))
			tempReal = minusDI + plusDI
			if !(((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001))) {
				outReal[outIdx] = (100.0 * (math.Abs(minusDI-plusDI) / tempReal))
			} else {
				outReal[outIdx] = outReal[outIdx-1]
			}
		} else {
			outReal[outIdx] = outReal[outIdx-1]
		}
		outIdx++
	}
	return outReal
}

// MinusDM - Minus Directional Movement
func MinusDM(inHigh []float64, inLow []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inHigh))

	lookbackTotal := 1
	if inTimePeriod > 1 {
		lookbackTotal = inTimePeriod - 1
	}
	startIdx := lookbackTotal
	outIdx := startIdx
	today := startIdx
	prevHigh := 0.0
	prevLow := 0.0
	if inTimePeriod <= 1 {
		today = startIdx - 1
		prevHigh = inHigh[today]
		prevLow = inLow[today]
		for today < len(inHigh)-1 {
			today++
			tempReal := inHigh[today]
			diffP := tempReal - prevHigh
			prevHigh = tempReal
			tempReal = inLow[today]
			diffM := prevLow - tempReal
			prevLow = tempReal
			if (diffM > 0) && (diffP < diffM) {
				outReal[outIdx] = diffM
			} else {
				outReal[outIdx] = 0
			}
			outIdx++
		}
		return outReal
	}
	prevMinusDM := 0.0
	today = startIdx - lookbackTotal
	prevHigh = inHigh[today]
	prevLow = inLow[today]
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
	}
	i = 0
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
	}
	outReal[startIdx] = prevMinusDM
	outIdx = startIdx + 1
	for today < len(inHigh)-1 {
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
		outReal[outIdx] = prevMinusDM
		outIdx++
	}
	return outReal
}

// PlusDM - Plus Directional Movement
func PlusDM(inHigh []float64, inLow []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inHigh))

	lookbackTotal := 1
	if inTimePeriod > 1 {
		lookbackTotal = inTimePeriod - 1
	}
	startIdx := lookbackTotal
	outIdx := startIdx
	today := startIdx
	prevHigh := 0.0
	prevLow := 0.0
	if inTimePeriod <= 1 {
		today = startIdx - 1
		prevHigh = inHigh[today]
		prevLow = inLow[today]
		for today < len(inHigh)-1 {
			today++
			tempReal := inHigh[today]
			diffP := tempReal - prevHigh
			prevHigh = tempReal
			tempReal = inLow[today]
			diffM := prevLow - tempReal
			prevLow = tempReal
			if (diffP > 0) && (diffP > diffM) {
				outReal[outIdx] = diffP
			} else {
				outReal[outIdx] = 0
			}
			outIdx++
		}
		return outReal
	}
	prevPlusDM := 0.0
	today = startIdx - lookbackTotal
	prevHigh = inHigh[today]
	prevLow = inLow[today]
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
	}
	i = 0
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
	}
	outReal[startIdx] = prevPlusDM
	outIdx = startIdx + 1
	for today < len(inHigh)-1 {
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
		outReal[outIdx] = prevPlusDM
		outIdx++
	}
	return outReal
}