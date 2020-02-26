/*
Copyright 2016 Mark Chenoweth
Copyright 2020 The GoCryptoTrader Developers
Licensed under terms of MIT license (see LICENSE)
*/

// Package talib is a pure Go port of TA-Lib (http://ta-lib.org) Technical Analysis Library
package indicators

import (
	"math"
)

/* Momentum Indicators */

// Bop - Balance Of Power
func Bop(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) []float64 {
	outReal := make([]float64, len(inClose))

	for i := 0; i < len(inClose); i++ {
		tempReal := inHigh[i] - inLow[i]
		if tempReal < (0.00000000000001) {
			outReal[i] = 0.0
		} else {
			outReal[i] = (inClose[i] - inOpen[i]) / tempReal
		}
	}

	return outReal
}

// Cci - Commodity Channel Index
func Cci(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
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

// MinusDI - Minus Directional Indicator
func MinusDI(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
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

// Mfi - Money Flow Index
func Mfi(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64, inTimePeriod int) []float64 {
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

// Mom - Momentum
func Mom(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inIdx, outIdx, trailingIdx := inTimePeriod, inTimePeriod, 0
	for inIdx < len(inReal) {
		outReal[outIdx] = inReal[inIdx] - inReal[trailingIdx]
		inIdx, outIdx, trailingIdx = inIdx+1, outIdx+1, trailingIdx+1
	}

	return outReal
}

// PlusDI - Plus Directional Indicator
func PlusDI(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
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

// Stoch - Stochastic
func Stoch(inHigh []float64, inLow []float64, inClose []float64, inFastKPeriod int, inSlowKPeriod int, inSlowKMAType MaType, inSlowDPeriod int, inSlowDMAType MaType) ([]float64, []float64) {
	outSlowK := make([]float64, len(inClose))
	outSlowD := make([]float64, len(inClose))

	lookbackK := inFastKPeriod - 1
	lookbackKSlow := inSlowKPeriod - 1
	lookbackDSlow := inSlowDPeriod - 1
	lookbackTotal := lookbackK + lookbackDSlow + lookbackKSlow
	startIdx := lookbackTotal
	outIdx := 0
	trailingIdx := startIdx - lookbackTotal
	today := trailingIdx + lookbackK
	lowestIdx, highestIdx := -1, -1
	diff, highest, lowest := 0.0, 0.0, 0.0
	tempBuffer := make([]float64, len(inClose)-today+1)
	for today < len(inClose) {
		tmp := inLow[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inLow[lowestIdx]
			i := lowestIdx + 1
			for i <= today {
				tmp := inLow[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
			diff = (highest - lowest) / 100.0
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
			diff = (highest - lowest) / 100.0
		}
		tmp = inHigh[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inHigh[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp := inHigh[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
			diff = (highest - lowest) / 100.0
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
			diff = (highest - lowest) / 100.0
		}
		if diff != 0.0 {
			tempBuffer[outIdx] = (inClose[today] - lowest) / diff
		} else {
			tempBuffer[outIdx] = 0.0
		}
		outIdx++
		trailingIdx++
		today++
	}

	tempBuffer1 := Ma(tempBuffer, inSlowKPeriod, inSlowKMAType)
	tempBuffer2 := Ma(tempBuffer1, inSlowDPeriod, inSlowDMAType)
	//for i, j := lookbackK, lookbackTotal; j < len(inClose); i, j = i+1, j+1 {
	for i, j := lookbackDSlow+lookbackKSlow, lookbackTotal; j < len(inClose); i, j = i+1, j+1 {
		outSlowK[j] = tempBuffer1[i]
		outSlowD[j] = tempBuffer2[i]
	}

	return outSlowK, outSlowD
}

// StochF - Stochastic Fast
func StochF(inHigh []float64, inLow []float64, inClose []float64, inFastKPeriod int, inFastDPeriod int, inFastDMAType MaType) ([]float64, []float64) {
	outFastK := make([]float64, len(inClose))
	outFastD := make([]float64, len(inClose))

	lookbackK := inFastKPeriod - 1
	lookbackFastD := inFastDPeriod - 1
	lookbackTotal := lookbackK + lookbackFastD
	startIdx := lookbackTotal
	outIdx := 0
	trailingIdx := startIdx - lookbackTotal
	today := trailingIdx + lookbackK
	lowestIdx, highestIdx := -1, -1
	diff, highest, lowest := 0.0, 0.0, 0.0
	tempBuffer := make([]float64, (len(inClose) - today + 1))

	for today < len(inClose) {
		tmp := inLow[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inLow[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmp = inLow[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
			diff = (highest - lowest) / 100.0
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
			diff = (highest - lowest) / 100.0
		}
		tmp = inHigh[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inHigh[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmp = inHigh[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
			diff = (highest - lowest) / 100.0
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
			diff = (highest - lowest) / 100.0
		}
		if diff != 0.0 {
			tempBuffer[outIdx] = (inClose[today] - lowest) / diff
		} else {
			tempBuffer[outIdx] = 0.0
		}
		outIdx++
		trailingIdx++
		today++
	}

	tempBuffer1 := Ma(tempBuffer, inFastDPeriod, inFastDMAType)
	for i, j := lookbackFastD, lookbackTotal; j < len(inClose); i, j = i+1, j+1 {
		outFastK[j] = tempBuffer[i]
		outFastD[j] = tempBuffer1[i]
	}

	return outFastK, outFastD
}

//Trix - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA
func Trix(inReal []float64, inTimePeriod int) []float64 {
	tmpReal := Ema(inReal, inTimePeriod)
	tmpReal = Ema(tmpReal[inTimePeriod-1:], inTimePeriod)
	tmpReal = Ema(tmpReal[inTimePeriod-1:], inTimePeriod)
	tmpReal = Roc(tmpReal, 1)

	outReal := make([]float64, len(inReal))
	for i, j := inTimePeriod, ((inTimePeriod-1)*3)+1; j < len(outReal); i, j = i+1, j+1 {
		outReal[j] = tmpReal[i]
	}

	return outReal
}

// UltOsc - Ultimate Oscillator
func UltOsc(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod1 int, inTimePeriod2 int, inTimePeriod3 int) []float64 {
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

// WillR - Williams' %R
func WillR(inHigh []float64, inLow []float64, inClose []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inClose))
	nbInitialElementNeeded := (inTimePeriod - 1)
	diff := 0.0
	outIdx := inTimePeriod - 1
	startIdx := inTimePeriod - 1
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	lowestIdx := -1
	highest := 0.0
	lowest := 0.0
	i := 0
	for today < len(inClose) {
		tmp := inLow[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inLow[lowestIdx]
			i = lowestIdx
			i++
			for i <= today {
				tmp = inLow[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
			diff = (highest - lowest) / (-100.0)
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
			diff = (highest - lowest) / (-100.0)
		}
		tmp = inHigh[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inHigh[highestIdx]
			i = highestIdx
			i++
			for i <= today {
				tmp = inHigh[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
			diff = (highest - lowest) / (-100.0)
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
			diff = (highest - lowest) / (-100.0)
		}
		if diff != 0.0 {
			outReal[outIdx] = (highest - inClose[today]) / diff
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
		trailingIdx++
		today++
	}
	return outReal
}

/* Volume Indicators */

// Ad - Chaikin A/D Line
func Ad(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64) []float64 {
	outReal := make([]float64, len(inClose))

	startIdx := 0
	nbBar := len(inClose) - startIdx
	currentBar := startIdx
	outIdx := 0
	ad := 0.0
	for nbBar != 0 {
		high := inHigh[currentBar]
		low := inLow[currentBar]
		tmp := high - low
		close := inClose[currentBar]
		if tmp > 0.0 {
			ad += (((close - low) - (high - close)) / tmp) * (inVolume[currentBar])
		}
		outReal[outIdx] = ad
		outIdx++
		currentBar++
		nbBar--
	}
	return outReal
}

// Obv - On Balance Volume
func Obv(inReal []float64, inVolume []float64) []float64 {
	outReal := make([]float64, len(inReal))
	startIdx := 0
	prevOBV := inVolume[startIdx]
	prevReal := inReal[startIdx]
	outIdx := 0
	for i := startIdx; i < len(inReal); i++ {
		tempReal := inReal[i]
		if tempReal > prevReal {
			prevOBV += inVolume[i]
		} else if tempReal < prevReal {
			prevOBV -= inVolume[i]
		}
		outReal[outIdx] = prevOBV
		prevReal = tempReal
		outIdx++
	}
	return outReal
}

/* Volatility Indicators */

/* Price Transform */

// AvgPrice - Average Price (o+h+l+c)/4
func AvgPrice(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) []float64 {
	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + inClose[i] + inOpen[i]) / 4
		outIdx++
	}
	return outReal
}

// MedPrice - Median Price (h+l)/2
func MedPrice(inHigh []float64, inLow []float64) []float64 {
	outReal := make([]float64, len(inHigh))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inHigh); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i]) / 2.0
		outIdx++
	}
	return outReal
}

// TypPrice - Typical Price (h+l+c)/3
func TypPrice(inHigh []float64, inLow []float64, inClose []float64) []float64 {
	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + inClose[i]) / 3.0
		outIdx++
	}
	return outReal
}

// WclPrice - Weighted Close Price
func WclPrice(inHigh []float64, inLow []float64, inClose []float64) []float64 {
	outReal := make([]float64, len(inClose))
	outIdx := 0
	startIdx := 0

	for i := startIdx; i < len(inClose); i++ {
		outReal[outIdx] = (inHigh[i] + inLow[i] + (inClose[i] * 2.0)) / 4.0
		outIdx++
	}
	return outReal
}

/* Cycle Indicators */

// HtDcPeriod - Hilbert Transform - Dominant Cycle Period (lookback=32)
func HtDcPeriod(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))

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
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 32
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i2 := 0.0
	q2 := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		hilbertTempReal := 0.0
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
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
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
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
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
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		if today >= startIdx {
			outReal[outIdx] = smoothPeriod
			outIdx++
		}
		today++
	}
	return outReal
}

// HtDcPhase - Hilbert Transform - Dominant Cycle Phase (lookback=63)
func HtDcPhase(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
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
	smoothPriceIdx := 0
	maxIdxSmoothPrice := (50 - 1)
	smoothPrice := make([]float64, maxIdxSmoothPrice+1)
	tempReal := math.Atan(1)
	rad2Deg := 45.0 / tempReal
	constDeg2RadBy360 := tempReal * 8.0
	lookbackTotal := 63
	startIdx := lookbackTotal
	trailingWMAIdx := startIdx - lookbackTotal
	today := trailingWMAIdx
	tempReal = inReal[today]
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
	i := 34
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
	detrender := 0.0
	prevDetrenderOdd := 0.0
	prevDetrenderEven := 0.0
	prevDetrenderInputOdd := 0.0
	prevDetrenderInputEven := 0.0
	q1 := 0.0
	prevq1Odd := 0.0
	prevq1Even := 0.0
	prevq1InputOdd := 0.0
	prevq1InputEven := 0.0
	jI := 0.0
	prevJIOdd := 0.0
	prevJIEven := 0.0
	prevJIInputOdd := 0.0
	prevJIInputEven := 0.0
	jQ := 0.0
	prevJQOdd := 0.0
	prevJQEven := 0.0
	prevJQInputOdd := 0.0
	prevJQInputEven := 0.0
	period := 0.0
	outIdx := 0
	previ2 := 0.0
	prevq2 := 0.0
	Re := 0.0
	Im := 0.0
	i1ForOddPrev3 := 0.0
	i1ForEvenPrev3 := 0.0
	i1ForOddPrev2 := 0.0
	i1ForEvenPrev2 := 0.0
	smoothPeriod := 0.0
	dcPhase := 0.0
	q2 := 0.0
	i2 := 0.0
	for today < len(inReal) {
		adjustedPrevPeriod := (0.075 * period) + 0.54
		todayValue := inReal[today]
		periodWMASub += todayValue
		periodWMASub -= trailingWMAValue
		periodWMASum += todayValue * 4.0
		trailingWMAValue = inReal[trailingWMAIdx]
		trailingWMAIdx++
		smoothedValue = periodWMASum * 0.1
		periodWMASum -= periodWMASub
		hilbertTempReal := 0.0
		smoothPrice[smoothPriceIdx] = smoothedValue
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
			jI -= prevJIEven
			prevJIEven = b * prevJIInputEven
			jI += prevJIEven
			prevJIInputEven = i1ForEvenPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQEven[hilbertIdx]
			jQEven[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQEven
			prevJQEven = b * prevJQInputEven
			jQ += prevJQEven
			prevJQInputEven = q1
			jQ *= adjustedPrevPeriod
			hilbertIdx++
			if hilbertIdx == 3 {
				hilbertIdx = 0
			}
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForEvenPrev3 - jQ)) + (0.8 * previ2)
			i1ForOddPrev3 = i1ForOddPrev2
			i1ForOddPrev2 = detrender
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
			jI -= prevJIOdd
			prevJIOdd = b * prevJIInputOdd
			jI += prevJIOdd
			prevJIInputOdd = i1ForOddPrev3
			jI *= adjustedPrevPeriod
			hilbertTempReal = a * q1
			jQ = -jQOdd[hilbertIdx]
			jQOdd[hilbertIdx] = hilbertTempReal
			jQ += hilbertTempReal
			jQ -= prevJQOdd
			prevJQOdd = b * prevJQInputOdd
			jQ += prevJQOdd
			prevJQInputOdd = q1
			jQ *= adjustedPrevPeriod
			q2 = (0.2 * (q1 + jI)) + (0.8 * prevq2)
			i2 = (0.2 * (i1ForOddPrev3 - jQ)) + (0.8 * previ2)
			i1ForEvenPrev3 = i1ForEvenPrev2
			i1ForEvenPrev2 = detrender
		}
		Re = (0.2 * ((i2 * previ2) + (q2 * prevq2))) + (0.8 * Re)
		Im = (0.2 * ((i2 * prevq2) - (q2 * previ2))) + (0.8 * Im)
		prevq2 = q2
		previ2 = i2
		tempReal = period
		if (Im != 0.0) && (Re != 0.0) {
			period = 360.0 / (math.Atan(Im/Re) * rad2Deg)
		}
		tempReal2 := 1.5 * tempReal
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
		smoothPeriod = (0.33 * period) + (0.67 * smoothPeriod)
		DCPeriod := smoothPeriod + 0.5
		DCPeriodInt := math.Floor(DCPeriod)
		realPart := 0.0
		imagPart := 0.0
		idx := smoothPriceIdx
		for i := 0; i < int(DCPeriodInt); i++ {
			tempReal = (float64(i) * constDeg2RadBy360) / (DCPeriodInt * 1.0)
			tempReal2 = smoothPrice[idx]
			realPart += math.Sin(tempReal) * tempReal2
			imagPart += math.Cos(tempReal) * tempReal2
			if idx == 0 {
				idx = 50 - 1
			} else {
				idx--
			}
		}
		tempReal = math.Abs(imagPart)
		if tempReal > 0.0 {
			dcPhase = math.Atan(realPart/imagPart) * rad2Deg
		} else if tempReal <= 0.01 {
			if realPart < 0.0 {
				dcPhase -= 90.0
			} else if realPart > 0.0 {
				dcPhase += 90.0
			}
		}
		dcPhase += 90.0
		dcPhase += 360.0 / smoothPeriod
		if imagPart < 0.0 {
			dcPhase += 180.0
		}
		if dcPhase > 315.0 {
			dcPhase -= 360.0
		}
		if today >= startIdx {
			outReal[outIdx] = dcPhase
			outIdx++
		}
		smoothPriceIdx++
		if smoothPriceIdx > maxIdxSmoothPrice {
			smoothPriceIdx = 0
		}

		today++
	}
	return outReal
}

/* Statistic Functions */

// Beta - Beta
func Beta(inReal0 []float64, inReal1 []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal0))

	x := 0.0
	y := 0.0
	sSS := 0.0
	sXY := 0.0
	sX := 0.0
	sY := 0.0
	tmpReal := 0.0
	n := 0.0
	nbInitialElementNeeded := inTimePeriod
	startIdx := nbInitialElementNeeded
	trailingIdx := startIdx - nbInitialElementNeeded
	trailingLastPriceX := inReal0[trailingIdx]
	lastPriceX := trailingLastPriceX
	trailingLastPriceY := inReal1[trailingIdx]
	lastPriceY := trailingLastPriceY
	trailingIdx++
	i := trailingIdx
	for i < startIdx {
		tmpReal := inReal0[i]
		x := 0.0
		if !((-0.00000000000001 < lastPriceX) && (lastPriceX < 0.00000000000001)) {
			x = (tmpReal - lastPriceX) / lastPriceX
		}
		lastPriceX = tmpReal
		tmpReal = inReal1[i]
		i++
		y := 0.0
		if !((-0.00000000000001 < lastPriceY) && (lastPriceY < 0.00000000000001)) {
			y = (tmpReal - lastPriceY) / lastPriceY
		}
		lastPriceY = tmpReal
		sSS += x * x
		sXY += x * y
		sX += x
		sY += y
	}
	outIdx := inTimePeriod
	n = float64(inTimePeriod)
	for ok := true; ok; {
		tmpReal = inReal0[i]
		if !((-0.00000000000001 < lastPriceX) && (lastPriceX < 0.00000000000001)) {
			x = (tmpReal - lastPriceX) / lastPriceX
		} else {
			x = 0.0
		}
		lastPriceX = tmpReal
		tmpReal = inReal1[i]
		i++
		if !((-0.00000000000001 < lastPriceY) && (lastPriceY < 0.00000000000001)) {
			y = (tmpReal - lastPriceY) / lastPriceY
		} else {
			y = 0.0
		}
		lastPriceY = tmpReal
		sSS += x * x
		sXY += x * y
		sX += x
		sY += y
		tmpReal = inReal0[trailingIdx]
		if !(((-(0.00000000000001)) < trailingLastPriceX) && (trailingLastPriceX < (0.00000000000001))) {
			x = (tmpReal - trailingLastPriceX) / trailingLastPriceX
		} else {
			x = 0.0
		}
		trailingLastPriceX = tmpReal
		tmpReal = inReal1[trailingIdx]
		trailingIdx++
		if !(((-(0.00000000000001)) < trailingLastPriceY) && (trailingLastPriceY < (0.00000000000001))) {
			y = (tmpReal - trailingLastPriceY) / trailingLastPriceY
		} else {
			y = 0.0
		}
		trailingLastPriceY = tmpReal
		tmpReal = (n * sSS) - (sX * sX)
		if !(((-(0.00000000000001)) < tmpReal) && (tmpReal < (0.00000000000001))) {
			outReal[outIdx] = ((n * sXY) - (sX * sY)) / tmpReal
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
		sSS -= x * x
		sXY -= x * y
		sX -= x
		sY -= y
		ok = i < len(inReal0)
	}

	return outReal
}

// Correl - Pearson's Correlation Coefficient (r)
func Correl(inReal0 []float64, inReal1 []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal0))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	trailingIdx := startIdx - lookbackTotal
	sumXY, sumX, sumY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
	today := trailingIdx
	for today = trailingIdx; today <= startIdx; today++ {
		x := inReal0[today]
		sumX += x
		sumX2 += x * x
		y := inReal1[today]
		sumXY += x * y
		sumY += y
		sumY2 += y * y
	}
	trailingX := inReal0[trailingIdx]
	trailingY := inReal1[trailingIdx]
	trailingIdx++
	tempReal := (sumX2 - ((sumX * sumX) / inTimePeriodF)) * (sumY2 - ((sumY * sumY) / inTimePeriodF))
	if !(tempReal < 0.00000000000001) {
		outReal[inTimePeriod-1] = (sumXY - ((sumX * sumY) / inTimePeriodF)) / math.Sqrt(tempReal)
	} else {
		outReal[inTimePeriod-1] = 0.0
	}
	outIdx := inTimePeriod
	for today < len(inReal0) {
		sumX -= trailingX
		sumX2 -= trailingX * trailingX
		sumXY -= trailingX * trailingY
		sumY -= trailingY
		sumY2 -= trailingY * trailingY
		x := inReal0[today]
		sumX += x
		sumX2 += x * x
		y := inReal1[today]
		today++
		sumXY += x * y
		sumY += y
		sumY2 += y * y
		trailingX = inReal0[trailingIdx]
		trailingY = inReal1[trailingIdx]
		trailingIdx++
		tempReal = (sumX2 - ((sumX * sumX) / inTimePeriodF)) * (sumY2 - ((sumY * sumY) / inTimePeriodF))
		if !(tempReal < (0.00000000000001)) {
			outReal[outIdx] = (sumXY - ((sumX * sumY) / inTimePeriodF)) / math.Sqrt(tempReal)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}
	return outReal
}

// StdDev - Standard Deviation
func StdDev(inReal []float64, inTimePeriod int, inNbDev float64) []float64 {
	outReal := Var(inReal, inTimePeriod)

	if inNbDev != 1.0 {
		for i := 0; i < len(inReal); i++ {
			tempReal := outReal[i]
			if !(tempReal < 0.00000000000001) {
				outReal[i] = math.Sqrt(tempReal) * inNbDev
			} else {
				outReal[i] = 0.0
			}
		}
	} else {
		for i := 0; i < len(inReal); i++ {
			tempReal := outReal[i]
			if !(tempReal < 0.00000000000001) {
				outReal[i] = math.Sqrt(tempReal)
			} else {
				outReal[i] = 0.0
			}
		}
	}
	return outReal
}

// Tsf - Time Series Forecast
func Tsf(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1.0) * 0.5
	sumXSqr := inTimePeriodF * (inTimePeriodF - 1) * (2*inTimePeriodF - 1) / 6
	divisor := sumX*sumX - inTimePeriodF*sumXSqr
	//initialize values of sumY and sumXY over first (inTimePeriod) input values
	sumXY := 0.0
	sumY := 0.0
	i := inTimePeriod
	for i != 0 {
		i--
		tempValue1 := inReal[today-i]
		sumY += tempValue1
		sumXY += float64(i) * tempValue1
	}
	for today < len(inReal) {
		//sumX and sumXY are already available for first output value
		if today > startIdx-1 {
			tempValue2 := inReal[today-inTimePeriod]
			sumXY += sumY - inTimePeriodF*tempValue2
			sumY += inReal[today] - tempValue2
		}
		m := (inTimePeriodF*sumXY - sumX*sumY) / divisor
		b := (sumY - m*sumX) / inTimePeriodF
		outReal[outIdx] = b + m*inTimePeriodF
		today++
		outIdx++
	}
	return outReal
}

// Var - Variance
func Var(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	periodTotal1 := 0.0
	periodTotal2 := 0.0
	trailingIdx := startIdx - nbInitialElementNeeded
	i := trailingIdx
	if inTimePeriod > 1 {
		for i < startIdx {
			tempReal := inReal[i]
			periodTotal1 += tempReal
			tempReal *= tempReal
			periodTotal2 += tempReal
			i++
		}
	}
	outIdx := startIdx
	for ok := true; ok; {
		tempReal := inReal[i]
		periodTotal1 += tempReal
		tempReal *= tempReal
		periodTotal2 += tempReal
		meanValue1 := periodTotal1 / float64(inTimePeriod)
		meanValue2 := periodTotal2 / float64(inTimePeriod)
		tempReal = inReal[trailingIdx]
		periodTotal1 -= tempReal
		tempReal *= tempReal
		periodTotal2 -= tempReal
		outReal[outIdx] = meanValue2 - meanValue1*meanValue1
		i++
		trailingIdx++
		outIdx++
		ok = i < len(inReal)
	}
	return outReal
}

// HeikinashiCandles - from candle values extracts heikinashi candle values.
//
// Returns highs, opens, closes and lows of the heikinashi candles (in this order).
//
//    NOTE: The number of Heikin-Ashi candles will always be one less than the number of provided candles, due to the fact
//          that a previous candle is necessary to calculate the Heikin-Ashi candle, therefore the first provided candle is not considered
//          as "current candle" in the algorithm, but only as "previous candle".
func HeikinashiCandles(highs []float64, opens []float64, closes []float64, lows []float64) ([]float64, []float64, []float64, []float64) {
	N := len(highs)

	heikinHighs := make([]float64, N)
	heikinOpens := make([]float64, N)
	heikinCloses := make([]float64, N)
	heikinLows := make([]float64, N)

	for currentCandle := 1; currentCandle < N; currentCandle++ {
		previousCandle := currentCandle - 1

		heikinHighs[currentCandle] = math.Max(highs[currentCandle], math.Max(opens[currentCandle], closes[currentCandle]))
		heikinOpens[currentCandle] = (opens[previousCandle] + closes[previousCandle]) / 2
		heikinCloses[currentCandle] = (highs[currentCandle] + opens[currentCandle] + closes[currentCandle] + lows[currentCandle]) / 4
		heikinLows[currentCandle] = math.Min(highs[currentCandle], math.Min(opens[currentCandle], closes[currentCandle]))
	}

	return heikinHighs, heikinOpens, heikinCloses, heikinLows
}

// Hlc3 returns the Hlc3 values
//
//     NOTE: Every Hlc item is defined as follows : (high + low + close) / 3
//           It is used as AvgPrice candle.
func Hlc3(highs []float64, lows []float64, closes []float64) []float64 {
	N := len(highs)
	result := make([]float64, N)
	for i := range highs {
		result[i] = (highs[i] + lows[i] + closes[i]) / 3
	}

	return result
}

// Crossover returns true if series1 is crossing over series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for buy signals.
//          It assumes first values are the most recent.
//          The crossover function does not use most recent value, since usually it's not a complete candle.
//          The second recent values and the previous are used, instead.
func Crossover(series1 []float64, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-2] <= series2[N-2] && series1[N-1] > series2[N-1]
}

// Crossunder returns true if series1 is crossing under series2.
//
//    NOTE: Usually this is used with Media Average Series to check if it crosses for sell signals.
func Crossunder(series1 []float64, series2 []float64) bool {
	if len(series1) < 3 || len(series2) < 3 {
		return false
	}

	N := len(series1)

	return series1[N-1] <= series2[N-1] && series1[N-2] > series2[N-2]
}
