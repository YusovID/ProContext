package main

import (
	"fmt"
	"log"
	"task_3/internal/calculator"
	"task_3/internal/model"
	"task_3/internal/work"
)

const (
	numberOfDays = 90
	workerCount  = 10
)

func main() {
	valutesMap := make(map[string][]model.JsonValute)
	errorsMap := make(map[string][]error)

	work.Dispatcher(&valutesMap, &errorsMap, numberOfDays, workerCount)

	if len(valutesMap) == 0 {
		log.Print("No data fetched")
		return
	}

	if len(errorsMap) > 0 {
		log.Println("--- Conversion Errors ---")
		for date, errors := range errorsMap {
			log.Printf("Date: %s, errors: %v", date, errors)
		}
		log.Println("--- End of Conversion Errors ---")
	}

	maxValute := calculator.MaxValute(valutesMap)
	minValute := calculator.MinValute(valutesMap)
	avgValue := calculator.AvgValue(valutesMap)

	fmt.Printf("Max valute:\n  Value: %f\n  Name: %s\n  Date: %s\n", maxValute.VunitRate, maxValute.Name, maxValute.Date)
	fmt.Printf("Min valute:\n  Value: %f\n  Name: %s\n  Date: %s\n", minValute.VunitRate, minValute.Name, minValute.Date)
	fmt.Printf("Avg value: %f\n", avgValue)
}
