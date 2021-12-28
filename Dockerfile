FROM golang:1.18-rc-alpine3.15 as build-env
WORKDIR /app
COPY . .
RUN apk add --no-cache git && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o minimail

FROM alpine:3.15
WORKDIR /app
COPY --from=build-env /app/minimail .
COPY static static
COPY .docs .docs

EXPOSE 8080/tcp
USER 1001
ENTRYPOINT ["./minimail"]
