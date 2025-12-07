clean:
	rm -rf bin
build:
	scripts/build.sh
	goreleaser build --snapshot --clean
test:
	scripts/test.sh
golangci-lint:
	golangci-lint run

set-version:
	scripts/set-version.sh
all: clean set-version test golangci-lint build

commit:
	scripts/git-commit-and-push.sh

release:
	scripts/create-release.sh

commit-watch: commit
	gh run watch

release-watch: release
	gh run watch

run:
	bin/json2table-linux-amd64 -h
	bin/json2table-linux-amd64
	bin/json2table-linux-amd64 ./test.json
run-i:
	bin/json2table-linux-amd64 ./test.json ./test.json
