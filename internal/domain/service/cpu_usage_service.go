package service

import (
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/internal/repository"
)

type CpuUsageService interface {
	//SearchByParams search by query params
	SearchByParams(queryParam *model.QueryParam) (*model.CpuUsage, error)
}

// NewCPuUsageService create new CpuUsageService
func NewCPuUsageService(queryRepository repository.CpuUsageRepository) CpuUsageService {
	return cPuUsageServiceImpl{
		repository: queryRepository,
	}
}

type cPuUsageServiceImpl struct {
	repository repository.CpuUsageRepository
}

func (s cPuUsageServiceImpl) SearchByParams(queryParam *model.QueryParam) (*model.CpuUsage, error) {
	err := queryParam.Validate()
	if err != nil {
		return nil, err
	}

	startTime, err := queryParam.GetStartTime()
	if err != nil {
		return nil, err
	}

	endTime, err := queryParam.GetEndTime()
	if err != nil {
		return nil, err
	}
	return s.repository.SearchByParams(queryParam.Hostname, startTime, endTime)
}
