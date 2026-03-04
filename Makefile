release:: generate-client mocks gofmt test;

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
	# The ontap swagger spec uses hyphens for "volume-count" but the actual API returns "volume_count".
	# The spec itself states: "REST API properties use underscores instead of hyphens between words."
	sed -i 's/volume-count/volume_count/g' api/models/aggregate.go

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

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out
