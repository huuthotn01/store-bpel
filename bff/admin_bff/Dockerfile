FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/bff/admin_bff/main

COPY go.mod go.sum ../../../
RUN go mod download

# copy admin-bff folder to working directory
COPY ./bff/admin_bff ../

# copy other dependencies
## account service schema
COPY ./account_service/schema /app/account_service/schema
## branch service schema
COPY ./branch_service/schema /app/branch_service/schema
## event service schema
COPY ./event_service/schema /app/event_service/schema
## goods service schema
COPY ./goods_service/schema /app/goods_service/schema
## order service schema
COPY ./order_service/schema /app/order_service/schema
## staff service schema
COPY ./staff_service/schema /app/staff_service/schema
## statistic service schema
COPY ./statistic_service/schema /app/statistic_service/schema
## warehouse service schema
COPY ./warehouse_service/schema /app/warehouse_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o admin-bff

FROM alpine:3 AS build-release-stage

WORKDIR /app/bff/admin-bff/main

COPY --from=build-stage /app/bff/admin_bff/main/admin-bff /admin-bff

# expose port, use bff admin port (10000)
EXPOSE 10000

# run
CMD ["/admin-bff"]