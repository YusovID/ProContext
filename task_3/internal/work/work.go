package work

import (
	"log"
	"sync"
	"task_3/internal/converter"
	"task_3/internal/fetcher"
	"task_3/internal/model"
	"task_3/internal/parser"
	"task_3/pkg/utils"
)

type Job struct {
	Date string
}

func worker(jobs <-chan Job, valutesMap *map[string][]model.JsonValute, errorsMap *map[string][]error, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for job := range jobs {
		dateStr := job.Date

		xmlData, err := fetcher.FetchXML(dateStr)
		if err != nil {
			log.Printf("failed to fetch daily XML: %v", err)
			continue
		}

		valCurs, err := parser.ParseXML(xmlData)
		if err != nil {
			log.Printf("failed to parse XML: %v", err)
			continue
		}

		jsonValutes, errors := converter.ValutesToJSONValutes(valCurs)
		if len(errors) > 0 {
			mu.Lock()
			(*errorsMap)[valCurs.Date] = append((*errorsMap)[valCurs.Date], errors...)
			mu.Unlock()
		}

		mu.Lock()
		(*valutesMap)[valCurs.Date] = jsonValutes
		mu.Unlock()
	}
}

func Dispatcher(valutesMap *map[string][]model.JsonValute, errorsMap *map[string][]error, numberOfDays, workerCount int) {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	jobs := make(chan Job, workerCount)

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(jobs, valutesMap, errorsMap, wg, mu)
	}

	go func() {
		for i := 0; i < numberOfDays; i++ {
			dateStr := utils.CalcDate(-i)
			jobs <- Job{Date: dateStr}
		}
		close(jobs)
	}()

	wg.Wait()
}
