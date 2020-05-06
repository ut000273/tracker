PROG=deepin_cve_tracker

build:
	go build -o ${PROG} cmd/main.go

docker:
	docker build -t deepin-cve -f deployments/mysql-Dockerfile .


clean:
	rm -f ${PROG}

rebuild: clean build
