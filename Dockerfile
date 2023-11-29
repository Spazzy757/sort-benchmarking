FROM golang:1.21

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY pkg ./pkg
COPY *.go ./

