FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/customer_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy customer_service folder to working directory
COPY ./customer_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o customer-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/customer-service/main

COPY --from=build-stage /app/customer_service/main/customer-service /customer-service

# expose port, use customer service port (14050)
EXPOSE 14050

# run
ENTRYPOINT ["/customer-service"]