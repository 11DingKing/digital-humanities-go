FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/digital-humanities .
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /out/digital-humanities /app/digital-humanities
EXPOSE 8080
ENTRYPOINT ["/app/digital-humanities"]
