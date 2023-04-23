FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/order_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy order_service folder to working directory
COPY ./order_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib
## goods service schema
COPY ./goods_service/schema /app/goods_service/schema
## statistic service schema
COPY ./statistic_service/schema /app/statistic_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/order-service/main

COPY --from=build-stage /app/order_service/main/order-service /order-service

# expose port, use order service port (14070)
EXPOSE 14070

# run
ENTRYPOINT ["/order-service"]