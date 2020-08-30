# whois json

[![aasaam](https://flat.badgen.net/badge/aasaam/software%20development%20group/0277bd?labelColor=000000&icon=https%3A%2F%2Fcdn.jsdelivr.net%2Fgh%2Faasaam%2Finformation%2Flogo%2Faasaam.svg)](https://github.com/aasaam)

[![travis](https://flat.badgen.net/travis/aasaam/whois-json)](https://travis-ci.org/aasaam/whois-json)
[![coveralls](https://flat.badgen.net/coveralls/c/github/aasaam/whois-json)](https://coveralls.io/github/aasaam/whois-json)
[![go-report-card](https://goreportcard.com/badge/github.com/gojp/goreportcard?style=flat-square)](https://goreportcard.com/report/github.com/aasaam/whois-json)

[![open-issues](https://flat.badgen.net/github/open-issues/aasaam/whois-json)](https://github.com/aasaam/whois-json/issues)
[![open-pull-requests](https://flat.badgen.net/github/open-prs/aasaam/whois-json)](https://github.com/aasaam/whois-json/pulls)
[![license](https://flat.badgen.net/github/license/aasaam/whois-json)](./LICENSE)

Simple tool for parse and create structured json for domain who is information.

## Why

Because Whois data is raw and we need support structured data.

## Download

Use [releases](https://github.com/aasaam/whois-json/releases) to download latest binary

## Usage

```bash
./whois-json -h
```

### CLI

```bash
./whois-json json -d github.com
./whois-json validate -d sub-domain-not-matter.google.com
```

### REST-API

```bash
./whois-json webserver
```

```bash
curl -s http://username:password@localhost:9000/whois/example.com | jq
```

## Thanks

Thanks to [Li Kexian](https://github.com/likexian) and others that help to build this tool.

* [whois-go](https://github.com/likexian/whois-go)
* [whois-parser-go](https://github.com/likexian/whois-parser-go)
* [dateparse](https://github.com/araddon/dateparse)
* [urfave/cli](https://github.com/urfave/cli/v2)
* [gofiber](https://github.com/gofiber/fiber)
