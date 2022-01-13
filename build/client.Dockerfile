FROM golang:1.17.6 AS build
WORKDIR /app
COPY . ./
RUN BUILD_PREFIX=/ make build_client

FROM scratch
WORKDIR /
COPY --from=build /app/client /
ENTRYPOINT ["/client"]