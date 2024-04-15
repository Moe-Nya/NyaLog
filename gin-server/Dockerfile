FROM golang:1.21.6 AS builder

WORKDIR /app

ENV GOPROXY https://goproxy.cn,direct

EXPOSE 1300

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o dist ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /app/dist /app/server
CMD sh -c "sleep 5 && /app/server"
