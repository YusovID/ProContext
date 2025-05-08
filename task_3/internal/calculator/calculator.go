package calculator

import (
	"math"
	"task_3/internal/model"
)

type Result struct {
	VunitRate float64
	Name      string
	Date      string
}

func MaxValute(jsonMap map[string][]model.JsonValute) Result {
	result := Result{}
	result.VunitRate = math.Inf(-1)

	for date, valuteSlice := range jsonMap {
		for _, valute := range valuteSlice {
			if valute.VunitRate > result.VunitRate {
				result.Date = date
				result.VunitRate = valute.VunitRate
				result.Name = valute.Name
			}
		}
	}

	return result
}

func MinValute(jsonMap map[string][]model.JsonValute) Result {
	result := Result{}
	result.VunitRate = math.Inf(1)

	for date, valuteSlice := range jsonMap {
		for _, valute := range valuteSlice {
			if valute.VunitRate < result.VunitRate {
				result.Date = date
				result.VunitRate = valute.VunitRate
				result.Name = valute.Name
			}
		}
	}

	return result
}

func AvgValue(jsonMap map[string][]model.JsonValute) float64 {
	var (
		sum   float64
		count int
	)

	for _, valuteSlice := range jsonMap {
		for _, valute := range valuteSlice {
			sum += valute.VunitRate
			count++
		}
	}

	return sum / float64(count)
}
