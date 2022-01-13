FROM golang:1.17.6 AS build
WORKDIR /app
COPY . ./
RUN BUILD_PREFIX=/ make build_server

FROM scratch
WORKDIR /
COPY --from=build /app/server /
COPY --from=build /app/resources /resources
EXPOSE 8999
ENTRYPOINT ["/server"]