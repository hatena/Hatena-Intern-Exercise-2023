run-exercise-%:
	go run ./exercise-$*/

test-exercise-%:
	go test ./exercise-$*/

.PHONY: test
test:
	go test ./...

.PHONY: goimports
goimports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w ./
