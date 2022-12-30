package test_commons

import "timescaledb-benchmark-assignment/internal/domain/model"

func GetCpuUsage(query *model.QueryParam) (*model.CpuUsage, error) {
	return &model.CpuUsage{
		Host: query.Hostname,
		Max:  9.1,
		Min:  1.0,
	}, nil
}

func GetQuery(hostname string) *model.QueryParam {
	return &model.QueryParam{
		Hostname:  hostname,
		StartTime: "2017-01-01 08:59:22",
		EndTime:   "2017-01-01 09:59:22",
	}
}
