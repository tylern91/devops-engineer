FROM golang:alpine AS builder
LABEL maintainer="devops-team@nfq.asia"
WORKDIR /app
ADD . /app/
RUN cd src/main && go build -o nfqasia
From golang:1.14.4-alpine3.12
WORKDIR /
COPY --from=builder /app/src/main/nfqasia .
COPY --from=builder /app/src/main/welcome .
COPY --from=builder /app/src/scripts/nfqasia.go /var/run/nfqasia.go
ENTRYPOINT ./nfqasia
