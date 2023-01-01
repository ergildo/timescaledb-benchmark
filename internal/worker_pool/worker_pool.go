package worker_pool

import (
	"fmt"
	"sync"
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/internal/domain/service"
)

type WorkersPool interface {
	ProcessQueries(queryParams []*model.QueryParam, numOfWorkers int) error
}

// NewQueryWorkersPool Create new WorkersPool
func NewQueryWorkersPool(results chan<- Result, queryService service.CpuUsageService) WorkersPool {
	workerPool := &queryWorkerPoolImpl{
		resultPool:      results,
		cPuUsageService: queryService,
	}
	return workerPool
}

type queryWorkerPoolImpl struct {
	workers         []*Worker
	resultPool      chan<- Result
	cPuUsageService service.CpuUsageService
}

// ProcessQueries process query params. Divide queries to be processed among workers according to hostname
func (p *queryWorkerPoolImpl) ProcessQueries(queryParams []*model.QueryParam, numOfWorkers int) error {
	waitGroup := &sync.WaitGroup{}

	//Create workers that will process the queries
	p.addWorkers(numOfWorkers, len(queryParams), waitGroup)

	fmt.Println("Adding tasks to workers pool")

	//Define which worker will process each query
	for _, queryParam := range queryParams {
		worker := p.getWorker(queryParam.Hostname)
		worker.AddTask(queryParam)
	}

	fmt.Println("Starting processing tasks")

	//Starting workers
	for _, worker := range p.workers {
		worker.Run()
	}
	//Wait until all worker complete their tasks
	waitGroup.Wait()
	return nil
}

func (p *queryWorkerPoolImpl) addWorkers(numOfWorkers int, bufferSize int, waitGroup *sync.WaitGroup) {
	fmt.Println("Creating workers")
	for i := 1; i <= numOfWorkers; i++ {
		fmt.Println("Creating worker", i)
		worker := NewWorker(i, p.resultPool, p.cPuUsageService, bufferSize, waitGroup)
		p.workers = append(p.workers, worker)
	}
}

func (p *queryWorkerPoolImpl) getWorker(host string) *Worker {
	idleWorker := p.workers[0]
	for _, worker := range p.workers {

		//Check if any worker has already processed a query with the same hostname
		if worker.IsOnMyTaskQueue(host) {
			return worker
		}

		//choose one of the workers that has fewer tasks in its queue
		if worker.QueueSize() < idleWorker.QueueSize() {
			idleWorker = worker
		}
	}
	return idleWorker
}
