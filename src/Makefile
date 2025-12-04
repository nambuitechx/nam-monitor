.PHONY: build push run down login

BACKEND_IMAGE = nam-monitor/backend
SCYLLA_IMAGE = nam-monitor/scylla

TAG = 1.0.0
REGION = ap-southeast-1

ACCOUNT_ID = 490004621103
ECR_REGISTRY = $(ACCOUNT_ID).dkr.ecr.$(REGION).amazonaws.com

build-backend:
	docker build -t $(BACKEND_IMAGE):$(TAG) ./backend

build-scylla:
	docker build -t $(SCYLLA_IMAGE):$(TAG) ./scylla

build: build-backend build-scylla

push-backend: login
	docker tag $(BACKEND_IMAGE):$(TAG) $(ECR_REGISTRY)/$(BACKEND_IMAGE):$(TAG)
	docker push $(ECR_REGISTRY)/$(BACKEND_IMAGE):$(TAG)

push-scylla: login
	docker tag $(SCYLLA_IMAGE):$(TAG) $(ECR_REGISTRY)/$(SCYLLA_IMAGE):$(TAG)
	docker push $(ECR_REGISTRY)/$(SCYLLA_IMAGE):$(TAG)

run:
	docker compose up -d

down:
	docker compose down --volumes

login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)
