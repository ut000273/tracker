PROG=deepin_cve_tracker

build:
	go build -o ${PROG} cmd/main.go

docker:
	docker run -d --name mysql-cve -p 32680:3306 --restart=always -v /data/db/mariadb:/var/lib/mysql mariadb
	docker build -t cve-tracker -f deployments/mysql-Dockerfile .
	docker run -it -d -p 10808:10808 --restart=always  --name tracker  cve-tracker:latest ./deepin_cve_tracker

docker-clean:
	docker stop tracker mysql-cve
	docker rm tracker mysql-cve
	docker image rm cve-tracker


clean:
	rm -f ${PROG}

rebuild: clean build

docker-rebuild: docker-clean docker
