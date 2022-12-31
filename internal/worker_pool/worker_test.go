package worker_pool

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"timescaledb-benchmark-assignment/mocks"
	test_commons "timescaledb-benchmark-assignment/tests/commons"
)

const hostname = "host_000001"
const workerId = 1

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)
	result := make(chan<- Result)
	waitGroup := &sync.WaitGroup{}
	worker := NewWorker(1, result, service, 1, waitGroup)

	query := test_commons.GetQuery(hostname)

	worker.AddTask(query)
	assert.Equal(t, 1, worker.QueueSize())
	assert.True(t, worker.ShouldProcess(hostname))
}

func TestAddRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)
	resultPool := make(chan Result, 1)
	waitGroup := &sync.WaitGroup{}

	worker := NewWorker(workerId, resultPool, service, 1, waitGroup)
	query := test_commons.GetQuery(hostname)

	service.EXPECT().SearchByParams(query).Return(test_commons.GetCpuUsage(query))

	worker.AddTask(query)
	assert.Equal(t, workerId, worker.QueueSize())
	worker.Run()
	waitGroup.Wait()
	assert.Equal(t, 0, worker.QueueSize())

	result := <-resultPool

	assert.Equal(t, result.QueryParam, query)
	assert.Equal(t, result.WorkerId, workerId)
}
