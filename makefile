# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# RSA Keys
# 	To generate a private/public key PEM file.
# 	$ openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# 	$ openssl rsa -pubout -in private.pem -out public.pem

run:
	go run apis/services/sales/main.go | go run apis/tooling/logfmt/main.go

help:
	go run apis/services/sales/main.go --help

version:
	go run apis/services/sales/main.go --version

curl-live:
	curl -il -X GET http://localhost:3000/liveliness

curl-ready:
	curl -il -X GET http://localhost:3000/readiness

curl-error:
	curl -il -X GET http://localhost:3000/testerror

curl-panic:
	curl -il -X GET http://localhost:3000/testpanic

admin:
	go run apis/tooling/admin/main.go

# admin token 
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiJlMDAzMzdkNS1mY2JiLTRkOWYtYTk5Yy0zZGY3ODRlZWQ1ZWMiLCJleHAiOjE3NTk4Mjg1NTUsImlhdCI6MTcyODI5MjU1NSwicm9sZXMiOlsiQURNSU4iXX0.SAgu8CX7H_UYD7SomxdhC-7Gd_unQ1NITeRtWZP4b-ofpzJb1GeowXfGkRG7wKE_3i1x0sj7PGMSLhJ9ywSy56mc1zwITF5pzltihPMp-Qkp-O-wZ9LPKP1tHF2iHQBbgNwHIVkZitwZ--qGsmuxJZpe8my5UZ-8iXoOax__mZ6NNzVbT9HBO4dREtkUkoWC6rDIKSe13qQcu--lI7-5gXY-RIFJkzqFKunTh1Pfb3PtKMYTALdfa4q0oB_Q_HZ0oCAIaOa-wuYMDYH4r40RtmOAfj9--e5y1ULO1xq-HaizXiKWfLwDAoA-Yipa7ojHx-eP3bewC0tAJAWn1OzwfA

# user token
# export TOKEN=eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiJlMDAzMzdkNS1mY2JiLTRkOWYtYTk5Yy0zZGY3ODRlZWQ1ZWMiLCJleHAiOjE3NTk5MDg2NDIsImlhdCI6MTcyODM3MjY0Miwicm9sZXMiOlsiVVNFUiJdfQ.gIhIflZ_jbppoSuWa9zMsxrM_j-EvVzXwglDwvnSlXA-XQVhKFub0iTy_It3NhJQzcrjkltkZ6iIUtqhrrwb7OhujHDG4dyK9bO3jq_vbe4i1INEqT4mjJgnbdLCyvJfT6NjCjYEadBzPgG_9Q_KPP-FBP3KjZAgcmFBocU9k9GLQTNQ_K3DLtoZ2PSwM99WXFdEIRRU_KVDQ8MYB0X_GJRVvVc4F4NZrF__vgiZc5SASZhlt0yOeohdol7rZ2ZE7tJy9HubxmFw16lY1OUMivGNynJKPKXGGsFC4Yc0Oz2y3UnC87UduRRO9rq-c3dOrvhfqXZ-IRBX5WxUEnZ1pQ

curl-auth:
	curl -il \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:3000/testauth"

token:
	curl -il \
	--user "admin@example.com:gophers" http://localhost:6000/auth/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1

curl-auth2:
	curl -il \
	-H "Authorization: Bearer ${TOKEN}" "http://localhost:6000/auth/authenticate"


# ==============================================================================
# Define dependencies

GOLANG          := golang:1.22
ALPINE          := alpine:3.19
KIND            := kindest/node:v1.29.2
POSTGRES        := postgres:16.2
GRAFANA         := grafana/grafana:10.4.0
PROMETHEUS      := prom/prometheus:v2.51.0
TEMPO           := grafana/tempo:2.4.0
LOKI            := grafana/loki:2.9.0
PROMTAIL        := grafana/promtail:2.9.0

KIND_CLUSTER    := ardan-starter-cluster
NAMESPACE       := sales-system
SALES_APP       := sales
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/ardanlabs
VERSION         := 0.0.1
SALES_IMAGE     := $(BASE_IMAGE_NAME)/$(SALES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install github.com/divan/expvarmon@latest
	go install github.com/rakyll/hey@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

dev-brew:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list pgcli || brew install pgcli
	brew list watch || brew install watch

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(ALPINE) & \
	docker pull $(KIND) & \
	docker pull $(POSTGRES) & \
	docker pull $(GRAFANA) & \
	docker pull $(PROMETHEUS) & \
	docker pull $(TEMPO) & \
	docker pull $(LOKI) & \
	docker pull $(PROMTAIL) & \
	wait;

# ==============================================================================
# Building containers

build: sales auth

sales:
	docker build \
		-f zarf/docker/dockerfile.sales \
		-t $(SALES_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

auth:
	docker build \
		-f zarf/docker/dockerfile.auth \
		-t $(AUTH_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)


dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces

# ==============================================================================
dev-load:
	kind load docker-image $(SALES_IMAGE) --name $(KIND_CLUSTER)
	kind load docker-image $(AUTH_IMAGE) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/auth | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(AUTH_APP) --timeout=120s --for=condition=Ready

	kustomize build zarf/k8s/dev/sales | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(SALES_APP) --timeout=120s --for=condition=Ready

dev-restart:
	kubectl rollout restart deployment $(AUTH_APP) --namespace=$(NAMESPACE)
	kubectl rollout restart deployment $(SALES_APP) --namespace=$(NAMESPACE)

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(SALES_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run apis/tooling/logfmt/main.go -service=$(SALES_APP)

dev-logs-auth:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(AUTH_APP) --all-containers=true -f --tail=100 | go run apis/tooling/logfmt/main.go

# ------------------------------------------------------------------------------

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $(SALES_APP)

dev-describe-sales:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(SALES_APP)

dev-describe-auth:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(AUTH_APP)

# ==============================================================================
# Metrics and Tracing

metrics:
	expvarmon -ports="localhost:3010" -vars="build,requests,goroutines,errors,panics,mem:memstats.HeapAlloc,mem:memstats.HeapSys,mem:memstats.Sys"
statsviz:
	open http://localhost:3010/debug/statsviz
# ==============================================================================
#  Modules support

tidy:
	go mod tidy
	GOWORK=off go mod vendor

# ==============================================================================
# Running tests within the local computer

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-r lint vuln-check
