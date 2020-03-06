package indicators

// Stoch - Stochastic
func Stoch(inHigh, inLow, inClose []float64, inFastKPeriod, inSlowKPeriod int, inSlowKMAType MaType, inSlowDPeriod int, inSlowDMAType MaType) ([]float64, []float64) {
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
			i := highestIdx + 1
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
func StochF(inHigh, inLow, inClose []float64, inFastKPeriod, inFastDPeriod int, inFastDMAType MaType) (outFastK, outFastD []float64) {
	outFastK = make([]float64, len(inClose))
	outFastD = make([]float64, len(inClose))

	lookbackK := inFastKPeriod - 1
	lookbackFastD := inFastDPeriod - 1
	lookbackTotal := lookbackK + lookbackFastD
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
