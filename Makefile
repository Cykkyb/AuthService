default:
	go run cmd/main.go --config=config/conf.yaml

migrate:
	go run cmd/migrate/main.go --config=config/conf.yaml