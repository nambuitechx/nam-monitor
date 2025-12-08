.PHONY: build push run down login


SCYLLA_IMAGE = nam-monitor/scylla
POSTGRES_IMAGE = nam-monitor/postgres
BACKEND_IMAGE = nam-monitor/backend

TAG = 1.0.0
REGION = ap-southeast-1

ACCOUNT_ID = 490004621103
ECR_REGISTRY = $(ACCOUNT_ID).dkr.ecr.$(REGION).amazonaws.com

build-scylla:
	docker build -t $(SCYLLA_IMAGE):$(TAG) ./scylla

build-postgres:
	docker build -t $(POSTGRES_IMAGE):$(TAG) ./postgres

build-backend:
	docker build -t $(BACKEND_IMAGE):$(TAG) ./backend

build: build-postgres build-backend

push-scylla: login
	docker tag $(SCYLLA_IMAGE):$(TAG) $(ECR_REGISTRY)/$(SCYLLA_IMAGE):$(TAG)
	docker push $(ECR_REGISTRY)/$(SCYLLA_IMAGE):$(TAG)

push-postgres: login
	docker tag $(POSTGRES_IMAGE):$(TAG) $(ECR_REGISTRY)/$(POSTGRES_IMAGE):$(TAG)
	docker push $(ECR_REGISTRY)/$(POSTGRES_IMAGE):$(TAG)

push-backend: login
	docker tag $(BACKEND_IMAGE):$(TAG) $(ECR_REGISTRY)/$(BACKEND_IMAGE):$(TAG)
	docker push $(ECR_REGISTRY)/$(BACKEND_IMAGE):$(TAG)

run:
	docker compose up -d

down:
	docker compose down --volumes

login:
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)
