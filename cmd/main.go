package main

import (
	"flag"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"timescaledb-benchmark-assignment/infra/db/timescaledb"
	"timescaledb-benchmark-assignment/internal/application"
	"timescaledb-benchmark-assignment/internal/domain/service"
	"timescaledb-benchmark-assignment/internal/repository"
)

var (
	queryFile string
	workers   int
)

func init() {

	flag.StringVar(&queryFile, "file", "", "Path to file containing queries to execute")
	flag.IntVar(&workers, "workers", 5, "Number of workers")

	flag.Parse()

	log.SetLevel(log.DebugLevel)
}

// main start application
func main() {
	log.Info("Starting application")
	db, err := timescaledb.GetDb()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	cpuUsageRepository := repository.NewCpuUsageRepository(db)

	cPuUsageService := service.NewCPuUsageService(cpuUsageRepository)

	timescaleDbQueryBenchmark := application.NewTimescaleDbQueryBenchmark(cPuUsageService)
	if err = timescaleDbQueryBenchmark.Run(queryFile, workers); err != nil {
		log.Fatalf("Unable to run application: %s", err)
	}

}
