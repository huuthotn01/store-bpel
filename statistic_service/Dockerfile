FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/statistic_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy statistic_service folder to working directory
COPY ./statistic_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o statistic-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/statistic-service/main

COPY --from=build-stage /app/statistic_service/main/statistic-service /statistic-service

# expose port, use statistic service port (14090)
EXPOSE 14090

# run
ENTRYPOINT ["/statistic-service"]