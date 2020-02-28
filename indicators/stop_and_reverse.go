package indicators

import "math"

// Sar - Parabolic SAR
// real = Sar(high, low, acceleration=0, maximum=0)
func Sar(inHigh, inLow []float64, inAcceleration, inMaximum float64) []float64 {
	outReal := make([]float64, len(inHigh))

	af := inAcceleration
	if af > inMaximum {
		af, inAcceleration = inMaximum, inMaximum
	}

	epTemp := MinusDM(inHigh, inLow, 1)
	isLong := 1
	if epTemp[1] > 0 {
		isLong = 0
	}
	startIdx := 1
	outIdx := startIdx
	todayIdx := startIdx
	newHigh := inHigh[todayIdx-1]
	newLow := inLow[todayIdx-1]
	sar, ep := 0.0, 0.0
	if isLong == 1 {
		ep = inHigh[todayIdx]
		sar = newLow
	} else {
		ep = inLow[todayIdx]
		sar = newHigh
	}
	newLow = inLow[todayIdx]
	newHigh = inHigh[todayIdx]
	prevLow := 0.0
	prevHigh := 0.0
	for todayIdx < len(inHigh) {
		prevLow = newLow
		prevHigh = newHigh
		newLow = inLow[todayIdx]
		newHigh = inHigh[todayIdx]
		todayIdx++
		if isLong == 1 {
			if newLow <= sar {
				isLong = 0
				sar = ep
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
				outReal[outIdx] = sar
				outIdx++
				af = inAcceleration
				ep = newLow
				sar += af * (ep - sar)
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
			} else {
				outReal[outIdx] = sar
				outIdx++
				if newHigh > ep {
					ep = newHigh
					af += inAcceleration
					if af > inMaximum {
						af = inMaximum
					}
				}
				sar += af * (ep - sar)
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
			}
		} else {
			if newHigh >= sar {
				isLong = 1
				sar = ep
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
				outReal[outIdx] = sar
				outIdx++
				af = inAcceleration
				ep = newHigh
				sar += af * (ep - sar)
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
			} else {
				outReal[outIdx] = sar
				outIdx++
				if newLow < ep {
					ep = newLow
					af += inAcceleration
					if af > inMaximum {
						af = inMaximum
					}
				}
				sar += af * (ep - sar)
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
			}
		}
	}
	return outReal
}

// SarExt - Parabolic SAR - Extended
// real = SAREXT(high, low, startvalue=0, offsetonreverse=0, accelerationinitlong=0, accelerationlong=0, accelerationmaxlong=0, accelerationinitshort=0, accelerationshort=0, accelerationmaxshort=0)
func SarExt(inHigh, inLow []float64,
	inStartValue,
	inOffsetOnReverse,
	inAccelerationInitLong,
	inAccelerationLong,
	inAccelerationMaxLong,
	inAccelerationInitShort,
	inAccelerationShort,
	inAccelerationMaxShort float64) []float64 {
	outReal := make([]float64, len(inHigh))

	startIdx := 1
	afLong := inAccelerationInitLong
	afShort := inAccelerationInitShort
	if afLong > inAccelerationMaxLong {
		afLong, inAccelerationInitLong = inAccelerationMaxLong, inAccelerationMaxLong
	}

	if inAccelerationLong > inAccelerationMaxLong {
		inAccelerationLong = inAccelerationMaxLong
	}

	if afShort > inAccelerationMaxShort {
		afShort, inAccelerationInitShort = inAccelerationMaxShort, inAccelerationMaxShort
	}

	if inAccelerationShort > inAccelerationMaxShort {
		inAccelerationShort = inAccelerationMaxShort
	}

	isLong := 0
	if inStartValue == 0 {
		epTemp := MinusDM(inHigh, inLow, 1)
		if epTemp[1] > 0 {
			isLong = 0
		} else {
			isLong = 1
		}
	} else if inStartValue > 0 {
		isLong = 1
	}
	outIdx := startIdx
	todayIdx := startIdx
	newHigh := inHigh[todayIdx-1]
	newLow := inLow[todayIdx-1]
	ep := 0.0
	sar := 0.0
	if inStartValue == 0 {
		if isLong == 1 {
			ep = inHigh[todayIdx]
			sar = newLow
		} else {
			ep = inLow[todayIdx]
			sar = newHigh
		}
	} else if inStartValue > 0 {
		ep = inHigh[todayIdx]
		sar = inStartValue
	} else {
		ep = inLow[todayIdx]
		sar = math.Abs(inStartValue)
	}
	newLow = inLow[todayIdx]
	newHigh = inHigh[todayIdx]
	prevLow := 0.0
	prevHigh := 0.0
	for todayIdx < len(inHigh) {
		prevLow = newLow
		prevHigh = newHigh
		newLow = inLow[todayIdx]
		newHigh = inHigh[todayIdx]
		todayIdx++
		if isLong == 1 {
			if newLow <= sar {
				isLong = 0
				sar = ep
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
				if inOffsetOnReverse != 0.0 {
					sar += sar * inOffsetOnReverse
				}
				outReal[outIdx] = -sar
				outIdx++
				afShort = inAccelerationInitShort
				ep = newLow
				sar += afShort * (ep - sar)
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
			} else {
				outReal[outIdx] = sar
				outIdx++
				if newHigh > ep {
					ep = newHigh
					afLong += inAccelerationLong
					if afLong > inAccelerationMaxLong {
						afLong = inAccelerationMaxLong
					}
				}
				sar += afLong * (ep - sar)
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
			}
		} else {
			if newHigh >= sar {
				isLong = 1
				sar = ep
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
				if inOffsetOnReverse != 0.0 {
					sar -= sar * inOffsetOnReverse
				}
				outReal[outIdx] = sar
				outIdx++
				afLong = inAccelerationInitLong
				ep = newHigh
				sar += afLong * (ep - sar)
				if sar > prevLow {
					sar = prevLow
				}
				if sar > newLow {
					sar = newLow
				}
			} else {
				outReal[outIdx] = -sar
				outIdx++
				if newLow < ep {
					ep = newLow
					afShort += inAccelerationShort
					if afShort > inAccelerationMaxShort {
						afShort = inAccelerationMaxShort
					}
				}
				sar += afShort * (ep - sar)
				if sar < prevHigh {
					sar = prevHigh
				}
				if sar < newHigh {
					sar = newHigh
				}
			}
		}
	}
	return outReal
}
