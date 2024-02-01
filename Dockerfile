FROM golang:1.19-alpine3.18 as buildbase

WORKDIR /go/src/github.com/rum-people-preseed/weather-harvester-svc

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/weather-harvester-svc main.go

FROM alpine:3.18

COPY --from=buildbase /usr/local/bin/weather-harvester-svc /usr/local/bin/weather-harvester-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["weather-harvester-svc"]