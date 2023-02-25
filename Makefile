.PHONY: run
run:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: restart
restart:
	docker-compose down && docker-compose up -d

.PHONY: build
build:
	docker-compose build

.PHONY: clean
clean:
	docker-compose down --volume

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: open
open:
	open http://localhost:8081

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: mysql
mysql:
	docker-compose run db mysql -h db -P 3306 -u root -p