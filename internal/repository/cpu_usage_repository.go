package repository

import (
	"database/sql"
	"fmt"
	"time"
	"timescaledb-benchmark-assignment/internal/domain/model"
)

const query = "select host, max(usage) as max, min(usage) as min from cpu_usage where host = $1 and ts between $2 and $3 group by host"

type CpuUsageRepository interface {
	//SearchByHostname search by query params
	SearchByHostname(hostname string, startTime, endTime *time.Time) (*model.CpuUsage, error)
}

// NewCpuUsageRepository create new CpuUsageRepository
func NewCpuUsageRepository(db *sql.DB) CpuUsageRepository {
	return cpuUsageRepositoryImpl{
		db: db,
	}
}

type cpuUsageRepositoryImpl struct {
	db *sql.DB
}

func (r cpuUsageRepositoryImpl) SearchByHostname(hostname string, startTime, endTime *time.Time) (*model.CpuUsage, error) {
	stm, err := r.db.Prepare(query)

	if err != nil {
		return nil, fmt.Errorf("error when preparing query:%w", err)
	}
	defer stm.Close()
	row := stm.QueryRow(hostname, startTime, endTime)

	if err != nil {
		return nil, fmt.Errorf("error when executing query:%w", err)
	}
	cpuUsage := &model.CpuUsage{}
	row.Scan(&cpuUsage.Host, &cpuUsage.Max, &cpuUsage.Min)

	return cpuUsage, nil
}
