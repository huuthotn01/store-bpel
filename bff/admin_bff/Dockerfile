FROM base-image:latest AS BUILD_STAGE

# copy admin-bff folder to working directory
COPY . .

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux cd ./bff/admin_bff/main && go build -o /admin-bff

# expose port, use bff admin port (10000)
EXPOSE 10000

# run
CMD ["/admin-bff"]