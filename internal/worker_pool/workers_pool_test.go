package worker_pool

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"timescaledb-benchmark-assignment/internal/common/query_params"
	"timescaledb-benchmark-assignment/mocks"
	test_commons "timescaledb-benchmark-assignment/tests/commons"
)

func TestProcessQueriesWithSameHostname(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)
	queryParams, _ := query_params.FromFile("tests/data/query_params.csv")

	for _, query := range queryParams {
		service.EXPECT().SearchByParams(query).Return(test_commons.GetCpuUsage(query)).Times(1)
	}
	resultPool := make(chan Result, len(queryParams))
	workerPool := NewQueryWorkersPool(resultPool, service)
	workerPool.ProcessQueries(queryParams, 3)

	numQueries := len(queryParams)

	assert.Len(t, resultPool, numQueries)

	results := make(map[string]int)

	for i := 1; i <= numQueries; i++ {
		result := <-resultPool
		if v, ok := results[result.QueryParam.Hostname]; ok {
			assert.Equal(t, v, result.WorkerId, "Query with the same hostname must be process by the same worker")
		} else {
			results[result.QueryParam.Hostname] = result.WorkerId
		}
	}

}
