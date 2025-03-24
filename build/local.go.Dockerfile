FROM golang:1.19

ENV GOPRIVATE=github.com/vunq41298

RUN apt-get update

RUN GO111MODULE=on go install golang.org/x/tools/cmd/goimports@v0.2.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.14.1 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.1 \
