FROM golang:1.18 AS BUILD_STAGE

# set destination for copy
WORKDIR /app

# download go modules
COPY . /app
RUN go mod download

# install dependencies
RUN go get .

# build
RUN CGO_ENABLED=0 GOOS=linux cd ./bff/admin_bff/main && go build -o admin-bff .

# expose port, use bff admin port (10000)
EXPOSE 10000

# run
CMD ["admin-bff"]