# Makefile
docker-build:
	docker build -t skill-marketplace:latest -f deploy/docker/Dockerfile .
# Makefile
k8s-deploy:
	kubectl apply -f deploy/k8s