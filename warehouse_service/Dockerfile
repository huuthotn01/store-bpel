FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/warehouse_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy warehouse_service folder to working directory
COPY ./warehouse_service ../

# copy other dependencies
## staff service schema
COPY ./staff_service/schema /app/staff_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o warehouse-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/warehouse-service/main

COPY --from=build-stage /app/warehouse_service/main/warehouse-service /warehouse-service

# expose port, use warehouse service port (14081)
EXPOSE 14081

# run
ENTRYPOINT ["/warehouse-service"]