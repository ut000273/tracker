PROG=deepin_cve_tracker

build:
	go build -o ${PROG} cmd/main.go

docker:
	docker build -t deepin-cve -f deployments/mysql-Dockerfile .
	docker run -it -d -p 10808:10808  --name tracker  deepin-cve:latest ./deepin_cve_tracker


clean:
	rm -f ${PROG}

rebuild: clean build
