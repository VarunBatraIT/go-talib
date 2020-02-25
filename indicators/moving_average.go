package indicators

import "math"

// Ma - Moving average
func Ma(inReal []float64, inTimePeriod int, inMAType MaType) []float64 {

	outReal := make([]float64, len(inReal))

	if inTimePeriod == 1 {
		copy(outReal, inReal)
		return outReal
	}

	switch inMAType {
	case SMA:
		outReal = Sma(inReal, inTimePeriod)
	case EMA:
		outReal = Ema(inReal, inTimePeriod)
	case WMA:
		outReal = Wma(inReal, inTimePeriod)
	case DEMA:
		outReal = Dema(inReal, inTimePeriod)
	case TEMA:
		outReal = Tema(inReal, inTimePeriod)
	case TRIMA:
		outReal = Trima(inReal, inTimePeriod)
	case KAMA:
		outReal = Kama(inReal, inTimePeriod)
	case MAMA:
		outReal, _ = Mama(inReal, 0.5, 0.05)
	case T3MA:
		outReal = T3(inReal, inTimePeriod, 0.7)
	}
	return outReal
}

// Dema - Double Exponential Moving Average
func Dema(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))
	firstEMA := Ema(inReal, inTimePeriod)
	secondEMA := Ema(firstEMA[inTimePeriod-1:], inTimePeriod)

	for outIdx, secondEMAIdx := (inTimePeriod*2)-2, inTimePeriod-1; outIdx < len(inReal); outIdx, secondEMAIdx = outIdx+1, secondEMAIdx+1 {
		outReal[outIdx] = (2.0 * firstEMA[outIdx]) - secondEMA[secondEMAIdx]
	}

	return outReal
}

// Ema - Exponential Moving Average
func ema(inReal []float64, inTimePeriod int, k1 float64) []float64 {
	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	i := inTimePeriod
	tempReal := 0.0
	for i > 0 {
		tempReal += inReal[today]
		today++
		i--
	}

	prevMA := tempReal / float64(inTimePeriod)

	for today <= startIdx {
		prevMA = ((inReal[today] - prevMA) * k1) + prevMA
		today++
	}
	outReal[startIdx] = prevMA
	outIdx := startIdx + 1
	for today < len(inReal) {
		prevMA = ((inReal[today] - prevMA) * k1) + prevMA
		outReal[outIdx] = prevMA
		today++
		outIdx++
	}

	return outReal
}

// Ema - Exponential Moving Average
func Ema(inReal []float64, inTimePeriod int) []float64 {
	k := 2.0 / float64(inTimePeriod+1)
	outReal := ema(inReal, inTimePeriod, k)
	return outReal
}

// Mama - MESA Adaptive Moving Average (lookback=32)
func Mama(inReal []float64, inFastLimit float64, inSlowLimit float64) ([]float64, []float64) {

	outMAMA := make([]float64, len(inReal))
	outFAMA := make([]float64, len(inReal))

	a := 0.0962
	b := 0.5769
	detrenderOdd := make([]float64, 3)
	detrenderEven := make([]float64, 3)
	q1Odd := make([]float64, 3)
	q1Even := make([]float64, 3)
	jIOdd := make([]float64, 3)
	jIEven := make([]float64, 3)
	jQOdd := make([]float64, 3)
	jQEven := make([]float64, 3)
	rad2Deg := 180.0 / (4.0 * math.Atan(1))
	lookbackTotal := 32
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal := inReal[today]
	today++
	periodWMASub := tempReal
	periodWMASum := tempReal
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 2.0
	tempReal = inReal[today]
	today++
	periodWMASub += tempReal
	periodWMASum += tempReal * 3.0
	trailingWMAValue := 0.0
	i := 9
	smoothedValue := 0.0
	for ok := true; ok; {
		tempReal = inReal[today]
		today++
		periodWMASub += tempReal
		periodWMASub -= trailingWMAValue
		periodWMASum += tempReal * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		i--
		ok = i != 0
	}
	hilbertIdx := 0
	detrenderOdd[0] = 0.0
	detrenderOdd[1] = 0.0
	detrenderOdd[2] = 0.0
	detrenderEven[0] = 0.0
	detrenderEven[1] = 0.0
	detrenderEven[2] = 0.0
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0

	q1Odd[0] = 0.0
	q1Odd[1] = 0.0
	q1Odd[2] = 0.0
	q1Even[0] = 0.0
	q1Even[1] = 0.0
	q1Even[2] = 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0

	jIOdd[0] = 0.0
	jIOdd[1] = 0.0
	jIOdd[2] = 0.0
	jIEven[0] = 0.0
	jIEven[1] = 0.0
	jIEven[2] = 0.0
	jI := 0.0
	prevjIOdd := 0.0
	prevjIEven := 0.0
	prevjIInputOdd := 0.0
	prevjIInputEven := 0.0

	jQOdd[0] = 0.0
	jQOdd[1] = 0.0
	jQOdd[2] = 0.0
	jQEven[0] = 0.0
	jQEven[1] = 0.0
	jQEven[2] = 0.0
	jQ := 0.0
	prevjQOdd := 0.0
	prevjQEven := 0.0
	prevjQInputOdd := 0.0
	prevjQInputEven := 0.0

	period := 0.0
	outIdx := startIdx
	previ2, prevq2 := 0.0, 0.0
	Re, Im := 0.0, 0.0
	mama, fama := 0.0, 0.0
	i1ForOddPrev3, i1ForEvenPrev3 := 0.0, 0.0
	i1ForOddPrev2, i1ForEvenPrev2 := 0.0, 0.0
	prevPhase := 0.0
	adjustedPrevPeriod := 0.0
	todayValue := 0.0
	hilbertTempReal := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod = (0.075 * period) + 0.54
		todayValue = inReal[today]

		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		q2, i2 := 0.0, 0.0
		tempReal2 := 0.0
		if (today % 2) == 0 {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderEven[hilbertIdx]
			detrenderEven[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderEven
			prevDetrenderEven = b * prevDetrenderInputEven
			detrender += prevDetrenderEven
			prevDetrenderInputEven = smoothedValue
			detrender *= adjustedPrevPeriod

			hilbertTempReal = a * detrender
			q1 = -q1Even[hilbertIdx]
			q1Even[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Even
			prevq1Even = b * prevq1InputEven
			q1 += prevq1Even
			prevq1InputEven = detrender
			q1 *= adjustedPrevPeriod

			hilbertTempReal = a * i1ForEvenPrev3
			jI = -jIEven[hilbertIdx]
			jIEven[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevjIEven
			prevjIEven = b * prevjIInputEven
			jI += prevjIEven
			prevjIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod

			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevjQEven
			prevjQEven = b * prevjQInputEven
			jQ += prevjQEven
			prevjQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
			if i1ForEvenPrev3 != 0.0 {
				tempReal2 = (math.Atan(q1/i1ForEvenPrev3) * rad2Deg)
			} else {
				tempReal2 = 0.0
			}
		} else {

			hilbertTempReal = a * smoothedValue
			detrender = -detrenderOdd[hilbertIdx]
			detrenderOdd[hilbertIdx] = hilbertTempReal
			detrender += hilbertTempReal
			detrender -= prevDetrenderOdd
			prevDetrenderOdd = b * prevDetrenderInputOdd
			detrender += prevDetrenderOdd
			prevDetrenderInputOdd = smoothedValue
			detrender *= adjustedPrevPeriod

			hilbertTempReal = a * detrender
			q1 = -q1Odd[hilbertIdx]
			q1Odd[hilbertIdx] = hilbertTempReal
			q1 += hilbertTempReal
			q1 -= prevq1Odd
			prevq1Odd = b * prevq1InputOdd
			q1 += prevq1Odd
			prevq1InputOdd = detrender
			q1 *= adjustedPrevPeriod

			hilbertTempReal = a * i1ForOddPrev3
			jI = -jIOdd[hilbertIdx]
			jIOdd[hilbertIdx] = hilbertTempReal
			jI += hilbertTempReal
			jI -= prevjIOdd
			prevjIOdd = b * prevjIInputOdd
			jI += prevjIOdd
			prevjIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod

			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevjQOdd
			prevjQOdd = b * prevjQInputOdd
			jQ += prevjQOdd
			prevjQInputOdd = q1
			jQ *= adjustedPrevPeriod

			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
			if i1ForOddPrev3 != 0.0 {
				tempReal2 = (math.Atan(q1/i1ForOddPrev3) * rad2Deg)
			} else {
				tempReal2 = 0.0
			}
		}
		tempReal = prevPhase - tempReal2
		prevPhase = tempReal2
		if tempReal < 1.0 {
			tempReal = 1.0
		}
		if tempReal > 1.0 {
			tempReal = inFastLimit / tempReal
			if tempReal < inSlowLimit {
				tempReal = inSlowLimit
			}
		} else {
			tempReal = inFastLimit
		}
		mama = (tempReal * todayValue) + ((1 - tempReal) * mama)
		tempReal *= 0.5
		fama = (tempReal * mama) + ((1 - tempReal) * fama)
		if today >= startIdx {
			outMAMA[outIdx] = mama
			outFAMA[outIdx] = fama
			outIdx++
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 = 1.5 * tempReal
		if period > tempReal2 {
			period = tempReal2
		}
		tempReal2 = 0.67 * tempReal
		if period < tempReal2 {
			period = tempReal2
		}
		if period < 6 {
			period = 6
		} else if period > 50 {
			period = 50
		}
		period = (0.2 * period) + (0.8 * tempReal)
		today++
	}
	return outMAMA, outFAMA
}

// MaVp - Moving average with variable period
func MaVp(inReal []float64, inPeriods []float64, inMinPeriod int, inMaxPeriod int, inMAType MaType) []float64 {

	outReal := make([]float64, len(inReal))
	startIdx := inMaxPeriod - 1
	outputSize := len(inReal)

	localPeriodArray := make([]float64, outputSize)
	for i := startIdx; i < outputSize; i++ {
		tempInt := int(inPeriods[i])
		if tempInt < inMinPeriod {
			tempInt = inMinPeriod
		} else if tempInt > inMaxPeriod {
			tempInt = inMaxPeriod
		}
		localPeriodArray[i] = float64(tempInt)
	}

	for i := startIdx; i < outputSize; i++ {
		curPeriod := int(localPeriodArray[i])
		if curPeriod != 0 {
			localOutputArray := Ma(inReal, curPeriod, inMAType)
			outReal[i] = localOutputArray[i]
			for j := i + 1; j < outputSize; j++ {
				if localPeriodArray[j] == float64(curPeriod) {
					localPeriodArray[j] = 0
					outReal[j] = localOutputArray[j]
				}
			}
		}
	}
	return outReal
}

// Kama - Kaufman Adaptive Moving Average
func Kama(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	constMax := 2.0 / (30.0 + 1.0)
	constDiff := 2.0/(2.0+1.0) - constMax
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	sumROC1 := 0.0
	today := startIdx - lookbackTotal
	trailingIdx := today
	i := inTimePeriod
	for i > 0 {
		tempReal := inReal[today]
		today++
		tempReal -= inReal[today]
		sumROC1 += math.Abs(tempReal)
		i--
	}
	prevKAMA := inReal[today-1]
	tempReal := inReal[today]
	tempReal2 := inReal[trailingIdx]
	trailingIdx++
	periodROC := tempReal - tempReal2
	trailingValue := tempReal2
	if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
		tempReal = 1.0
	} else {
		tempReal = math.Abs(periodROC / sumROC1)
	}
	tempReal = (tempReal * constDiff) + constMax
	tempReal *= tempReal
	prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
	today++
	for today <= startIdx {
		tempReal = inReal[today]
		tempReal2 = inReal[trailingIdx]
		trailingIdx++
		periodROC = tempReal - tempReal2
		sumROC1 -= math.Abs(trailingValue - tempReal2)
		sumROC1 += math.Abs(tempReal - inReal[today-1])
		trailingValue = tempReal2
		if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
			tempReal = 1.0
		} else {
			tempReal = math.Abs(periodROC / sumROC1)
		}
		tempReal = (tempReal * constDiff) + constMax
		tempReal *= tempReal
		prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
		today++
	}
	outReal[inTimePeriod] = prevKAMA
	outIdx := inTimePeriod + 1
	for today < len(inReal) {
		tempReal = inReal[today]
		tempReal2 = inReal[trailingIdx]
		trailingIdx++
		periodROC = tempReal - tempReal2
		sumROC1 -= math.Abs(trailingValue - tempReal2)
		sumROC1 += math.Abs(tempReal - inReal[today-1])
		trailingValue = tempReal2
		if (sumROC1 <= periodROC) || (((-(0.00000000000001)) < sumROC1) && (sumROC1 < (0.00000000000001))) {
			tempReal = 1.0
		} else {
			tempReal = math.Abs(periodROC / sumROC1)
		}
		tempReal = (tempReal * constDiff) + constMax
		tempReal *= tempReal
		prevKAMA = ((inReal[today] - prevKAMA) * tempReal) + prevKAMA
		today++
		outReal[outIdx] = prevKAMA
		outIdx++
	}

	return outReal
}


// Macd - Moving Average Convergence/Divergence
// unstable period ~= 100
func Macd(inReal []float64, inFastPeriod int, inSlowPeriod int, inSignalPeriod int) ([]float64, []float64, []float64) {

	if inSlowPeriod < inFastPeriod {
		inSlowPeriod, inFastPeriod = inFastPeriod, inSlowPeriod
	}

	k1 := 0.0
	k2 := 0.0
	if inSlowPeriod != 0 {
		k1 = 2.0 / float64(inSlowPeriod+1)
	} else {
		inSlowPeriod = 26
		k1 = 0.075
	}
	if inFastPeriod != 0 {
		k2 = 2.0 / float64(inFastPeriod+1)
	} else {
		inFastPeriod = 12
		k2 = 0.15
	}

	lookbackSignal := inSignalPeriod - 1
	lookbackTotal := lookbackSignal
	lookbackTotal += (inSlowPeriod - 1)

	fastEMABuffer := ema(inReal, inFastPeriod, k2)
	slowEMABuffer := ema(inReal, inSlowPeriod, k1)
	for i := 0; i < len(fastEMABuffer); i++ {
		fastEMABuffer[i] = fastEMABuffer[i] - slowEMABuffer[i]
	}

	outMACD := make([]float64, len(inReal))
	for i := lookbackTotal - 1; i < len(fastEMABuffer); i++ {
		outMACD[i] = fastEMABuffer[i]
	}
	outMACDSignal := ema(outMACD, inSignalPeriod, (2.0 / float64(inSignalPeriod+1)))

	outMACDHist := make([]float64, len(inReal))
	for i := lookbackTotal; i < len(outMACDHist); i++ {
		outMACDHist[i] = outMACD[i] - outMACDSignal[i]
	}

	return outMACD, outMACDSignal, outMACDHist
}

// MacdExt - MACD with controllable MA type
// unstable period ~= 100
func MacdExt(inReal []float64, inFastPeriod int, inFastMAType MaType, inSlowPeriod int, inSlowMAType MaType, inSignalPeriod int, inSignalMAType MaType) ([]float64, []float64, []float64) {

	lookbackLargest := 0
	if inFastPeriod < inSlowPeriod {
		lookbackLargest = inSlowPeriod
	} else {
		lookbackLargest = inFastPeriod
	}
	lookbackTotal := (inSignalPeriod - 1) + (lookbackLargest - 1)

	outMACD := make([]float64, len(inReal))
	outMACDSignal := make([]float64, len(inReal))
	outMACDHist := make([]float64, len(inReal))

	slowMABuffer := Ma(inReal, inSlowPeriod, inSlowMAType)
	fastMABuffer := Ma(inReal, inFastPeriod, inFastMAType)
	tempBuffer1 := make([]float64, len(inReal))

	for i := 0; i < len(slowMABuffer); i++ {
		tempBuffer1[i] = fastMABuffer[i] - slowMABuffer[i]
	}
	tempBuffer2 := Ma(tempBuffer1, inSignalPeriod, inSignalMAType)

	for i := lookbackTotal; i < len(outMACDHist); i++ {
		outMACD[i] = tempBuffer1[i]
		outMACDSignal[i] = tempBuffer2[i]
		outMACDHist[i] = outMACD[i] - outMACDSignal[i]
	}

	return outMACD, outMACDSignal, outMACDHist
}

// MacdFix - MACD Fix 12/26
// unstable period ~= 100
func MacdFix(inReal []float64, inSignalPeriod int) ([]float64, []float64, []float64) {
	return Macd(inReal, 0, 0, inSignalPeriod)
}



// Sma - Simple Moving Average
func Sma(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	periodTotal := 0.0
	trailingIdx := startIdx - lookbackTotal
	i := trailingIdx
	if inTimePeriod > 1 {
		for i < startIdx {
			periodTotal += inReal[i]
			i++
		}
	}
	outIdx := startIdx
	for ok := true; ok; {
		periodTotal += inReal[i]
		tempReal := periodTotal
		periodTotal -= inReal[trailingIdx]
		outReal[outIdx] = tempReal / float64(inTimePeriod)
		trailingIdx++
		i++
		outIdx++
		ok = i < len(outReal)
	}

	return outReal
}

// T3 - Triple Exponential Moving Average (T3) (lookback=6*inTimePeriod)
func T3(inReal []float64, inTimePeriod int, inVFactor float64) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := 6 * (inTimePeriod - 1)
	startIdx := lookbackTotal
	today := startIdx - lookbackTotal
	k := 2.0 / (float64(inTimePeriod) + 1.0)
	oneMinusK := 1.0 - k
	tempReal := inReal[today]
	today++
	for i := inTimePeriod - 1; i > 0; i-- {
		tempReal += inReal[today]
		today++
	}
	e1 := tempReal / float64(inTimePeriod)
	tempReal = e1
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		tempReal += e1
		today++
	}
	e2 := tempReal / float64(inTimePeriod)
	tempReal = e2
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		tempReal += e2
		today++
	}
	e3 := tempReal / float64(inTimePeriod)
	tempReal = e3
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		tempReal += e3
		today++
	}
	e4 := tempReal / float64(inTimePeriod)
	tempReal = e4
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		tempReal += e4
		today++
	}
	e5 := tempReal / float64(inTimePeriod)
	tempReal = e5
	for i := inTimePeriod - 1; i > 0; i-- {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		tempReal += e5
		today++
	}
	e6 := tempReal / float64(inTimePeriod)
	for today <= startIdx {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		e6 = (k * e5) + (oneMinusK * e6)
		today++
	}
	tempReal = inVFactor * inVFactor
	c1 := -(tempReal * inVFactor)
	c2 := 3.0 * (tempReal - c1)
	c3 := -6.0*tempReal - 3.0*(inVFactor-c1)
	c4 := 1.0 + 3.0*inVFactor - c1 + 3.0*tempReal
	outIdx := lookbackTotal
	outReal[outIdx] = c1*e6 + c2*e5 + c3*e4 + c4*e3
	outIdx++
	for today < len(inReal) {
		e1 = (k * inReal[today]) + (oneMinusK * e1)
		e2 = (k * e1) + (oneMinusK * e2)
		e3 = (k * e2) + (oneMinusK * e3)
		e4 = (k * e3) + (oneMinusK * e4)
		e5 = (k * e4) + (oneMinusK * e5)
		e6 = (k * e5) + (oneMinusK * e6)
		outReal[outIdx] = c1*e6 + c2*e5 + c3*e4 + c4*e3
		outIdx++
		today++
	}

	return outReal
}

// Tema - Triple Exponential Moving Average
func Tema(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))
	firstEMA := Ema(inReal, inTimePeriod)
	secondEMA := Ema(firstEMA[inTimePeriod-1:], inTimePeriod)
	thirdEMA := Ema(secondEMA[inTimePeriod-1:], inTimePeriod)

	outIdx := (inTimePeriod * 3) - 3
	secondEMAIdx := (inTimePeriod * 2) - 2
	thirdEMAIdx := inTimePeriod - 1

	for outIdx < len(inReal) {
		outReal[outIdx] = thirdEMA[thirdEMAIdx] + ((3.0 * firstEMA[outIdx]) - (3.0 * secondEMA[secondEMAIdx]))
		outIdx++
		secondEMAIdx++
		thirdEMAIdx++
	}

	return outReal
}

// Trima - Triangular Moving Average
func Trima(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	outIdx := inTimePeriod - 1
	var factor float64

	if inTimePeriod%2 == 1 {
		i := inTimePeriod >> 1
		factor = (float64(i) + 1.0) * (float64(i) + 1.0)
		factor = 1.0 / factor
		trailingIdx := startIdx - lookbackTotal
		middleIdx := trailingIdx + i
		todayIdx := middleIdx + i
		numerator := 0.0
		numeratorSub := 0.0
		for i := middleIdx; i >= trailingIdx; i-- {
			tempReal := inReal[i]
			numeratorSub += tempReal
			numerator += numeratorSub
		}
		numeratorAdd := 0.0
		middleIdx++
		for i := middleIdx; i <= todayIdx; i++ {
			tempReal := inReal[i]
			numeratorAdd += tempReal
			numerator += numeratorAdd
		}
		outIdx = inTimePeriod - 1
		tempReal := inReal[trailingIdx]
		trailingIdx++
		outReal[outIdx] = numerator * factor
		outIdx++
		todayIdx++
		for todayIdx < len(inReal) {
			numerator -= numeratorSub
			numeratorSub -= tempReal
			tempReal = inReal[middleIdx]
			middleIdx++
			numeratorSub += tempReal
			numerator += numeratorAdd
			numeratorAdd -= tempReal
			tempReal = inReal[todayIdx]
			todayIdx++
			numeratorAdd += tempReal
			numerator += tempReal
			tempReal = inReal[trailingIdx]
			trailingIdx++
			outReal[outIdx] = numerator * factor
			outIdx++
		}

	} else {

		i := (inTimePeriod >> 1)
		factor = float64(i) * (float64(i) + 1)
		factor = 1.0 / factor
		trailingIdx := startIdx - lookbackTotal
		middleIdx := trailingIdx + i - 1
		todayIdx := middleIdx + i
		numerator := 0.0
		numeratorSub := 0.0
		for i := middleIdx; i >= trailingIdx; i-- {
			tempReal := inReal[i]
			numeratorSub += tempReal
			numerator += numeratorSub
		}
		numeratorAdd := 0.0
		middleIdx++
		for i := middleIdx; i <= todayIdx; i++ {
			tempReal := inReal[i]
			numeratorAdd += tempReal
			numerator += numeratorAdd
		}
		outIdx = inTimePeriod - 1
		tempReal := inReal[trailingIdx]
		trailingIdx++
		outReal[outIdx] = numerator * factor
		outIdx++
		todayIdx++

		for todayIdx < len(inReal) {
			numerator -= numeratorSub
			numeratorSub -= tempReal
			tempReal = inReal[middleIdx]
			middleIdx++
			numeratorSub += tempReal
			numeratorAdd -= tempReal
			numerator += numeratorAdd
			tempReal = inReal[todayIdx]
			todayIdx++
			numeratorAdd += tempReal
			numerator += tempReal
			tempReal = inReal[trailingIdx]
			trailingIdx++
			outReal[outIdx] = numerator * factor
			outIdx++
		}
	}
	return outReal
}

// Wma - Weighted Moving Average
func Wma(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal

	if inTimePeriod == 1 {
		copy(outReal, inReal)
		return outReal
	}
	divider := (inTimePeriod * (inTimePeriod + 1)) >> 1
	outIdx := inTimePeriod - 1
	trailingIdx := startIdx - lookbackTotal
	periodSum, periodSub := 0.0, 0.0
	inIdx := trailingIdx
	i := 1
	for inIdx < startIdx {
		tempReal := inReal[inIdx]
		periodSub += tempReal
		periodSum += tempReal * float64(i)
		inIdx++
		i++
	}
	trailingValue := 0.0
	for inIdx < len(inReal) {
		tempReal := inReal[inIdx]
		periodSub += tempReal
		periodSub -= trailingValue
		periodSum += tempReal * float64(inTimePeriod)
		trailingValue = inReal[trailingIdx]
		outReal[outIdx] = periodSum / float64(divider)
		periodSum -= periodSub
		inIdx++
		trailingIdx++
		outIdx++
	}
	return outReal
}