package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/mocks"
	test_commons "timescaledb-benchmark-assignment/tests/commons"
)

const hostname = "host_000001"

var invalidQueries = []*model.QueryParam{&model.QueryParam{
	StartTime: "2017-01-01 08:59:22",
	EndTime:   "2017-01-01 09:59:22",
},

	&model.QueryParam{
		Hostname:  "",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},
	&model.QueryParam{
		Hostname:  "  ",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},

	&model.QueryParam{
		Hostname: "hostname",
		EndTime:  "2017-01-01 09:59:22",
	},

	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "",
		EndTime:   "2017-01-01 09:59:22",
	},

	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: " ",
		EndTime:   "2017-01-01 09:59:22",
	},
	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
	},

	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "",
	},
	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   " ",
	},
	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	},
	&model.QueryParam{
		Hostname:  "hostname",
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01",
	},
}

func TestSearchByParamsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockCpuUsageRepository(ctrl)
	service := NewCPuUsageService(repository)
	query := test_commons.GetQuery(hostname)
	cpuUsage, err := test_commons.GetCpuUsage(query)
	startTime, _ := query.GetStartTime()
	endTime, _ := query.GetEndTime()
	repository.EXPECT().SearchByParams(query.Hostname, startTime, endTime).Return(cpuUsage, err)
	result, err := service.SearchByParams(query)
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
	repository.EXPECT().SearchByParams(query.Hostname, startTime, endTime).Return(nil, errors.New("error when executing query"))
	result, err := service.SearchByParams(query)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestSearchByInvalidQueries(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mocks.NewMockCpuUsageRepository(ctrl)
	service := NewCPuUsageService(repository)
	for _, query := range invalidQueries {
		result, err := service.SearchByParams(query)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	}

}
