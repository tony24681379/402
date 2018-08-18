# build stage
ARG GO_VERSION=1.10.3
FROM golang:${GO_VERSION}-alpine3.7 as builder
ADD $WORKSPACE /go/src/github.com/tony24681379/402
ENV TZ Asia/Taipei
RUN apk --no-cache add make tzdata git
WORKDIR /go/src/github.com/tony24681379/402
RUN go install

# final stage
FROM alpine:3.7
RUN apk --no-cache add ca-certificates tzdata curl
ENV TZ Asia/Taipei
WORKDIR /app
COPY --from=builder /go/bin/402 /app/402

EXPOSE 4000

CMD ["/app/402"]