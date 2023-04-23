FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/cart_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy cart_service folder to working directory
COPY ./cart_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib
## goods service schema
COPY ./goods_service/schema /app/goods_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o cart-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/cart-service/main

COPY --from=build-stage /app/cart_service/main/cart-service /cart-service

# expose port, use cart service port (14061)
EXPOSE 14061

# run
ENTRYPOINT ["/cart-service"]