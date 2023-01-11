FROM golang as builder
WORKDIR /sources
COPY . .
RUN go build

FROM ubuntu
WORKDIR /usr/local/static
COPY --from=builder /sources/static .
RUN apt-get -y update && apt-get install -y
CMD ["/usr/local/static/static", "generate", "index"]
