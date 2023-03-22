FROM golang:1.19-alpine3.15 AS build
RUN apk --no-cache add ca-certificates git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener ./cmd/main.go

FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN ls -l
COPY --from=build /app/url-shortener .
EXPOSE 8080
CMD ["./url-shortener"]