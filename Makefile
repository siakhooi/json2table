clean:
	rm -rf bin
build:
	scripts/build.sh
	goreleaser build --snapshot --clean
build-linux:
	scripts/build.sh -l
test:
	scripts/test.sh
golangci-lint:
	golangci-lint run

all: clean test golangci-lint build

commit:
	scripts/git-commit-and-push.sh

release:
	scripts/create-release.sh

commit-watch: commit
	gh run watch

release-watch: release
	gh run watch

run-help:
	bin/json2table-linux-amd64 -h
run-version:
	bin/json2table-linux-amd64 -v
run-no-arguments:
	bin/json2table-linux-amd64
run-too-many-arguments:
	bin/json2table-linux-amd64 ./test.json ./test.json
run-build:
	bin/json2table-linux-amd64 --build
run-input-1:
	bin/json2table-linux-amd64 ./test.json
run-input-3:
	cat ./test.json| bin/json2table-linux-amd64
run-spec-1:
	bin/json2table-linux-amd64 -s ./spec.json ./test.json
run-spec-2:
	bin/json2table-linux-amd64 --spec ./spec.json ./test.json
run-spec-3:
	cat ./test.json |bin/json2table-linux-amd64 -s ./spec.json 
run-spec-4:
	JSON2TABLE_SPEC_FILE=./spec.json bin/json2table-linux-amd64 ./test.json
run-spec-5:
	cat  ./test.json |JSON2TABLE_SPEC_FILE=./spec.json bin/json2table-linux-amd64
