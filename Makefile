.PHONY: test-day-1

# Target to run tests in day-1/main_test.go
test-day-1:
	go test -v -cover ./day-1/...
