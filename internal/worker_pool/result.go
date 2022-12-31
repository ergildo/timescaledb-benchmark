package worker_pool

import (
	"timescaledb-benchmark-assignment/internal/domain/model"
)

type Result struct {
	CpuUsage    *model.CpuUsage
	QueryParam  *model.QueryParam
	WorkerId    int
	ElapsedTime float64
}
