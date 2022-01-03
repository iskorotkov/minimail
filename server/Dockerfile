FROM golang:1.18-rc-alpine3.15 as build
WORKDIR /app

RUN apk add --no-cache git

COPY ["server/go.mod", "server/go.sum", "./"]
RUN go mod download

COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o minimail

FROM alpine:3.15 as run
WORKDIR /app

COPY --from=build /app/minimail .
COPY ["openapi.yaml", "openapi.yml"]
COPY ["client-common", "static/common"]
COPY ["client-swagger-ui", "static/swagger-ui"]
COPY ["client-templates", "static/templates"]
COPY ["client-ajax", "static/ajax"]

EXPOSE 8080/tcp
USER 1001
ENTRYPOINT ["./minimail"]
