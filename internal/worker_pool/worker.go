package worker_pool

import (
	"fmt"
	"sync"
	"time"
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/internal/domain/service"
)

type Worker struct {
	id              int
	TaskQueue       chan *model.QueryParam
	hosts           []string
	cpuUsageService service.CpuUsageService
	resultsPool     chan<- Result
	waitGroup       *sync.WaitGroup
}

// NewWorker create new Worker
func NewWorker(id int, results chan<- Result,
	cpuUsageService service.CpuUsageService, bufferSize int, waitGroup *sync.WaitGroup) *Worker {
	taskPool := make(chan *model.QueryParam, bufferSize)
	host := make([]string, 0)
	waitGroup.Add(1)
	return &Worker{
		id:              id,
		TaskQueue:       taskPool,
		resultsPool:     results,
		hosts:           host,
		cpuUsageService: cpuUsageService,
		waitGroup:       waitGroup,
	}

}

// Run execute query params and calculate the elapsed time
func (w *Worker) Run() {

	go func() {

		defer w.waitGroup.Done()
		defer close(w.TaskQueue)

		queueSize := len(w.TaskQueue)

		for i := 1; i <= queueSize; i++ {
			queryParam := <-w.TaskQueue
			fmt.Println("Worker", w.id, "processing ", i, "of ", queueSize, "queries")
			start := time.Now()
			cpuUsages, err := w.cpuUsageService.SearchByHostname(queryParam)
			if err != nil {
				panic(fmt.Sprintf("error while processing query: %s", err))
			}
			elapsedTime := float64(time.Since(start).Milliseconds())

			w.resultsPool <- Result{
				WorkerId:    w.id,
				CpuUsage:    cpuUsages,
				QueryParam:  queryParam,
				ElapsedTime: elapsedTime,
			}

		}

	}()

}

// AddTask add task to worker task queue
func (w *Worker) AddTask(query *model.QueryParam) {
	if !w.IsOnMyTaskQueue(query.Hostname) {
		w.hosts = append(w.hosts, query.Hostname)
	}
	w.TaskQueue <- query
}

// IsOnMyTaskQueue Check if hostname is added to task queue
func (w *Worker) IsOnMyTaskQueue(hostname string) bool {
	for _, host := range w.hosts {
		if host == hostname {
			return true
		}
	}
	return false
}

// QueueSize return task queue size
func (w *Worker) QueueSize() int {
	return len(w.TaskQueue)
}
