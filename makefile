#################### DEFINITIONS ########################
projectName := cecilia
clusterName := cecilia-cluster
imageName := fkaanoz/$(projectName):latest
build :=
service :=


#################### LOCAL ########################

tidy:
	go mod tidy && go mod vendor

run:
	go run ./cmd/$(projectName)/main.go

local-build:
	go build -o bin/$(projectName) ./cmd/$(projectName)/main.go

hot-reload:
	air --build.cmd "go build -o bin/$(projectName) cmd/$(projectName)/main.go" --build.bin "./bin/$(projectName)"


#################### KIND CLUSTER ########################

cluster-up:
	kind create cluster --name $(clusterName) --config=./zarf/kind/config.yaml

cluster-down:
	kind delete cluster --name $(clusterName)



#################### IMAGE ########################

image-build:
	docker build -t $(imageName)  -f ./zarf/docker/Dockerfile .
