package indicators

import "errors"

// GroupCandles groups a set of candles in another set of candles, basing on a grouping factor.
//
// This is pretty useful if you want to transform, for example, 15min candles into 1h candles using same data.
//
// This avoid calling multiple times the exchange for multiple contexts.
//
// Example:
//     To transform 15 minute candles in 30 minutes candles you have a grouping factor = 2
//
//     To transform 15 minute candles in 1 hour candles you have a grouping factor = 4
//
//     To transform 30 minute candles in 1 hour candles you have a grouping factor = 2
func GroupCandles(highs []float64, opens []float64, closes []float64, lows []float64, groupingFactor int) ([]float64, []float64, []float64, []float64, error) {
	N := len(highs)
	if groupingFactor == 0 {
		return nil, nil, nil, nil, errors.New("Grouping factor must be > 0")
	} else if groupingFactor == 1 {
		return highs, opens, closes, lows, nil // no need to group in this case, return the parameters.
	}
	if N%groupingFactor > 0 {
		return nil, nil, nil, nil, errors.New("Cannot group properly, need a groupingFactor which is a factor of the number of candles")
	}

	groupedN := N / groupingFactor

	groupedHighs := make([]float64, groupedN)
	groupedOpens := make([]float64, groupedN)
	groupedCloses := make([]float64, groupedN)
	groupedLows := make([]float64, groupedN)

	lastOfCurrentGroup := groupingFactor - 1

	k := 0
	for i := 0; i < N; i += groupingFactor { // scan all param candles
		groupedOpens[k] = opens[i]
		groupedCloses[k] = closes[i+lastOfCurrentGroup]

		groupedHighs[k] = highs[i]
		groupedLows[k] = lows[i]

		endOfCurrentGroup := i + lastOfCurrentGroup
		for j := i + 1; j <= endOfCurrentGroup; j++ { // group high lows candles here
			if lows[j] < groupedLows[k] {
				groupedLows[k] = lows[j]
			}
			if highs[j] > groupedHighs[k] {
				groupedHighs[k] = highs[j]
			}
		}
		k++
	}
	return groupedHighs, groupedOpens, groupedCloses, groupedLows, nil
}
