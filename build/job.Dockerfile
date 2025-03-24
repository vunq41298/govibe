ARG RELEASE_IMAGE_NAME=alpine
ARG RELEASE_IMAGE_TAG=3.18

FROM ${RELEASE_IMAGE_NAME}:${RELEASE_IMAGE_TAG} AS base

RUN apk --no-cache add tzdata ca-certificates

FROM scratch
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./api/binaries/job /
ENTRYPOINT ["/job"]
