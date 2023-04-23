FROM golang:1.18 AS build-stage

# working directory
WORKDIR /app/event_service/main

COPY go.mod go.sum ../../
RUN go mod download

# copy event_service folder to working directory
COPY ./event_service ../

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o event-service

FROM alpine:3 AS build-release-stage

WORKDIR /app/event-service/main

COPY --from=build-stage /app/event_service/main/event-service /event-service

# expose port, use event service port (14060)
EXPOSE 14060

# run
ENTRYPOINT ["/event-service"]