generate-mocks:
		mockgen -package=mocks -source=internal/domain/service/cpu_usage_service.go -destination=mocks/cpu_usage_service_mock.go
		mockgen -package=mocks -source=internal/repository/cpu_usage_repository.go -destination=mocks/cpu_usage_repository_mock.go
		mockgen -package=mocks -source=internal/workers_pool/worker_pool.go -destination=mocks/worker_pool_mock.go

