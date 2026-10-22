# syntax=docker/dockerfile:1

FROM golang:1.27.3 AS development

WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.7

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

FROM golang:1.27.3 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/wilayah-indonesia-api ./api

FROM gcr.io/distroless/static-debian12:nonroot AS production

WORKDIR /app

COPY --from=builder /out/wilayah-indonesia-api ./wilayah-indonesia-api

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/app/wilayah-indonesia-api"]
