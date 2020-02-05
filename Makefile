PROJECT_ID := nnnewta

docker-push: docker-build
	@docker push gcr.io/$(PROJECT_ID)/newta

docker-build:
	@docker build -t gcr.io/$(PROJECT_ID)/newta -f build/docker/newta/Dockerfile .
