default:
	just --list

benchmark:
	go test ./... -bench=. -benchmem -run=^$

test:
	go test ./... --cover -coverprofile reports/coverage

coverage-report: test
	go tool cover -html reports/coverage
