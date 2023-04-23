FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/bff/shared_bff/main

COPY go.mod go.sum ../../../
RUN go mod download

# copy shared-bff folder to working directory
COPY ./bff/shared_bff ../

# copy other dependencies
## account service schema
COPY ./account_service/schema /app/account_service/schema
## event service schema
COPY ./event_service/schema /app/event_service/schema
## goods service schema
COPY ./goods_service/schema /app/goods_service/schema
## order service schema
COPY ./order_service/schema /app/order_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o shared-bff

FROM alpine:3 AS build-release-stage

WORKDIR /app/bff/shared-bff/main

COPY --from=build-stage /app/bff/shared_bff/main/shared-bff /shared-bff

# expose port, use bff admin port (10000)
EXPOSE 30000

# run
CMD ["/shared-bff"]