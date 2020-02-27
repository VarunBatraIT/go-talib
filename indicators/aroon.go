package indicators

// Aroon - Aroon
// down, up = Aroon(high, low, timeperiod)
func Aroon(inHigh, inLow []float64, inTimePeriod int) (down, up []float64) {
	return aroon(inHigh, inLow, inTimePeriod, false)
}

// AroonOsc - Aroon Oscillator
// down, up = AroonOsc(high, low, timeperiod)
func AroonOsc(inHigh, inLow []float64, inTimePeriod int) []float64 {
	up, _ := aroon(inHigh, inLow, inTimePeriod, true)
	return up
}

func aroon(inHigh, inLow []float64, inTimePeriod int, osc bool) (down, up []float64) {
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
		if osc {
			outAroonDown[outIdx] = factor * float64(highestIdx-lowestIdx)
		} else {
			outAroonUp[outIdx] = factor * float64(inTimePeriod-(today-highestIdx))
			outAroonDown[outIdx] = factor * float64(inTimePeriod-(today-lowestIdx))
		}
		outIdx++
		trailingIdx++
		today++
	}
	return outAroonDown, outAroonUp
}
