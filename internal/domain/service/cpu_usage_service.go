package service

import (
	"timescaledb-benchmark-assignment/internal/domain/model"
	"timescaledb-benchmark-assignment/internal/repository"
)

type CpuUsageService interface {
	//SearchByHostname search by query params
	SearchByHostname(queryParam *model.QueryParam) (*model.CpuUsage, error)
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

func (s cPuUsageServiceImpl) SearchByHostname(queryParam *model.QueryParam) (*model.CpuUsage, error) {

	startTime, err := queryParam.GetStartTime()
	if err != nil {
		return nil, err
	}

	endTime, err := queryParam.GetEndTime()
	if err != nil {
		return nil, err
	}
	return s.repository.SearchByHostname(queryParam.Hostname, startTime, endTime)
}
