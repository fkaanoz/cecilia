#################### DEFINITIONS ########################

service 		:= cecilia
clusterName 	:= cecilia-cluster
serviceNS 		:= cecilia-system
imageName 		:= fkaanoz/$(service):latest
release 		:= `git rev-parse --short=8 HEAD`

#################### LOCAL ########################

tidy:
	go mod tidy && go mod vendor

run:
	go run ./cmd/$(service)/main.go

local-build:
	go build -o bin/$(service) ./cmd/$(service)/main.go

hot-reload:
	air --build.cmd "go build -o bin/$(service) cmd/$(service)/main.go" --build.bin "./bin/$(service)"


#################### KIND CLUSTER ########################

cluster-up:
	kind create cluster --name $(clusterName) --config=./zarf/kind/config.yaml
	kubectl create namespace $(serviceNS)
	kubectl config set-context --current --namespace=$(serviceNS)

cluster-down:
	kind delete cluster --name $(clusterName)

cluster-image-upload:
	kind load docker-image --name $(clusterName) $(imageName)

cluster-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --namespace=$(serviceNS)

cluster-apply-deployment:
	kubectl apply -f ./zarf/k8s/deployment.yaml

#
#cluster-logs:
#	kubectl logs pods/

cluster-deploy: image-build cluster-image-upload cluster-apply-deployment

#################### DOCKER ########################

image-build:
	docker build -t $(imageName)  -f ./zarf/docker/Dockerfile --build-arg RELEASE=$(release) .
