package indicators

import "math"

// LinearReg - Linear Regression
func LinearReg(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1) * 0.5
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
		outReal[outIdx] = b + m*(inTimePeriodF-1)
		outIdx++
		today++
	}
	return outReal
}

// LinearRegAngle - Linear Regression Angle
func LinearRegAngle(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1) * 0.5
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
		outReal[outIdx] = math.Atan(m) * (180.0 / math.Pi)
		outIdx++
		today++
	}
	return outReal
}

// LinearRegIntercept - Linear Regression Intercept
func LinearRegIntercept(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1) * 0.5
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
		outReal[outIdx] = (sumY - m*sumX) / inTimePeriodF
		outIdx++
		today++
	}
	return outReal
}

// LinearRegSlope - Linear Regression Slope
func LinearRegSlope(inReal []float64, inTimePeriod int) []float64 {
	outReal := make([]float64, len(inReal))

	inTimePeriodF := float64(inTimePeriod)
	lookbackTotal := inTimePeriod
	startIdx := lookbackTotal
	outIdx := startIdx - 1
	today := startIdx - 1
	sumX := inTimePeriodF * (inTimePeriodF - 1) * 0.5
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
		outReal[outIdx] = (inTimePeriodF*sumXY - sumX*sumY) / divisor
		outIdx++
		today++
	}
	return outReal
}
