FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/bff/customer_bff/main

COPY go.mod go.sum ../../../
RUN go mod download

# copy customer-bff folder to working directory
COPY ./bff/customer_bff ../

# copy other dependencies
## account service schema
COPY ./account_service/schema /app/account_service/schema
## cart service schema
COPY ./cart_service/schema /app/cart_service/schema
## customer service schema
COPY ./customer_service/schema /app/customer_service/schema
## order service schema
COPY ./order_service/schema /app/order_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o customer-bff

FROM alpine:3 AS build-release-stage

WORKDIR /app/bff/customer-bff/main

COPY --from=build-stage /app/bff/customer_bff/main/customer-bff /customer-bff

# expose port, use bff admin port (10000)
EXPOSE 20000

# run
CMD ["/customer-bff"]