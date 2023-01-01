package application

import (
	"errors"
	"fmt"
	"math"
	"timescaledb-benchmark-assignment/internal/common/query_params"
	"timescaledb-benchmark-assignment/internal/common/util"
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/internal/domain/service"
	"timescaledb-benchmark-assignment/internal/worker_pool"
)

type TimescaleDbQueryBenchmarkApp struct {
	cpuUsageService service.CpuUsageService
}

// NewTimescaleDbQueryBenchmark create new TimescaleDbQueryBenchmarkApp
func NewTimescaleDbQueryBenchmark(cpuUsageService service.CpuUsageService) *TimescaleDbQueryBenchmarkApp {
	return &TimescaleDbQueryBenchmarkApp{cpuUsageService: cpuUsageService}
}

// Run start TimescaleDbQueryBenchmark Application
func (t TimescaleDbQueryBenchmarkApp) Run(queryFile string,
	workers int) error {

	if workers <= 0 {
		return errors.New("workers should be greater than 0")
	}
	//Read  query params file
	queryParams, err := query_params.FromFile(queryFile)

	if err != nil {
		return err
	}

	//Create workers pool
	resultPool := make(chan worker_pool.Result, len(queryParams))
	defer close(resultPool)

	workerPool := worker_pool.NewQueryWorkersPool(resultPool, t.cpuUsageService)

	//Send the query param list to the workers pool
	if err := workerPool.ProcessQueries(queryParams, workers); err != nil {
		return err
	}

	//Process query results
	processResults(queryParams, resultPool)

	return nil
}

// processResults process results and print report
func processResults(queries []*model.QueryParam, resultPool chan worker_pool.Result) {
	maxTime := 0.0
	minTime := math.MaxFloat64
	sumTime := 0.0
	queryTimes := make([]float64, 0)

	numQueries := len(queries)
	for i := 1; i <= numQueries; i++ {
		result := <-resultPool
		maxTime = math.Max(maxTime, result.ElapsedTime)
		minTime = math.Min(minTime, result.ElapsedTime)
		sumTime += result.ElapsedTime
		queryTimes = append(queryTimes, result.ElapsedTime)
		fmt.Println("----------------------------------------------------")
		fmt.Println("Host:", result.CpuUsage.Host)
		fmt.Println("Max Usage:", result.CpuUsage.Max)
		fmt.Println("Min Usage:", result.CpuUsage.Min)
		fmt.Println("Processing Time: ", result.ElapsedTime, "ms")
		fmt.Println("Query processed by:", "Worker", result.WorkerId)
		fmt.Println("----------------------------------------------------")
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("Total of queries processed:", numQueries)
	fmt.Println("Total Processing Time:", sumTime, "ms")
	fmt.Println("Minimum Query Time:", minTime, "ms")
	fmt.Println("Maximum Query Time:", maxTime, "ms")
	fmt.Println("Average Query Time:", sumTime/float64(numQueries), "ms")
	fmt.Println("Median Query Time:", util.Median(queryTimes), "ms")
	fmt.Println("----------------------------------------------------")
}
