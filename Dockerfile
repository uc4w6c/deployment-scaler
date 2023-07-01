# Build
FROM golang:1.20-buster As Build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /deployment-scaler

# Deploy
FROM gcr.io/distroless/base-debian11:latest

WORKDIR /

COPY --from=build /deployment-scaler /deployment-scaler

USER nonroot:nonroot

ENTRYPOINT ["/deployment-scaler"]
