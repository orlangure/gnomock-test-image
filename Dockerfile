FROM golang:latest as builder

WORKDIR /app
ADD go.mod .

RUN go mod download

ADD . .
ARG GOARCH
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
