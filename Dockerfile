FROM golang:1.9

RUN mkdir -p /go/src/github.com/transactcharlie/riemann-spawn
WORKDIR /go/src/github.com/transactcharlie/riemann-spawn
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o riemann-spawn .

FROM scratch
COPY --from=0 /go/src/github.com/transactcharlie/riemann-spawn/riemann-spawn /riemann-spawn
CMD ["/riemann-spawn"]
