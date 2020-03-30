package indicators

import (
	"math"
)

/* Math Transform Functions */

// Acos - Vector Trigonometric ACOS
func Acos(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Acos(inReal[i])
	}
	return outReal
}

// Asin - Vector Trigonometric ASIN
func Asin(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Asin(inReal[i])
	}
	return outReal
}

// Atan - Vector Trigonometric ATAN
func Atan(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Atan(inReal[i])
	}
	return outReal
}

// Ceil - Vector CEIL
func Ceil(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Ceil(inReal[i])
	}
	return outReal
}

// Cos - Vector Trigonometric COS
func Cos(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Cos(inReal[i])
	}
	return outReal
}

// Cosh - Vector Trigonometric COSH
func Cosh(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Cosh(inReal[i])
	}
	return outReal
}

// Exp - Vector atrithmetic EXP
func Exp(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Exp(inReal[i])
	}
	return outReal
}

// Floor - Vector FLOOR
func Floor(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Floor(inReal[i])
	}
	return outReal
}

// Ln - Vector natural log LN
func Ln(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Log(inReal[i])
	}
	return outReal
}

// Log10 - Vector LOG10
func Log10(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Log10(inReal[i])
	}
	return outReal
}

// Sin - Vector Trigonometric SIN
func Sin(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Sin(inReal[i])
	}
	return outReal
}

// Sinh - Vector Trigonometric SINH
func Sinh(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Sinh(inReal[i])
	}
	return outReal
}

// Sqrt - Vector SQRT
func Sqrt(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Sqrt(inReal[i])
	}
	return outReal
}

// Tan - Vector Trigonometric TAN
func Tan(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Tan(inReal[i])
	}
	return outReal
}

// Tanh - Vector Trigonometric TANH
func Tanh(inReal []float64) []float64 {
	outReal := make([]float64, len(inReal))
	for i := 0; i < len(inReal); i++ {
		outReal[i] = math.Tanh(inReal[i])
	}
	return outReal
}

/* Math Operator Functions */

// Add - Vector arithmetic addition
func Add(inReal0, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] + inReal1[i]
	}
	return outReal
}

// Div - Vector arithmetic division
func Div(inReal0, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] / inReal1[i]
	}
	return outReal
}

// Max - Highest value over a period
func Max(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0

	for today < len(outReal) {
		tmp := inReal[today]

		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outReal[outIdx] = highest
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

// MaxIndex - Index of highest value over a specified period
func MaxIndex(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	for today < len(inReal) {
		tmp := inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp > highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outReal[outIdx] = float64(highestIdx)
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

// Min - Lowest value over a period
func Min(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	lowestIdx := -1
	lowest := 0.0
	for today < len(outReal) {
		tmp := inReal[today]

		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		outReal[outIdx] = lowest
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}

// MinIndex - Index of lowest value over a specified period
func MinIndex(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	if inTimePeriod < 2 {
		return outReal
	}

	nbInitialElementNeeded := inTimePeriod - 1
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmp := inReal[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx + 1
			for i <= today {
				tmp = inReal[i]
				if tmp < lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		outReal[outIdx] = float64(lowestIdx)
		outIdx++
		trailingIdx++
		today++
	}
	return outReal
}

// MinMax - Lowest and highest values over a specified period
func MinMax(inReal []float64, inTimePeriod int) (outMin, outMax []float64) {
	outMin = make([]float64, len(inReal))
	outMax = make([]float64, len(inReal))

	nbInitialElementNeeded := (inTimePeriod - 1)
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmpLow, tmpHigh := inReal[today], inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmpHigh = inReal[i]
				if tmpHigh > highest {
					highestIdx = i
					highest = tmpHigh
				}
				i++
			}
		} else if tmpHigh >= highest {
			highestIdx = today
			highest = tmpHigh
		}
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmpLow = inReal[i]
				if tmpLow < lowest {
					lowestIdx = i
					lowest = tmpLow
				}
				i++
			}
		} else if tmpLow <= lowest {
			lowestIdx = today
			lowest = tmpLow
		}
		outMax[outIdx] = highest
		outMin[outIdx] = lowest
		outIdx++
		trailingIdx++
		today++
	}
	return outMin, outMax
}

// MinMaxIndex - Indexes of lowest and highest values over a specified period
func MinMaxIndex(inReal []float64, inTimePeriod int) (outMinIdx, outMaxIdx []float64) {
	outMinIdx = make([]float64, len(inReal))
	outMaxIdx = make([]float64, len(inReal))

	nbInitialElementNeeded := (inTimePeriod - 1)
	startIdx := nbInitialElementNeeded
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - nbInitialElementNeeded
	highestIdx := -1
	highest := 0.0
	lowestIdx := -1
	lowest := 0.0
	for today < len(inReal) {
		tmpLow, tmpHigh := inReal[today], inReal[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inReal[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmpHigh = inReal[i]
				if tmpHigh > highest {
					highestIdx = i
					highest = tmpHigh
				}
				i++
			}
		} else if tmpHigh >= highest {
			highestIdx = today
			highest = tmpHigh
		}
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inReal[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmpLow = inReal[i]
				if tmpLow < lowest {
					lowestIdx = i
					lowest = tmpLow
				}
				i++
			}
		} else if tmpLow <= lowest {
			lowestIdx = today
			lowest = tmpLow
		}
		outMaxIdx[outIdx] = float64(highestIdx)
		outMinIdx[outIdx] = float64(lowestIdx)
		outIdx++
		trailingIdx++
		today++
	}
	return outMinIdx, outMaxIdx
}

// Mult - Vector arithmetic multiply
func Mult(inReal0, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] * inReal1[i]
	}
	return outReal
}

// Sub - Vector arithmetic subtraction
func Sub(inReal0, inReal1 []float64) []float64 {
	outReal := make([]float64, len(inReal0))
	for i := 0; i < len(inReal0); i++ {
		outReal[i] = inReal0[i] - inReal1[i]
	}
	return outReal
}

// Sum - Vector summation
func Sum(inReal []float64, inTimePeriod int) []float64 {
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
	for i < len(inReal) {
		periodTotal += inReal[i]
		tempReal := periodTotal
		periodTotal -= inReal[trailingIdx]
		outReal[outIdx] = tempReal
		i++
		trailingIdx++
		outIdx++
	}

	return outReal
}

// Bop - Balance Of Power
func Bop(inOpen, inHigh, inLow, inClose []float64) []float64 {
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

// OBV implements the on balance volume indicator
func OBV(ohlc [][]float64, cryptowatch bool) []float64 {
	var result []float64

	for x := range ohlc {
		if x == 0 {
			if cryptowatch {
				result = append(result, 0)
			} else {
				result = append(result, ohlc[x][5])
			}
			continue
		}
		// nolint gocritic ifElseChain: switch statement complexity not eeded
		if ohlc[x][4] > ohlc[x-1][4] {
			result = append(result, result[x-1]+ohlc[x][5])
		} else if ohlc[x][4] < ohlc[x-1][4] {
			result = append(result, result[x-1]-ohlc[x][5])
		} else if ohlc[x][4] == ohlc[x-1][4] {
			result = append(result, result[x-1])
		}
	}
	return result
}

// Beta - Beta
func Beta(inReal0, inReal1 []float64, inTimePeriod int) []float64 {
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
		tmpReal = inReal0[i]
		x = 0.0
		if !((-0.00000000000001 < lastPriceX) && (lastPriceX < 0.00000000000001)) {
			x = (tmpReal - lastPriceX) / lastPriceX
		}
		lastPriceX = tmpReal
		tmpReal = inReal1[i]
		i++
		y = 0.0
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

// Trix - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA
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

// WillR - Williams' %R
func WillR(inHigh, inLow, inClose []float64, inTimePeriod int) []float64 {
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

// Ad - Chaikin A/D Line
func Ad(inHigh, inLow, inClose, inVolume []float64) []float64 {
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
		closeVal := inClose[currentBar]
		if tmp > 0.0 {
			ad += (((closeVal - low) - (high - closeVal)) / tmp) * (inVolume[currentBar])
		}
		outReal[outIdx] = ad
		outIdx++
		currentBar++
		nbBar--
	}
	return outReal
}

// Correl - Pearson's Correlation Coefficient (r)
func Correl(inReal0, inReal1 []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal0))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod - 1
	startIdx := lookbackTotal
	trailingIdx := startIdx - lookbackTotal
	sumXY, sumX, sumY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
	//nolint // to do fix me
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

// HeikinashiCandles - from candle values extracts heikinashi candle values.
//
// Returns highs, opens, closes and lows of the heikinashi candles (in this order).
//
//    NOTE: The number of Heikin-Ashi candles will always be one less than the number of provided candles, due to the fact
//          that a previous candle is necessary to calculate the Heikin-Ashi candle, therefore the first provided candle is not considered
//          as "current candle" in the algorithm, but only as "previous candle".
func HeikinashiCandles(highs, opens, closes, lows []float64) (heikinHighs, heikinOpens, heikinCloses, heikinLows []float64) {
	N := len(highs)

	heikinHighs = make([]float64, N)
	heikinOpens = make([]float64, N)
	heikinCloses = make([]float64, N)
	heikinLows = make([]float64, N)

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
func Hlc3(highs, lows, closes []float64) []float64 {
	N := len(highs)
	result := make([]float64, N)
	for i := range highs {
		result[i] = (highs[i] + lows[i] + closes[i]) / 3
	}

	return result
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
