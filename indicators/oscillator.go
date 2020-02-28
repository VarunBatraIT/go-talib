package indicators

import "math"

// Apo - Absolute Price Oscillator
func Apo(inReal []float64, inFastPeriod, inSlowPeriod int, inMAType MaType) []float64 {
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
func Ppo(inReal []float64, inFastPeriod, inSlowPeriod int, inMAType MaType) []float64 {
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
func AdOsc(inHigh, inLow, inClose, inVolume []float64, inFastPeriod, inSlowPeriod int) []float64 {
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
	closeToday := inClose[today]
	if tmp > 0.0 {
		ad += (((closeToday - low) - (high - closeToday)) / tmp) * (inVolume[today])
	}
	today++
	fastEMA := ad
	slowEMA := ad

	for today < startIdx {
		high = inHigh[today]
		low = inLow[today]
		tmp = high - low
		closeToday = inClose[today]
		if tmp > 0.0 {
			ad += (((closeToday - low) - (high - closeToday)) / tmp) * (inVolume[today])
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
		closeToday = inClose[today]
		if tmp > 0.0 {
			ad += (((closeToday - low) - (high - closeToday)) / tmp) * (inVolume[today])
		}
		today++
		fastEMA = (fastk * ad) + (oneMinusfastk * fastEMA)
		slowEMA = (slowk * ad) + (oneMinusslowk * slowEMA)
		outReal[outIdx] = fastEMA - slowEMA
		outIdx++
	}

	return outReal
}

// UltOsc - Ultimate Oscillator
func UltOsc(inHigh, inLow, inClose []float64, inTimePeriod1, inTimePeriod2, inTimePeriod3 int) []float64 {
	outReal := make([]float64, len(inClose))

	usedFlag := make([]int, 3)
	periods := make([]int, 3)
	sortedPeriods := make([]int, 3)

	periods[0] = inTimePeriod1
	periods[1] = inTimePeriod2
	periods[2] = inTimePeriod3

	for i := 0; i < 3; i++ {
		longestPeriod := 0
		longestIndex := 0
		for j := 0; j < 3; j++ {
			if (usedFlag[j] == 0) && (periods[j] > longestPeriod) {
				longestPeriod = periods[j]
				longestIndex = j
			}
		}
		usedFlag[longestIndex] = 1
		sortedPeriods[i] = longestPeriod
	}
	inTimePeriod1 = sortedPeriods[2]
	inTimePeriod2 = sortedPeriods[1]
	inTimePeriod3 = sortedPeriods[0]

	lookbackTotal := 0
	if inTimePeriod1 > inTimePeriod2 {
		lookbackTotal = inTimePeriod1
	}
	if inTimePeriod3 > lookbackTotal {
		lookbackTotal = inTimePeriod3
	}
	lookbackTotal++

	startIdx := lookbackTotal - 1

	a1Total := 0.0
	b1Total := 0.0
	for i := startIdx - inTimePeriod1 + 1; i < startIdx; i++ {
		tempLT := inLow[i]
		tempHT := inHigh[i]
		tempCY := inClose[i-1]
		trueLow := 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow := inClose[i] - trueLow
		trueRange := tempHT - tempLT
		tempDouble := math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a1Total += closeMinusTrueLow
		b1Total += trueRange
	}

	a2Total := 0.0
	b2Total := 0.0
	for i := startIdx - inTimePeriod2 + 1; i < startIdx; i++ {
		tempLT := inLow[i]
		tempHT := inHigh[i]
		tempCY := inClose[i-1]
		trueLow := 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow := inClose[i] - trueLow
		trueRange := tempHT - tempLT
		tempDouble := math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a2Total += closeMinusTrueLow
		b2Total += trueRange
	}

	a3Total := 0.0
	b3Total := 0.0
	for i := startIdx - inTimePeriod3 + 1; i < startIdx; i++ {
		tempLT := inLow[i]
		tempHT := inHigh[i]
		tempCY := inClose[i-1]
		trueLow := 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow := inClose[i] - trueLow
		trueRange := tempHT - tempLT
		tempDouble := math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a3Total += closeMinusTrueLow
		b3Total += trueRange
	}

	//today := startIdx
	//outIdx := startIdx
	//trailingIdx1 := today - inTimePeriod1 + 1
	//trailingIdx2 := today - inTimePeriod2 + 1
	//trailingIdx3 := today - inTimePeriod3 + 1

	today := startIdx
	outIdx := startIdx
	trailingIdx1 := today - inTimePeriod1 + 1
	trailingIdx2 := today - inTimePeriod2 + 1
	trailingIdx3 := today - inTimePeriod3 + 1

	for today < len(inClose) {
		tempLT := inLow[today]
		tempHT := inHigh[today]
		tempCY := inClose[today-1]
		trueLow := 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow := inClose[today] - trueLow
		trueRange := tempHT - tempLT
		tempDouble := math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a1Total += closeMinusTrueLow
		a2Total += closeMinusTrueLow
		a3Total += closeMinusTrueLow
		b1Total += trueRange
		b2Total += trueRange
		b3Total += trueRange
		output := 0.0
		if !(((-(0.00000000000001)) < b1Total) && (b1Total < (0.00000000000001))) {
			output += 4.0 * (a1Total / b1Total)
		}
		if !(((-(0.00000000000001)) < b2Total) && (b2Total < (0.00000000000001))) {
			output += 2.0 * (a2Total / b2Total)
		}
		if !(((-(0.00000000000001)) < b3Total) && (b3Total < (0.00000000000001))) {
			output += a3Total / b3Total
		}
		tempLT = inLow[trailingIdx1]
		tempHT = inHigh[trailingIdx1]
		tempCY = inClose[trailingIdx1-1]
		trueLow = 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow = inClose[trailingIdx1] - trueLow
		trueRange = tempHT - tempLT
		tempDouble = math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a1Total -= closeMinusTrueLow
		b1Total -= trueRange
		tempLT = inLow[trailingIdx2]
		tempHT = inHigh[trailingIdx2]
		tempCY = inClose[trailingIdx2-1]
		trueLow = 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow = inClose[trailingIdx2] - trueLow
		trueRange = tempHT - tempLT
		tempDouble = math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a2Total -= closeMinusTrueLow
		b2Total -= trueRange
		tempLT = inLow[trailingIdx3]
		tempHT = inHigh[trailingIdx3]
		tempCY = inClose[trailingIdx3-1]
		trueLow = 0.0
		if tempLT < tempCY {
			trueLow = tempLT
		} else {
			trueLow = tempCY
		}
		closeMinusTrueLow = inClose[trailingIdx3] - trueLow
		trueRange = tempHT - tempLT
		tempDouble = math.Abs(tempCY - tempHT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}
		tempDouble = math.Abs(tempCY - tempLT)
		if tempDouble > trueRange {
			trueRange = tempDouble
		}

		a3Total -= closeMinusTrueLow
		b3Total -= trueRange
		outReal[outIdx] = 100.0 * (output / 7.0)
		outIdx++
		today++
		trailingIdx1++
		trailingIdx2++
		trailingIdx3++
	}
	return outReal
}
