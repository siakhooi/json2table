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
run-build:
	bin/json2table-linux-amd64 --build
run-1:
	bin/json2table-linux-amd64 ./test.json
run-3:
	cat ./test.json| bin/json2table-linux-amd64
run-2:
	bin/json2table-linux-amd64
run-i:
	bin/json2table-linux-amd64
	bin/json2table-linux-amd64 ./test.json ./test.json
