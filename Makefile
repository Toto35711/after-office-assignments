.PHONY: test-day-1 run-day-1

test-day-1:
	go test -v -cover ./day-1/...

run-day-1:
	go run ./day-1 main.go
