FROM golang:1.16-buster

ADD . /src

RUN cd /src \
  && go get -u -v golang.org/x/lint/golint \
  && golint . \
  && export CI=1 \
  && go test -covermode=count -coverprofile=coverage.tmp.out \
  && cat coverage.out | grep -v "http_whois.go" | grep -v "domain_whois.go" | grep -v "main.go" > coverage.txt \

