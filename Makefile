docker:
	docker-build && docker-run

docker-build:
	docker build -t service-booking .

docker-run:
	docker run -p 8080:8080 service-booking --name booking
