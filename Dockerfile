FROM golang:1.17 as base
RUN mkdir /attn
ADD . /attn
WORKDIR /attn
RUN go clean --modcache
RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go test ./... -v
RUN go build

EXPOSE 5001

CMD ["./attn"]
