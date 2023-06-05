FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o binary ./

FROM scratch

COPY --from=builder /app/binary .
COPY --from=builder /app/migration/ ./app/migration/

WORKDIR /app

CMD ["/binary", "getall"]


# sudo docker run -it --network=host my-image:create
# sudo docker run -it --network=host my-image:delete
# sudo docker run -it --network=host my-image:get
# sudo docker run -it --network=host my-image:getall
