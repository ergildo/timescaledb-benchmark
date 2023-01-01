package application

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"timescaledb-benchmark-assignment/internal/common/query_params"
	"timescaledb-benchmark-assignment/mocks"
	test_commons "timescaledb-benchmark-assignment/tests/commons"
)

const (
	queryFile = "../../tests/data/query_params.csv"
	workers   = 16
)

func TestRunSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)
	queryParams, _ := query_params.FromFile(queryFile)

	for _, query := range queryParams {
		service.EXPECT().SearchByHostname(query).Return(test_commons.GetCpuUsage(query)).Times(1)
	}

	application := NewTimescaleDbQueryBenchmark(service)

	err := application.Run(queryFile, workers)
	assert.Nil(t, err)
}

func TestRunQueryFileNotSpecified(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)

	application := NewTimescaleDbQueryBenchmark(service)

	err := application.Run("", workers)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "query-file not specified")
}

func TestRunQueryFileWorkersInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockCpuUsageService(ctrl)

	application := NewTimescaleDbQueryBenchmark(service)

	err := application.Run(queryFile, 0)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "workers should be greater than 0")
}
