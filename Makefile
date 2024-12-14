dc-up:
	docker-compose -f docker-compose.yml up

dc-build:
	docker build -t go-docker-go-app .