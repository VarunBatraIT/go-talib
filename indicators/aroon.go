package indicators

// Aroon - Aroon
// aroondown, aroonup = AROON(high, low, timeperiod=14)
func Aroon(inHigh []float64, inLow []float64, inTimePeriod int) ([]float64, []float64) {
	outAroonUp := make([]float64, len(inHigh))
	outAroonDown := make([]float64, len(inHigh))

	startIdx := inTimePeriod
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - inTimePeriod
	lowestIdx := -1
	highestIdx := -1
	lowest := 0.0
	highest := 0.0
	factor := 100.0 / float64(inTimePeriod)
	for today < len(inHigh) {
		tmp := inLow[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inLow[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmp = inLow[i]
				if tmp <= lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		tmp = inHigh[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inHigh[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmp = inHigh[i]
				if tmp >= highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		outAroonUp[outIdx] = factor * float64(inTimePeriod-(today-highestIdx))
		outAroonDown[outIdx] = factor * float64(inTimePeriod-(today-lowestIdx))
		outIdx++
		trailingIdx++
		today++
	}
	return outAroonDown, outAroonUp
}

// AroonOsc - Aroon Oscillator
func AroonOsc(inHigh []float64, inLow []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inHigh))

	startIdx := inTimePeriod
	outIdx := startIdx
	today := startIdx
	trailingIdx := startIdx - inTimePeriod
	lowestIdx := -1
	highestIdx := -1
	lowest := 0.0
	highest := 0.0
	factor := 100.0 / float64(inTimePeriod)
	for today < len(inHigh) {
		tmp := inLow[today]
		if lowestIdx < trailingIdx {
			lowestIdx = trailingIdx
			lowest = inLow[lowestIdx]
			i := lowestIdx
			i++
			for i <= today {
				tmp = inLow[i]
				if tmp <= lowest {
					lowestIdx = i
					lowest = tmp
				}
				i++
			}
		} else if tmp <= lowest {
			lowestIdx = today
			lowest = tmp
		}
		tmp = inHigh[today]
		if highestIdx < trailingIdx {
			highestIdx = trailingIdx
			highest = inHigh[highestIdx]
			i := highestIdx
			i++
			for i <= today {
				tmp = inHigh[i]
				if tmp >= highest {
					highestIdx = i
					highest = tmp
				}
				i++
			}
		} else if tmp >= highest {
			highestIdx = today
			highest = tmp
		}
		aroon := factor * float64(highestIdx-lowestIdx)
		outReal[outIdx] = aroon
		outIdx++
		trailingIdx++
		today++
	}

	return outReal
}
