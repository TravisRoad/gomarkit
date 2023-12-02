from golang:1.21.1-alpine3.18 as base

ENV GOPROXY https://goproxy.cn,direct
RUN apk add --no-cache --update gcc g++
WORKDIR /go/build
COPY . /go/build
RUN go mod download
RUN CGO_ENABLED=1 go build -ldflags "-s -w" -o /go/build/exe

FROM alpine:3.18

ENV MODE="PROD"
ENV PORT="8080"
WORKDIR /app
COPY --from=base /go/build/exe /
COPY --from=base /go/build/config.prod.yaml /app
EXPOSE 8080
CMD /exe

