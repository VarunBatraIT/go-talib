package indicators

import "math"

// Cci - Commodity Channel Index
func Cci(inHigh, inLow, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))

	circBufferIdx := 0
	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	circBuffer := make([]float64, inTimePeriod)
	maxIdxCircBuffer := (inTimePeriod - 1)
	i := startIdx - lookbackTotal
	if inTimePeriod > 1 {
		for i < startIdx {
			circBuffer[circBufferIdx] = (inHigh[i] + inLow[i] + inClose[i]) / 3
			i++
			circBufferIdx++
			if circBufferIdx > maxIdxCircBuffer {
				circBufferIdx = 0
			}
		}
	}
	outIdx := inTimePeriod - 1
	for i < len(inClose) {
		lastValue := (inHigh[i] + inLow[i] + inClose[i]) / 3
		circBuffer[circBufferIdx] = lastValue
		theAverage := 0.0
		for j := 0; j < inTimePeriod; j++ {
			theAverage += circBuffer[j]
		}

		theAverage /= float64(inTimePeriod)
		tempReal2 := 0.0
		for j := 0; j < inTimePeriod; j++ {
			tempReal2 += math.Abs(circBuffer[j] - theAverage)
		}
		tempReal := lastValue - theAverage
		if (tempReal != 0.0) && (tempReal2 != 0.0) {
			outReal[outIdx] = tempReal / (0.015 * (tempReal2 / float64(inTimePeriod)))
		} else {
			outReal[outIdx] = 0.0
		}
		{
			circBufferIdx++
			if circBufferIdx > maxIdxCircBuffer {
				circBufferIdx = 0
			}
		}
		outIdx++
		i++
	}

	return outReal
}

// Mfi - Money Flow Index
func Mfi(inHigh, inLow, inClose, inVolume []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))
	mflowIdx := 0
	maxIdxMflow := (50 - 1)
	mflow := make([]moneyFlow, inTimePeriod)
	maxIdxMflow = inTimePeriod - 1
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx
	today := startIdx - lookbackTotal
	prevValue := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
	posSumMF := 0.0
	negSumMF := 0.0
	today++
	for i := inTimePeriod; i > 0; i-- {
		tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
		tempValue2 := tempValue1 - prevValue
		prevValue = tempValue1
		tempValue1 *= inVolume[today]
		today++
		if tempValue2 < 0 {
			(mflow[mflowIdx]).negative = tempValue1
			negSumMF += tempValue1
			(mflow[mflowIdx]).positive = 0.0
		} else if tempValue2 > 0 {
			(mflow[mflowIdx]).positive = tempValue1
			posSumMF += tempValue1
			(mflow[mflowIdx]).negative = 0.0
		} else {
			(mflow[mflowIdx]).positive = 0.0
			(mflow[mflowIdx]).negative = 0.0
		}
		mflowIdx++
		if mflowIdx > maxIdxMflow {
			mflowIdx = 0
		}
	}
	if today > startIdx {
		tempValue1 := posSumMF + negSumMF
		if tempValue1 < 1.0 {
		} else {
			outReal[outIdx] = 100.0 * (posSumMF / tempValue1)
			outIdx++
		}
	} else {
		for today < startIdx {
			posSumMF -= mflow[mflowIdx].positive
			negSumMF -= mflow[mflowIdx].negative
			tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
			tempValue2 := tempValue1 - prevValue
			prevValue = tempValue1
			tempValue1 *= inVolume[today]
			today++
			if tempValue2 < 0 {
				(mflow[mflowIdx]).negative = tempValue1
				negSumMF += tempValue1
				(mflow[mflowIdx]).positive = 0.0
			} else if tempValue2 > 0 {
				(mflow[mflowIdx]).positive = tempValue1
				posSumMF += tempValue1
				(mflow[mflowIdx]).negative = 0.0
			} else {
				(mflow[mflowIdx]).positive = 0.0
				(mflow[mflowIdx]).negative = 0.0
			}
			mflowIdx++
			if mflowIdx > maxIdxMflow {
				mflowIdx = 0
			}
		}
	}
	for today < len(inClose) {
		posSumMF -= (mflow[mflowIdx]).positive
		negSumMF -= (mflow[mflowIdx]).negative
		tempValue1 := (inHigh[today] + inLow[today] + inClose[today]) / 3.0
		tempValue2 := tempValue1 - prevValue
		prevValue = tempValue1
		tempValue1 *= inVolume[today]
		today++
		if tempValue2 < 0 {
			(mflow[mflowIdx]).negative = tempValue1
			negSumMF += tempValue1
			(mflow[mflowIdx]).positive = 0.0
		} else if tempValue2 > 0 {
			(mflow[mflowIdx]).positive = tempValue1
			posSumMF += tempValue1
			(mflow[mflowIdx]).negative = 0.0
		} else {
			(mflow[mflowIdx]).positive = 0.0
			(mflow[mflowIdx]).negative = 0.0
		}
		tempValue1 = posSumMF + negSumMF
		if tempValue1 < 1.0 {
			outReal[outIdx] = 0.0
		} else {
			outReal[outIdx] = 100.0 * (posSumMF / tempValue1)
		}
		outIdx++
		mflowIdx++
		if mflowIdx > maxIdxMflow {
			mflowIdx = 0
		}
	}
	return outReal
}
