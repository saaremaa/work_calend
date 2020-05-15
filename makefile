win:
	go build -ldflags "-s -w" -o work-calend.exe
linux:
	go build -ldflags "-s -w" -o work-calend
lint:
	golangci-lint run
docker:
	docker build -t workday:latest .

