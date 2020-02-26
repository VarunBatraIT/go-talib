package indicators

// Apo - Absolute Price Oscillator
func Apo(inReal []float64, inFastPeriod int, inSlowPeriod int, inMAType MaType) []float64 {
	if inSlowPeriod < inFastPeriod {
		inSlowPeriod, inFastPeriod = inFastPeriod, inSlowPeriod
	}
	tempBuffer := Ma(inReal, inFastPeriod, inMAType)
	outReal := Ma(inReal, inSlowPeriod, inMAType)
	for i := inSlowPeriod - 1; i < len(inReal); i++ {
		outReal[i] = tempBuffer[i] - outReal[i]
	}

	return outReal
}

// Cmo - Chande Momentum Oscillator
func Cmo(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx
	if inTimePeriod == 1 {
		copy(outReal, inReal)
		return outReal
	}
	today := startIdx - lookbackTotal
	prevValue := inReal[today]
	prevGain := 0.0
	prevLoss := 0.0
	today++
	for i := inTimePeriod; i > 0; i-- {
		tempValue1 := inReal[today]
		tempValue2 := tempValue1 - prevValue
		prevValue = tempValue1
		if tempValue2 < 0 {
			prevLoss -= tempValue2
		} else {
			prevGain += tempValue2
		}
		today++
	}
	prevLoss /= float64(inTimePeriod)
	prevGain /= float64(inTimePeriod)
	if today > startIdx {
		tempValue1 := prevGain + prevLoss
		if !(((-(0.00000000000001)) < tempValue1) && (tempValue1 < (0.00000000000001))) {
			outReal[outIdx] = 100.0 * ((prevGain - prevLoss) / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	} else {
		for today < startIdx {
			tempValue1 := inReal[today]
			tempValue2 := tempValue1 - prevValue
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
		tempValue1 := inReal[today]
		today++
		tempValue2 := tempValue1 - prevValue
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
		if !(((-(0.00000000000001)) < tempValue1) && (tempValue1 < (0.00000000000001))) {
			outReal[outIdx] = 100.0 * ((prevGain - prevLoss) / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}
	return outReal
}

// Ppo - Percentage Price Oscillator
func Ppo(inReal []float64, inFastPeriod int, inSlowPeriod int, inMAType MaType) []float64 {
	if inSlowPeriod < inFastPeriod {
		inSlowPeriod, inFastPeriod = inFastPeriod, inSlowPeriod
	}
	tempBuffer := Ma(inReal, inFastPeriod, inMAType)
	outReal := Ma(inReal, inSlowPeriod, inMAType)

	for i := inSlowPeriod - 1; i < len(inReal); i++ {
		tempReal := outReal[i]
		if !(((-(0.00000000000001)) < tempReal) && (tempReal < (0.00000000000001))) {
			outReal[i] = ((tempBuffer[i] - tempReal) / tempReal) * 100.0
		} else {
			outReal[i] = 0.0
		}
	}

	return outReal
}

// AdOsc - Chaikin A/D Oscillator
func AdOsc(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64, inFastPeriod int, inSlowPeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	if (inFastPeriod < 2) || (inSlowPeriod < 2) {
		return outReal
	}

	slowestPeriod := 0
	if inFastPeriod < inSlowPeriod {
		slowestPeriod = inSlowPeriod
	} else {
		slowestPeriod = inFastPeriod
	}
	lookbackTotal := slowestPeriod - 1
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	ad := 0.0
	fastk := 2.0 / (float64(inFastPeriod) + 1.0)
	oneMinusfastk := 1.0 - fastk
	slowk := 2.0 / (float64(inSlowPeriod) + 1.0)
	oneMinusslowk := 1.0 - slowk
	high := inHigh[today]
	low := inLow[today]
	tmp := high - low
	close := inClose[today]
	if tmp > 0.0 {
		ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
	}
	today++
	fastEMA := ad
	slowEMA := ad

	for today < startIdx {
		high = inHigh[today]
		low = inLow[today]
		tmp = high - low
		close = inClose[today]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
		}
		today++

		fastEMA = (fastk * ad) + (oneMinusfastk * fastEMA)
		slowEMA = (slowk * ad) + (oneMinusslowk * slowEMA)
	}
	outIdx := lookbackTotal
	for today < len(inClose) {
		high = inHigh[today]
		low = inLow[today]
		tmp = high - low
		close = inClose[today]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[today])
		}
		today++
		fastEMA = (fastk * ad) + (oneMinusfastk * fastEMA)
		slowEMA = (slowk * ad) + (oneMinusslowk * slowEMA)
		outReal[outIdx] = fastEMA - slowEMA
		outIdx++
	}

	return outReal
}