FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/goods_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy goods_service folder to working directory
COPY ./goods_service ../

# copy other dependencies
## event service schema
COPY ./event_service/schema /app/event_service/schema
## order service schema
COPY ./order_service/schema /app/order_service/schema
## warehouse service schema
COPY ./warehouse_service/schema /app/warehouse_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o goods-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/goods-service/main

COPY --from=build-stage /app/goods_service/main/goods-service /goods-service

# expose port, use goods service port (14080)
EXPOSE 14080

# run
ENTRYPOINT ["/goods-service"]