all: compile build push deploy

test:
	go test ./...

dep:
	dep ensure

pretty:
	gofmt -d -w $$(find . -type f -name '*.go' -not -path "./vendor/*")
	goimports -d -w $$(find . -type f -name '*.go' -not -path "./vendor/*")
	go tool vet $$(find . -type f -name '*.go' -not -path "./vendor/*")

cover:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out