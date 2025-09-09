CERT_DIR := /tmp/ontap-go
HOST_IP := $(shell ip route get 8.8.8.8 2>/dev/null | awk -F"src " 'NR==1{split($$2,a," ");print a[1]}' || echo "127.0.0.1")
GO := $(shell which go || echo "/usr/local/go/bin/go")
release:: generate-client generate-server mocks gofmt test;

.PHONY: generate-client
generate-client:
	rm -rf api
	mkdir -p api
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-v ${PWD}:/work \
		-w /work \
		ghcr.io/metal-stack/builder swagger generate client -A Ontap -f spec/ontap.yaml -t api --struct-tags json --struct-tags yaml
	rm api/models/application_template.go # this redeclares ApplicationTemplate which is a BUG in the ontap swagger spec.


.PHONY: generate-server
generate-server:
	# https://goswagger.io/go-swagger/generate/server/
	rm -rf pkg/server/cmd pkg/server/models pkg/server/restapi
	mkdir -p pkg/server/fake
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-v ${PWD}:/work \
		-w /work \
		ghcr.io/metal-stack/builder swagger generate server -A ontap-fake-server -f spec/ontap.yaml -t pkg/server --skip-models --existing-models=github.com/metal-stack/ontap-go/api/models

.PHONY: mocks
mocks:
	rm -rf test/mocks
	docker run --rm \
		--user $$(id -u):$$(id -g) \
		-w /work \
		-v ${PWD}:/work \
		vektra/mockery:v2.52.3 -r --keeptree --inpackage --dir api/client --output test/mocks --all

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: start-dev-server
start-dev-server:
	@echo 'Using Host ip $(HOST_IP)'
	rm -rf $(CERT_DIR)
	mkdir -p $(CERT_DIR)
	openssl req -x509 -newkey rsa:4096 -keyout $(CERT_DIR)/key.pem -out $(CERT_DIR)/cert.pem -days 365 -nodes -subj "/CN=localhost"
	sudo $(GO) run pkg/server/cmd/ontap-fake-server-server/main.go --tls-certificate $(CERT_DIR)/cert.pem --tls-key $(CERT_DIR)/key.pem --tls-port 443 --tls-host $(HOST_IP)

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out
