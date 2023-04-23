FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/staff_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy staff_service folder to working directory
COPY ./staff_service ../

# copy other dependencies
## kafka lib
COPY ./library/kafka_lib /app/library/kafka_lib
## account service schema
COPY ./account_service/schema /app/account_service/schema
## branch service schema
COPY ./branch_service/schema /app/branch_service/schema
## goods service schema
COPY ./goods_service/schema /app/goods_service/schema

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o staff-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/staff-service/main

COPY --from=build-stage /app/staff_service/main/staff-service /staff-service

# expose port, use staff service port (14082)
EXPOSE 14082

# run
ENTRYPOINT ["/staff-service"]