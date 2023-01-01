package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"timescaledb-benchmark-assignment/mocks"
	test_commons "timescaledb-benchmark-assignment/tests/commons"
)

const hostname = "host_000001"

func TestSearchByParamsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockCpuUsageRepository(ctrl)
	service := NewCPuUsageService(repository)
	query := test_commons.GetQuery(hostname)
	cpuUsage, err := test_commons.GetCpuUsage(query)
	startTime, _ := query.GetStartTime()
	endTime, _ := query.GetEndTime()
	repository.EXPECT().SearchByHostname(query.Hostname, startTime, endTime).Return(cpuUsage, err)
	result, err := service.SearchByHostname(query)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, cpuUsage, result)
}

func TestSearchByParamsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockCpuUsageRepository(ctrl)
	service := NewCPuUsageService(repository)
	query := test_commons.GetQuery(hostname)
	startTime, _ := query.GetStartTime()
	endTime, _ := query.GetEndTime()
	repository.EXPECT().SearchByHostname(query.Hostname, startTime, endTime).Return(nil, errors.New("error when executing query"))
	result, err := service.SearchByHostname(query)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
