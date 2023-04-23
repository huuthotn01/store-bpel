FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/account_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy account_service folder to working directory
COPY ./account_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib
## cart service schema
COPY ./cart_service/schema /app/cart_service/schema
## customer service schema
COPY ./customer_service/schema /app/customer_service/schema
## staff service schema
COPY ./staff_service/schema /app/staff_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o account-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/account-service/main

COPY --from=build-stage /app/account_service/main/account-service /account-service

# expose port, use account service port (14083)
EXPOSE 14083

# run
ENTRYPOINT ["/account-service"]