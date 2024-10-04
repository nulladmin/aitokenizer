FROM golang:1.23.0 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY config.go .
COPY go.sum .
COPY go.mod .
COPY Makefile .

RUN make build

# RUN go build -o aitokenizer .

FROM scratch AS runtime
COPY --from=build /go/src/aitokenizer ./
EXPOSE 8080/tcp
ENTRYPOINT ["./aitokenizer"]
