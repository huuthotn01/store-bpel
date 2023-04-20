FROM base-image:latest AS BUILD_STAGE

# working directory
WORKDIR /app/event_service/main

# copy admin-bff folder to working directory
COPY ./event_service ../

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /event-service

# expose port, use bff admin port (10000)
EXPOSE 14060

# run
CMD ["/event-service"]