# json2table

[![License](https://img.shields.io/github/license/siakhooi/json2table)](https://github.com/siakhooi/json2table/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/siakhooi/json2table)](https://github.com/siakhooi/json2table/releases/latest)
[![Build](https://img.shields.io/github/actions/workflow/status/siakhooi/json2table/build.yaml?label=build)](https://github.com/siakhooi/json2table/actions/workflows/build.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/siakhooi/json2table.svg)](https://pkg.go.dev/github.com/siakhooi/json2table)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_json2table&metric=alert_status)](https://sonarcloud.io/project/overview?id=siakhooi_json2table)
[![Coverage](https://qlty.sh/gh/siakhooi/projects/json2table/coverage.svg)](https://qlty.sh/gh/siakhooi/projects/json2table)
[![Funding](https://img.shields.io/badge/Funding-Wise-33cb56.svg?logo=wise)](https://wise.com/pay/me/siakn3)
![visitors](https://hit-tztugwlsja-uc.a.run.app/?outputtype=badge&counter=ghmd-json2table)

print json data in tabular format

## Usage
```
NAME:
   json2table - print json data in tabular format

USAGE:
   json2table [global options] [dataFile]

VERSION:
   v1.1.0

GLOBAL OPTIONS:
   --build                      print build info and exit
   --spec string, -s string     read spec from specFile.json, or from environment variable JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if not provided
   --columns string, -c string  Comma separated list of columns to print, ignore -s and JSON2TABLE_SPEC or JSON2TABLE_SPEC_FILE if provided
   --help, -h                   show help
   --version, -v                print the version
```
## Examples

### Cli
```
$ json2table -s ./samples/spec1.json ./samples/data1.json

$ json2table -c 'id,desc,url,display.name' ./samples/data2.json

```
### Scripts with Spec

- [Picsum list](./examples/picsum-list.md)

## Installation

See [Installation.md](Installation.md) for Homebrew, Scoop, Linux packages, Windows winget, and manual binary installs.

## Spec Reference

- [Spec file guide](./SPEC.md)

## Quality

- https://sonarcloud.io/project/overview?id=siakhooi_json2table
- https://qlty.sh/gh/siakhooi/projects/json2table

## Deliverables

- https://pkg.go.dev/github.com/siakhooi/json2table

## Reference

- https://github.com/savioxavier/termlink
- https://github.com/fatih/color
