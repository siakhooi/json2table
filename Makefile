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
quick: clean test golangci-lint build-linux
commit:
	scripts/git-commit-and-push.sh

release:
	scripts/create-release.sh

commit-watch: commit
	gh run watch

release-watch: release
	gh run watch

run:
	bin/json2table-linux-amd64
run-help:
	bin/json2table-linux-amd64 -h
run-version:
	bin/json2table-linux-amd64 -v
run-build:
	bin/json2table-linux-amd64 --build
run-no-arguments-1:
	bin/json2table-linux-amd64
run-no-arguments-2:
	bin/json2table-linux-amd64 -s ./samples/spec1.json
run-too-many-arguments-1:
	bin/json2table-linux-amd64 ./samples/data1.json ./samples/data1.json
run-too-many-arguments-2:
	bin/json2table-linux-amd64 -s ./samples/spec.json ./samples/data1.json ./samples/data1.json
run-1:
	bin/json2table-linux-amd64 -s ./samples/spec1.json ./samples/data1.json
run-1a:
	bin/json2table-linux-amd64 -s ./samples/spec2.json ./samples/data1.json
run-2:
	bin/json2table-linux-amd64 --spec ./samples/spec1.json ./samples/data1.json
run-3:
	cat ./samples/data1.json |bin/json2table-linux-amd64 -s ./samples/spec1.json
run-4:
	JSON2TABLE_SPEC_FILE=./samples/spec1.json bin/json2table-linux-amd64 ./samples/data1.json
run-5:
	cat  ./samples/data1.json |JSON2TABLE_SPEC_FILE=./samples/spec1.json bin/json2table-linux-amd64
run-6:
	JSON2TABLE_SPEC='{"dataPath":"$.data2","columns":[{"path":"id","title":"ID"},{"path":"display.name"}]}' bin/json2table-linux-amd64 ./samples/data1.json
