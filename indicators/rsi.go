package indicators

// Rsi - Relative strength index
func Rsi(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	// variable declarations
	tempValue1 := 0.0
	tempValue2 := 0.0
	outIdx := inTimePeriod
	today := 0
	prevValue := inReal[today]
	prevGain := 0.0
	prevLoss := 0.0
	today++

	for i := inTimePeriod; i > 0; i-- {
		tempValue1 = inReal[today]
		today++
		tempValue2 = tempValue1 - prevValue
		prevValue = tempValue1
		if tempValue2 < 0 {
			prevLoss -= tempValue2
		} else {
			prevGain += tempValue2
		}
	}

	prevLoss /= float64(inTimePeriod)
	prevGain /= float64(inTimePeriod)

	if today > 0 {
		tempValue1 = prevGain + prevLoss
		if !((-0.00000000000001 < tempValue1) && (tempValue1 < 0.00000000000001)) {
			outReal[outIdx] = 100.0 * (prevGain / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	} else {
		for today < 0 {
			tempValue1 = inReal[today]
			tempValue2 = tempValue1 - prevValue
			prevValue = tempValue1
			prevLoss *= float64(inTimePeriod - 1)
			prevGain *= float64(inTimePeriod - 1)
			if tempValue2 < 0 {
				prevLoss -= tempValue2
			} else {
				prevGain += tempValue2
			}
			prevLoss /= float64(inTimePeriod)
			prevGain /= float64(inTimePeriod)
			today++
		}
	}

	for today < len(inReal) {
		tempValue1 = inReal[today]
		today++
		tempValue2 = tempValue1 - prevValue
		prevValue = tempValue1
		prevLoss *= float64(inTimePeriod - 1)
		prevGain *= float64(inTimePeriod - 1)
		if tempValue2 < 0 {
			prevLoss -= tempValue2
		} else {
			prevGain += tempValue2
		}
		prevLoss /= float64(inTimePeriod)
		prevGain /= float64(inTimePeriod)
		tempValue1 = prevGain + prevLoss
		if !((-0.00000000000001 < tempValue1) && (tempValue1 < 0.00000000000001)) {
			outReal[outIdx] = 100.0 * (prevGain / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}

	return outReal
}

// StochRsi - Stochastic Relative Strength Index
func StochRsi(inReal []float64, inTimePeriod, inFastKPeriod, inFastDPeriod int, inFastDMAType MaType) ([]float64, []float64) {
	outFastK := make([]float64, len(inReal))
	outFastD := make([]float64, len(inReal))

	lookbackSTOCHF := (inFastKPeriod - 1) + (inFastDPeriod - 1)
	lookbackTotal := inTimePeriod + lookbackSTOCHF
	startIdx := lookbackTotal
	tempRSIBuffer := Rsi(inReal, inTimePeriod)
	tempk, tempd := StochF(tempRSIBuffer, tempRSIBuffer, tempRSIBuffer, inFastKPeriod, inFastDPeriod, inFastDMAType)

	for i := startIdx; i < len(inReal); i++ {
		outFastK[i] = tempk[i]
		outFastD[i] = tempd[i]
	}

	return outFastK, outFastD
}
