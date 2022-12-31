package model

import (
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"

type QueryParam struct {
	Hostname  string `csv:"hostname"`
	StartTime string `csv:"start_time"`
	EndTime   string `csv:"end_time"`
}

func (q QueryParam) GetStartTime() (*time.Time, error) {
	start, err := time.Parse(layout, q.StartTime)
	if err != nil {
		return nil, fmt.Errorf("startTime %s invalid format", q.StartTime)
	}
	return &start, nil
}

func (q QueryParam) GetEndTime() (*time.Time, error) {
	end, err := time.Parse(layout, q.EndTime)
	if err != nil {
		return nil, fmt.Errorf("endtime %s invalid format", q.EndTime)
	}
	return &end, nil
}

func (q QueryParam) Validate() error {

	if len(strings.TrimSpace(q.Hostname)) == 0 {
		return fmt.Errorf("hostname is required")
	}

	if len(strings.TrimSpace(q.StartTime)) == 0 {
		return fmt.Errorf("startTime is required")
	}

	if len(strings.TrimSpace(q.EndTime)) == 0 {
		return fmt.Errorf("endTime is required")
	}

	if _, err := q.GetStartTime(); err != nil {
		return err
	}

	if _, err := q.GetEndTime(); err != nil {
		return err
	}

	return nil
}

func (q QueryParam) String() string {
	return fmt.Sprintf("( hostname: %s, startTime: %s, endTime: %s )", q.Hostname, q.StartTime, q.EndTime)
}
