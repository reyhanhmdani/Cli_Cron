FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o binary ./

FROM alpine

COPY --from=builder /app/binary .
COPY --from=builder /app/migration/ ./app/migration/
#COPY cmd/database/ca.crt /etc/ssl/certs/


RUN apk add --no-cache ca-certificates
RUN update-ca-certificates

WORKDIR /app

CMD ["/binary", "delete"]


# sudo docker run -it --network=host my-image:create
# sudo docker run -it --network=host my-image:delete
# sudo docker run -it --network=host my-image:get
# sudo docker run -it --network=host my-image:getall
