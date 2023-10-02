
run:
	go run .

test_unit:
	go test -v -tags=unit

test_integration:
	go test -v -tags=unit,integration

test_e2e:
	go test -v -tags=unit,integration,e2e
