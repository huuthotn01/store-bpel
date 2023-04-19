FROM golang:1.18 AS BUILD_STAGE

# set destination for copy
WORKDIR /app

# download go modules
COPY go.mod go.sum ./
RUN go mod download
