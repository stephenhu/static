FROM golang as builder
WORKDIR /sources
COPY . .
RUN go get github.com/eknkc/amber && \
    go get github.com/gorilla/mux && \
    go get github.com/russross/blackfriday && \
    go get github.com/fatih/color && \
    go get gopkg.in/alecthomas/kingpin.v2 && \
    go build

FROM ubuntu
WORKDIR /usr/local/static
COPY --from=builder /sources/static .
RUN apt-get -y update && apt-get install -y
CMD ["/usr/local/static/static", "build"]
