package main

import "errors"

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("max(): requires at least one value")
	}
	maxVal := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > maxVal {
			maxVal = vals[i]
		}
	}
	return maxVal, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("min(): requires at least one value")
	}
	minVal := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < minVal {
			minVal = vals[i]
		}
	}
	return minVal, nil
}

func max2(val int, vals ...int) int {
	maxVal := val
	for _, v := range vals {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func min2(val int, vals ...int) int {
	minVal := val
	for _, v := range vals {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
