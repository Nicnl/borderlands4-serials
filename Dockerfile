FROM golang:1.25.3 AS builder
MAINTAINER nicnl25@gmail.com

# 1) Install dependencies
WORKDIR /go/src/borderlands_4_serials
ADD go.mod go.sum ./
RUN go mod download -x

# 2) Add sources and compile
ADD . .
RUN GOOS=linux go build -ldflags "-s -w" -v -o /borderlands_4_serials ./cmd/api/main


FROM scratch
COPY --from=builder /borderlands_4_serials /borderlands_4_serials
EXPOSE 8080
CMD ["/borderlands_4_serials"]
