lint:
	golangci-lint run

migration:
	@read -p "migration file name:" module; \
	cd kite/app/db/migrations && goose create $$module sql

start-db:
	cd deploy && docker-compose up --build

down:
	docker-compose down --volumes

start:
	docker-compose up --build -d

restart:
	docker-compose up -d --build --no-deps app