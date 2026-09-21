FROM golang:1.27-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /out/data \
    && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /out/go-login-cli ./cmd

FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /app
COPY --from=build /out/go-login-cli ./go-login-cli
COPY --from=build --chown=65532:65532 /out/data ./data

VOLUME ["/app/data"]
ENTRYPOINT ["/app/go-login-cli"]
