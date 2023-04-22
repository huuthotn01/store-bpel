FROM base-image:latest AS build-stage

# working directory
WORKDIR /app/event_service/main

# copy admin-bff folder to working directory
COPY ./event_service ../

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /event-service

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /event-service/main

COPY --from=build-stage /event-service ../

# expose port, use event service port (14060)
EXPOSE 14060

# run
ENTRYPOINT ["/event-service"]