## Timescaledb Benchmark Assignment

Implement a command line tool that can be used to benchmark SELECT query performance across multiple workers/clients against a TimescaleDB instance. The tool should take as its input a CSV file (whose format is specified below) and a flag to specify the number of concurrent workers.
Your tool should take the CSV row values hostname, start time, end time and use them to
generate a SQL query for each row that returns the max cpu usage and min cpu usage of the given hostname for every minute in the time range specified by the start time and end time.
Each query should then be executed by one of the concurrent workers your tool creates, with the constraint that queries for the same hostname be executed by the same worker each time.
Note that the constraint does not mean that the worker _only_ executes for that hostname (i.e., it can execute for multiple hostnames).
After processing all the queries specified by the parameters in the CSV file, the tool should output a summary with the following stats:

* of queries processed,
* total processing time across all queries,
* the minimum query time (for a single query),
* the median query time,
* the average query time,
* and the maximum query time.

## Solution

![solution architecture diagram](timescaledb-benchmark-assignment-solution.png "Solution Architecture Diagram")

### Process File

Read the file, validate and parse it to a list of queries. Create the workers pool according to the number of workers informed by the user,  and then send the list to the workers pool. Each query param represents a task to be processed on the workers pool.

### Worker Pool
Pool of workers where each worker has its own task queue. Queries with the same hostname are always queued in the same task queue. In other words, queries with the same hostname are always processed by the same worker.

### Add Task

Check whether any worker has already processed a query with the same hostname. If so, it will put the query in the worker task queue, otherwise, it chooses one of the workers that has fewer tasks in its queue. By doing this, it will guarantee that queries for the same hostname will be executed by the same worker each time.

### Process Query

Worker process query from its own task queue, calculate the elapsed time and add the result to results pool.


### Process Results

After all workers process all queries, process the result for each query and then output a summary with the following stats:

* of queries processed,
* total processing time across all queries,
* the minimum query time (for a single query),
* the median query time,
* the average query time,
* and the maximum query time.

## Prerequisite
Golang  1.19.4<br/>
GoMock v1.6.0 <br/>
Docker 20.11.1<br/>
GNU Make 3.81<br/>
TimescaleDB<br/>

## Technologies
Golang, Gocsv, Golang-migrate, GoMock, Pq (Pure Go Postgres Driver), Testify<br/>
Docker, git, GNU Make, TimescaleDB<br/>

## Installation
In order to simplify the installation process, the application will run as a docker container. So, the query params files that you want to test have to be accessible from the container. 
The file query_param.csv provided with the assignment will be copied to the container directory /timescaleBD automatically.
Just in case you want to test any other files, you have to copy them to tests/data, and it will be copied to the container directory /timescaleBD automatically.
### Download
To clone the repository, run:

``` 
git clone git@github.com:ergildo/timescaledb-benchmark-assignment.git && cd timescaledb-benchmark-assignment

```

### Configurations
To configure database, edit the file .env

``` 
DB_HOST=<< Database host >>
DB_PORT=<< Database port >>
DB_NAME=<< Database name >>
DB_USER=<< Database username >>
DB_PASSWORD=<< User password >>
DB_MAX_CONNECTIONS=<< Max open connections >>
```

### Notes
DB_MAX_CONNECTIONS has to be less or equals the number of connections available on your database otherwise you will get the follow error:

``` 
error when preparing query:pq: remaining connection slots are reserved for non-replication superuser connections

```

If you get this error, you can fix it by increasing the number of max connections on timescaleDB platform. 

### Build
To build docker image, run:

``` 
docker build --tag timescaledb-benchmark-query:latest .

```

### Start Application
To start the application, run: 

``` 
docker run -ti --env-file .env timescaledb-benchmark-query:latest

```

### Database migrations
To execute database migrations, run:

``` 
benchmark-migrations

```

### Usage
To use the application, run:

``` 
benchmark-query -file=query_params.csv -workers=16

```

Where **-file** is the path to a file containing queries to execute **-workers** is the number of workers that will execute the queries.

## Run local
Just in case you want to run the application on your computer instead of running as a docker container,  you need to set up the fallow environment variables: 

``` 
export DB_HOST=<< Database host >>
export DB_PORT=<< Database port >>
export DB_NAME=<< Database name >>
export DB_USER=<< Database username >>
export DB_PASSWORD=<< User password >>
export DB_MAX_CONNECTIONS=<< Max open connections >>

```
Make sure that you have Golang installed, see prerequisite section.

### Database migrations
To execute database migrations, run:

``` 
go run ./infra/db/migrations/...

```

### Run
To run the application, run:

``` 
go run ./cmd/... -file=tests/data/query_params.csv -workers=16

```

## Tests
To execute tests, run:

``` 
go test ./internal/...

```

## Contacts
#### If you have any questions, please contact me

**e-mail:** ergildo@gmail<br/>

**whatsapp:** +46 76 081 36 43<br/>
