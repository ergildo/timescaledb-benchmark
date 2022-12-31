package query_params

import (
	"errors"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
	"timescaledb-benchmark-assignment/internal/domain/model"
)

// FromFile Read the file, validate and parse it to a list of query params
func FromFile(queryFile string) ([]*model.QueryParam, error) {
	fmt.Println("Reading file", queryFile)
	if len(strings.TrimSpace(queryFile)) <= 0 {
		return nil, errors.New("query-file not specified")
	}

	file, err := os.Open(queryFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open %s: %w", queryFile, err)
	}
	defer file.Close()

	var queryParams []*model.QueryParam

	if err := gocsv.UnmarshalFile(file, &queryParams); err != nil {
		return nil, fmt.Errorf("unable to parse file %s: %w", queryFile, err)
	}

	for _, query := range queryParams {
		err := query.Validate()
		if err != nil {
			fmt.Errorf("invalid file %s: %w", queryFile, err)
		}
	}

	return queryParams, nil

}
