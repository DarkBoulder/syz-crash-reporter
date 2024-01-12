all:
	go build ./syz-crash-reporter.go
start:
	./syz-crash-reporter -config=config
