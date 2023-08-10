# Golang
.PHONY: go-test
go-test:
	docker compose run --rm golang make test

go-run-exercise-%:
	docker compose run --rm golang make run-exercise-$*
go-test-exercise-%:
	docker compose run --rm golang make test-exercise-$*

.PHONY: go-goimports
go-goimports:
	docker-compose run --rm golang make goimports

# TypeScript
.PHONY: ts-dev
ts-dev:
	docker compose up typescript

.PHONY: ts-test
ts-test:
	docker compose run --rm typescript yarn test

.PHONY: ts-lint ts-format
ts-lint:
	docker compose run --rm typescript yarn lint
ts-format:
	docker compose run --rm typescript yarn format
