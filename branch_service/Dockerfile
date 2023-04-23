FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/branch_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy branch_service folder to working directory
COPY ./branch_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o branch-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/branch-service/main

COPY --from=build-stage /app/branch_service/main/branch-service /branch-service

# expose port, use branch service port (14000)
EXPOSE 14000

# run
ENTRYPOINT ["/branch-service"]